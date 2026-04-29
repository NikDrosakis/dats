package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	//"path/filepath"
	"regexp"
	//"strconv"
	"strings"

	"log"

	"github.com/gin-gonic/gin"

	alarms "main/src/alarms"
//	"main/src/controllers"
//	"main/src/controllers/db"
//	"main/src/views"
//	"main/src/models/global_model"

	// variantsLicenseKey "main/src/views/license_key"
	// variantsLicenseKeyModel "main/src/views/variants/variant_license_key/models/model"

	"github.com/gocql/gocql"
//	"github.com/redis/go-redis/v9"

	_ "github.com/mattn/go-sqlite3"
    _ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/gzip"

	_ "main/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apiKey HeadersUserId
// @in header
// @name User-Id

// @securityDefinitions.apiKey HeadersSession
// @in header
// @name Session

// @securityDefinitions.apiKey Cookie
// @in header
// @name Cookie
func main() {
	var err error
//	var serial string
    r := gin.Default()
	var _cassandraCluster *gocql.ClusterConfig
	var cassandraConn *gocql.Session
	var redisConn *redis.Client

	log.Println("Server starting...")
  //  var mariaSchemaPath string
  //  var mariaConn *sql.DB

//	var sqliteSchemaPath string
//	var sqlitePath string
	var sqliteConn *sql.DB

//	var sqliteMiniExtrasSchemaPath string
	//var sqliteMiniExtrasPath string
	var sqliteMiniExtrasConn *sql.DB

	//var env global_model.TypeEnv

//	var ctx global_model.TypeCtx

	/*** * * ***/

	// cassandraConn
	defer (func(cassandraConn *gocql.Session) {
		if os.Getenv("COMPOSE_PROFILES") == "transmitter_mini" {
			return
		}

		if os.Getenv("COMPOSE_PROFILES") == "receiver_mini" {
			return
		}

		/*** * * ***/

		defer cassandraConn.Close()
	})(cassandraConn)


    // mariadb
    // mariadb : schema
    //mariaSchemaPath = "db/maria/schema.sql"

    // mariadb : conn
 /*   dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
        "root",                // root
        "130177",            // 130177
        "db_mariadb",                         // container hostname
        "3306",
        "stsy_db",        // stsy_db
    )
*/
    sqliteConn, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("MariaDB conn error: %v", err)
    }

    if err = sqliteConn.Ping(); err != nil {
        log.Fatalf("MariaDB ping error: %v", err)
    }

    // mariadb : schema : read
  // mariaSchema, err := os.ReadFile(mariaSchemaPath)
//    if err != nil {
  //      log.Fatalf("MariaDB schema read error: %v", err)
//    }

    // mariadb : schema : exec
  //  _, err = mariaConn.Exec(string(mariaSchema))
//    if err != nil {
  //      log.Fatalf("MariaDB schema exec error: %v", err)
//    }



	// sqliteConn
	defer sqliteConn.Close()

	// sqliteMiniExtrasConn
	defer sqliteMiniExtrasConn.Close()

	/*** * * ***/

	// serial
//	serial = os.Getenv("SERIAL")

	// cassandra

	// cassandra : cluster
	_cassandraCluster = gocql.NewCluster("db_cassandra")
	// cassandra : cluster : keyspace
	_cassandraCluster.Keyspace = "stsy_ks"

	// cassandra : conn
	(func() {
		if os.Getenv("COMPOSE_PROFILES") == "transmitter_mini" {
			return
		}

		if os.Getenv("COMPOSE_PROFILES") == "receiver_mini" {
			return
		}

		/*** * * ***/

		cassandraConn, err = _cassandraCluster.CreateSession()
		if err != nil {
			log.Fatal(err)
		}
	}())

	// redis
	// redis : conn
/*	redisConn = redis.NewClient(&redis.Options{
		Addr:     "db_redis:6379",
		Password: "",
		DB:       0, // Use default DB
	})
*/
	// sqlite
	// sqlite : path

	//sqlitePath = "mnt/db_sqlite/local.db"
	// sqlite : schema
	// sqlite : schema : path
//	sqliteSchemaPath = "db/maria/schema.sql"
	//sqliteSchemaPath = "db/sqlite/init/schema.sql"
	//err = os.MkdirAll(filepath.Dir(sqlitePath), os.ModePerm)
	//if err != nil {
	//	fmt.Printf("Error creating data directory: %v\n", err)
	//	return
	//}

	// sqlite : conn
	//sqliteConn, err = sql.Open("sqlite3", sqlitePath)
	//if err != nil {
	//	fmt.Printf("Error opening database: %v\n", err)
