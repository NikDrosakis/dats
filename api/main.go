package main

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"
    "log"

    "github.com/gocql/gocql"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
    	"github.com/redis/go-redis/v9"
    _ "main/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "main/src/global_model"
)

func main() {
    log.Println("Server starting...")

    var r = gin.Default()
    var _cassandraCluster *gocql.ClusterConfig
    var cassandraConn *gocql.Session
    var serial string
	var redisConn *redis.Client
	var session string
	var userId int
    // ========== SQLITE CONNECTION ==========
    sqlitePath := "mnt/db_sqlite/local.db"
    sqliteSchemaPath := "db/sqlite/schema.sql"

    // Create directory if not exists
    err := os.MkdirAll(filepath.Dir(sqlitePath), os.ModePerm)
    if err != nil {
        fmt.Printf("Error creating data directory: %v\n", err)
        return
    }

	// redis
	// redis : conn
	redisConn = redis.NewClient(&redis.Options{
		Addr:     "db_redis:6379",
		Password: "",
		DB:       0, // Use default DB
	})

    // Cassandra connection
    _cassandraCluster = gocql.NewCluster("db_cassandra")
    _cassandraCluster.Keyspace = "stsy_ks"

    // Open SQLite connection
    sqliteConn, err := sql.Open("sqlite3", sqlitePath)
    if err != nil {
        fmt.Printf("Error opening database: %v\n", err)
        return
    }
    defer sqliteConn.Close()

    // PRAGMA settings
    sqliteConn.Exec("PRAGMA journal_mode=WAL;")
    sqliteConn.Exec("PRAGMA busy_timeout=60000;")

    // Execute schema
    schema, err := os.ReadFile(sqliteSchemaPath)
    if err != nil {
        fmt.Printf("Error reading schema file: %v\n", err)
        return
    }
    sqliteConn.Exec(string(schema))

    // Serial
    serial = os.Getenv("SERIAL")

    // Environment
    env := global_model.TypeEnv{
        COMPOSE_PROFILES:                       os.Getenv("COMPOSE_PROFILES"),
        TRANSMITTER_FALLOCATE_MB:               os.Getenv("TRANSMITTER_FALLOCATE_MB"),
        TRANSMITTER_FALLOCATE_ENABLE:           os.Getenv("TRANSMITTER_FALLOCATE_ENABLE"),
        SAVE_LAST_HOUR_INTERVAL_PER_SERIAL_SEC: os.Getenv("SAVE_LAST_HOUR_INTERVAL_PER_SERIAL_SEC"),
        TTL_RECORDS_HISTORY_DAYS:               os.Getenv("TTL_RECORDS_HISTORY_DAYS"),
        TTL_ALARMS_HISTORY_DAYS:                os.Getenv("TTL_ALARMS_HISTORY_DAYS"),
    }

    // Context
    ctx := global_model.TypeCtx{
        GinEngine:            r,
        RedisConn:            nil, // Add Redis if needed
        CassandraConn:        cassandraConn,
        SqliteConn:           sqliteConn,
        Serial:               serial,
        Env:                  env,
    }

    _ = ctx // Use ctx to avoid unused variable warning

    // ========== ROUTES ==========
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Health check
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // ========== START SERVER ==========
    log.Println("API is running on port 8083")
    r.Run(":8083")
}