//		return
//	}
//	sqliteConn.Exec("PRAGMA journal_mode=WAL;")
//	sqliteConn.Exec("PRAGMA busy_timeout=60000;")

	// sqlite mini extras
	(func() {
		if os.Getenv("COMPOSE_PROFILES") == "transmitter" {
			return
		}

		if os.Getenv("COMPOSE_PROFILES") == "receiver" {
			return
		}

		/*** * * ***/

		// sqlite mini extras : path
	//	sqliteMiniExtrasPath = "mnt/db_sqlite_mini_extras/local.db"
		// sqlite mini extras : schema
		// sqlite mini extras : schema : path
	//	sqliteMiniExtrasSchemaPath = "db/maria/schema_mini_extras.sql"
		//sqliteMiniExtrasSchemaPath = "db/sqlite_mini_extras/init/schema.sql"


		//err = os.MkdirAll(filepath.Dir(sqliteMiniExtrasPath), os.ModePerm)
		//if err != nil {
		//	fmt.Printf("Error creating data directory: %v\n", err)
		//	return
		//}
        dsnMini := fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
        "root",                // root
        "130177",            // 130177
        "db_mariadb",                         // container hostname
        "3306",
        "stsy_db",        // stsy_db
        )
		// sqlite mini extras : conn
		sqliteMiniExtrasConn, err = sql.Open("mysql", dsnMini)
		if err != nil {
			fmt.Printf("Error opening database: %v\n", err)
			return
		}
	//	sqliteMiniExtrasConn.Exec("PRAGMA journal_mode=WAL;")
	//	sqliteMiniExtrasConn.Exec("PRAGMA busy_timeout=60000;")
	}())

	/*** * * ***/

	// sqlite
	// sqlite : schema
//	sqliteSchema, err := os.ReadFile(sqliteSchemaPath)
//	if err != nil {
//		fmt.Printf("Error reading schema file: %v\n", err)
//		return
//	}
	// sqlite : schema : exec
	//_, err = sqliteConn.Exec(string(sqliteSchema))
	//if err != nil {
	//	fmt.Printf("Error executing schema SQL: %v\n", err)
	//	return
	//}

	// sqlite mini extras
	(func() {
		if os.Getenv("COMPOSE_PROFILES") == "transmitter" {
			return
		}

		if os.Getenv("COMPOSE_PROFILES") == "receiver" {
			return
		}

		/*** * * ***/

		// sqlite mini extras : schema
	//	sqliteMiniExtrasSchema, err := os.ReadFile(sqliteMiniExtrasSchemaPath)
	//	if err != nil {
	//		fmt.Printf("Error reading schema file: %v\n", err)
	//		return
	//	}
		// sqlite mini extras : schema : exec
	//	_, err = sqliteMiniExtrasConn.Exec(string(sqliteMiniExtrasSchema))
	//	if err != nil {
	//		fmt.Printf("Error executing schema SQL: %v\n", err)
	//		return
	//	}
	}())

	// gin

	// gin : mode
	gin.SetMode(gin.ReleaseMode)

	// gin : engine : CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Origin"+
				", "+"Authorization"+
				", "+"Content-Type"+
				", "+"Api-Key", // legacy, removed (predecessor of HEADER_SESSION)
			//	", "+global_model.CONST_HTTP_HEADER_USER_ID_KEY+
			//	", "+global_model.CONST_HTTP_HEADER_SESSION_KEY,
		)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		/*** * * ***/

		c.Next()
	})

	// gin : engine : compression
	r.Use(gzip.Gzip(gzip.BestCompression))

	// gin : engine : auth
	r.Use(func(c *gin.Context) {
		var err error

		var userIdStr string
//		var userId int

		var session string

		var checkPass bool

		/*** * * ***/

		if !strings.HasPrefix(c.Request.URL.Path, "/v2") {
			return
		}

		/*** * * ***/

		// check skip
		if (func() bool {
			var checkSkip bool

			var checkSkipPaths []string

			/*** * * ***/

			// checkSkip
			checkSkip = false

			// checkSkip : paths
			checkSkipPaths = []string{
				// records

				// records : set

				// records : set : io
				"/v2/io/records/set/[^/]",

				// records : set : alias
				"/v2/set/[^/]", // alias of "/v2/io/records/set/[^/]"

				// records : get

				// records : get : db redis
				"/v2/db/redis/records_latests/get_stream", // legacy, alias of "/v2/db/redis/records_latest/get_stream"
				"/v2/db/redis/records_latest/get_stream",

				// records : abstract

				// records : abstract : live
				"/v2/io/records/abstract/live",

				// user

				// user : auth
				"/v2/user/auth", // legacy
				"/v2/io/users/auth",

				// commands

				// commands : get
				"/v2/io/commands/get/[^/]",

				// alarms

				// alarms : events

				"/v2/io/alarms/events/set",

				// modules

				// modules : ping
				"/v2/modules/ping",

				"/swagger/[^/]",
			}

			/*** * * ***/

			// checkSkip (based on skipCheckPaths)
			for _, skipCheckPath := range checkSkipPaths {
				if regexp.MustCompile(`^` + skipCheckPath + `+(/|$)$`).MatchString(c.Request.URL.Path) {
					checkSkip = true
				}
			}

			/*** * * ***/

			return checkSkip
		})() {
			return
		}

		/*** * * ***/

		// checkPass
		checkPass = false

		// userIdStr

	//	userIdStr, err = c.Cookie(global_model.CONST_HTTP_COOKIE_USER_ID_KEY)
	//	if err != nil {
	//		userIdStr = ""
	//		err = nil
	//	}

	//	if len(userIdStr) == 0 {
	//		userIdStr = c.GetHeader(global_model.CONST_HTTP_HEADER_USER_ID_KEY)
	//	}

		if len(userIdStr) == 0 {

	//		var r global_model.TypeRes
//			r.Status.ErrCode = 2.0
	//		r.Status.ErrDesc = "Please Provide a Cookie or Header " + global_model.CONST_HTTP_HEADER_USER_ID_KEY

		//	c.AbortWithStatusJSON(400, r)

			return
		}

		// userId (based on userIdStr)

	//	userId, err = strconv.Atoi(userIdStr)
		if err != nil {
		//	var r global_model.TypeRes

		//	r.Status.ErrCode = 2.0
			//r.Status.ErrDesc = "Malformed Cookie or Header" + global_model.CONST_HTTP_COOKIE_USER_ID_KEY

	//		c.AbortWithStatusJSON(400, r)

			return
		}

		// apiKey

	//	session, err = c.Cookie(global_model.CONST_HTTP_COOKIE_SESSION_KEY)
		if err != nil {
			session = ""
			err = nil
		}

		if len(session) == 0 {
		//	session = c.GetHeader(global_model.CONST_HTTP_HEADER_SESSION_KEY)
		}

		if len(session) == 0 {
	//		var r global_model.TypeRes
		//	r.Status.ErrCode = 2.0
		//	r.Status.ErrDesc = "Please Provide a Cookie or Header " + global_model.CONST_HTTP_COOKIE_SESSION_KEY

//			c.AbortWithStatusJSON(400, r)

			return
		}

		/*** * * ***/

		// checkPass
	//	checkPass, _, _, err = db.Redis.Keys.Sessions.IsValid(redisConn, session, userId)
		if err != nil {
		//	var r global_model.TypeRes
		//	r.Status.ErrCode = global_model.CONST_ERR_CODE_GENERIC
		//	r.Status.ErrDesc = err.Error()
		//	c.AbortWithStatusJSON(500, r)
			return
		}

		/*** * * ***/

		// checkPass
		if !checkPass {
		//	var r global_model.TypeRes
		//	r.Status.ErrCode = global_model.CONST_ERR_CODE_AUTH_UNAUTHORIZED
	//		r.Status.ErrDesc = "Didn't Pass Auth Check"
		//	c.AbortWithStatusJSON(400, r)
			return
		}

	//	c.Next()
	})
/*
	env = global_model.TypeEnv{
		COMPOSE_PROFILES:                       os.Getenv("COMPOSE_PROFILES"),
		TRANSMITTER_FALLOCATE_MB:               os.Getenv("TRANSMITTER_FALLOCATE_MB"),
		TRANSMITTER_FALLOCATE_ENABLE:           os.Getenv("TRANSMITTER_FALLOCATE_ENABLE"),
		SAVE_LAST_HOUR_INTERVAL_PER_SERIAL_SEC: os.Getenv("SAVE_LAST_HOUR_INTERVAL_PER_SERIAL_SEC"),
		TTL_RECORDS_HISTORY_DAYS:               os.Getenv("TTL_RECORDS_HISTORY_DAYS"),
		TTL_ALARMS_HISTORY_DAYS:                os.Getenv("TTL_ALARMS_HISTORY_DAYS"),
	}

	// ctx
	ctx = global_model.TypeCtx{
		GinEngine:            _ginEngine,
		RedisConn:            redisConn,
		CassandraConn:        cassandraConn,
		SqliteConn:           sqliteConn,
		SqliteMiniExtrasConn: sqliteMiniExtrasConn,
		Serial:               serial,
		Env:                  env,
	}

	// views
	views.Trigger(ctx)
*/
	// controllers
//	controllers.Trigger(ctx)
// Initialize alarms module
    alarms.Init(&alarms.AlarmContext{
        CassandraConn: cassandraConn,
        SqliteConn:    sqliteConn,
    })

    // Register alarm routes
    r.POST("/v2/alarms/get", alarms.HandleGetAlarms)
    r.POST("/v2/alarms/set", alarms.HandleSetAlarm)
    r.GET("/v2/alarms/ws", alarms.HandleWebSocket)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/*** * * ***/

	// gin
	// gin : run
	r.Run(":8083")
    	log.Println("API is running on port 8083")
}
