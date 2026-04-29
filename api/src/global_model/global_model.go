package global_model

import (
    "database/sql"
        "github.com/gocql/gocql"
            "github.com/gin-gonic/gin"
            	"github.com/redis/go-redis/v9"
	"sync/atomic"
	"time"
)

/*** * * ***/

// Metadata
type TypeMetadata map[string]any

// Primitives

// Primitives : Graph
type TypeGraph struct {
	X []time.Time `json:"x"`
	Y []any       `json:"y"`
}

// Primitives : Value
type TypeValue any

// Primitives : Tag
type TypeTag string

// Primitives : Device
type TypeDevice string

// Primitives : Source
type TypeSource string

// Primitives : Record
type TypeRecord map[TypeTag]TypeValue

// Primitives : User Id
type TypeUserId int

// Primitives : User username
type TypeUsername string

type TypeEnabled int

type TypePass string
type TypePassHash string

type TypePin string

// Not-primitives

// Not-primitives : Tags
type TypeTags []TypeTag

// Not-primitives : Devices
type TypeDevices []TypeDevice

// Not-primitives : Source

// Not-primitives : Source : Sources
type TypeSources []TypeSource

// Not-primitives : TypeUserId : TypeUserIds
type TypeUserIds []TypeUserId

// Not-primitives : TypeUsernames : TypeUsername
type TypeUsernames []TypeUsername

// Not-primitives : Source : Sources : Tags
type TypeSourcesTags map[TypeSource]TypeTags

// Not-primitives : Record

// Not-primitives : Record : Times
type TypeTimesRecord map[time.Time]TypeRecord

// Not-primitives : Record : Sources
type TypeSourcesRecord map[TypeSource]TypeRecord

// Non-primitives : Record : Times : Sources
type TypeSourcesTimesRecord map[TypeSource]TypeTimesRecord

// Not-primitives : Record : Sources : Times
type TypeTimesSourcesRecord map[time.Time]TypeSourcesRecord

/*** * IO * ***/

// Request

type TypeReq struct {
	Data     any            `json:"data"`     // feel free to override
	Metadata map[string]any `json:"metadata"` // override, but better not
}

// Response

type TypeRes struct {
	Data     any            `json:"data"`     // feel free to override
	Metadata map[string]any `json:"metadata"` // override, but better not
	Status   struct {
		ErrCode int32  `json:"errCode"`
		ErrDesc string `json:"errDesc"`
	} `json:"status"` // never override
}

// Response : Graph

type TypeResGraph struct {
	TypeRes
	Data TypeGraph `json:"data"`
}

// Alarms
// Alarms : Events
type TypeAlarmEvent struct {
	// Please note that the order of alarm types is: "on" -> (optionally) "ack" -> "off".
	//
	// Each type implies that the previous state was reached, and we record the time for all types.
	//
	// "on" / "ack" / "off"
	Type string `json:"type"`

	// Identifier
	TransmitterSerial string `json:"transmitterSerial"`
	// Identifier
	Device string `json:"device"`
	// Identifier
	DevicesAlarmGroup string `json:"devicesAlarmGroup"`
	// Identifier
	Alarm string `json:"alarm"`

	// DevicesAlarmGroupNicename string `json:"devicesAlarmGroupNicename"`

	// Message is raw... message. In contrast to "MessageAlarmLemma", here, what you see is what it is.
	//
	// You should not have both "Message" and "MessageAlarmLemma" set at the same time.
	Message string `json:"message"`
	// MessageAlarmLemma is the key for an "alarm lemma".
	//
	// Alarm lemmas are predefined messages.
	//
	// Instead of storing the same message repeatedly, we store it once as an alarm lemma,
	// associate it with a key, and use that key here as MessageAlarmLemma.
	//
	// This isn't always possible—for example, some messages are dynamically generated and vary,
	// which is why we have both "Message" and "MessageAlarmLemma".
	//
	// You should not have both "Message" and "MessageAlarmLemma" set at the same time.
	MessageAlarmLemma string `json:"messageAlarmLemma"`

	// Please note that the order of alarm types is: "on" -> (optionally) "ack" -> "off".
	//
	// Each type implies that the previous state was reached, and we record the time for all types.
	OnTs time.Time `json:"onTs"`
	// Please note that the order of alarm types is: "on" -> (optionally) "ack" -> "off".
	//
	// Each type implies that the previous state was reached, and we record the time for all types.
	AckTs time.Time `json:"ackTs"`
	// Please note that the order of alarm types is: "on" -> (optionally) "ack" -> "off".
	//
	// Each type implies that the previous state was reached, and we record the time for all types.
	OffTs time.Time `json:"offTs"`
}

type TypeAlarmEvents []TypeAlarmEvent

type TypeResAlarmEvent struct {
	TypeAlarmEvent

	OnTsValid  bool `json:"onTsValid"`
	AckTsValid bool `json:"ackTsValid"`
	OffTsValid bool `json:"offTsValid"`

	DevicesAlarmGroupNicename string `json:"devicesAlarmGroupNicename"`
}

type TypeResAlarmEvents []TypeResAlarmEvent

type TypeRecordBatchDatum struct {
	V_float float32 `cql:"v_float"`
	V_bson  []byte  `cql:"v_bson"`
	Repeat  int     `cql:"repeat"`
}

type TypeRecordBatchData []TypeRecordBatchDatum

type TypeTagsRecordBatchData map[string]TypeRecordBatchData

type TypeBsonDocument struct {
	V any `bson:"v"`
}

type TypeRecordBatch struct {
	Synced       time.Time           `cql:"t_synced_ts"`
	Serial       string              `cql:"transmitter_serial"`
	Hour         time.Time           `cql:"hour"`
	Tag          string              `cql:"tag"`
	Lat          float64             `cql:"lat"`
	Long         float64             `cql:"long"`
	Data         TypeRecordBatchData `cql:"data"`
	Frequency_ms int                 `cql:"frequency_ms"`
}

type TypeRecordBatchDatumNoRepeat struct {
	V_float float32 `cql:"v_float"`
	V_bson  []byte  `cql:"v_bson"`
}

type TypeRecordBatchDataNoRepeat []TypeRecordBatchDatumNoRepeat

// type TypeTagsRecordBatchNoRepeatData map[string]TypeRecordBatchNoRepeatData

type TypeRecordBatchNoRepeat struct {
	TypeRecordBatch

	Data TypeRecordBatchDataNoRepeat `cql:"data"`
}

/*** * * ***/

type AtomicFlagBit struct {
	// atomic operations work on hardware level and
	// CPUs generally can't perform atomic operations on single bits,
	// so we stick with int32 even if it's more than we need.
	v int32
}

func (f *AtomicFlagBit) Get() int {
	return (int)(atomic.LoadInt32(&f.v))
}
func (f *AtomicFlagBit) Set1() {
	atomic.StoreInt32(&f.v, 1)
}
func (f *AtomicFlagBit) Set0() {
	atomic.StoreInt32(&f.v, 0)
}

type AtomicFlagBool AtomicFlagBit

func (f *AtomicFlagBool) Get() bool {
	if (int32)(atomic.LoadInt32(&f.v)) == (int32)(1) {
		return true
	}

	return false
}
func (f *AtomicFlagBool) SetTrue() {
	(*AtomicFlagBit)(f).Set1()
}
func (f *AtomicFlagBool) SetFalse() {
	(*AtomicFlagBit)(f).Set0()
}

/*
 * if there is only one row result
 * from controller
 */
type TypeUser struct {
	UserId   TypeUserId   `json:"userId"`
	Username TypeUsername `json:"username"`
}

/*
 * if there is more than one row results
 * from controller (i.e. GetAllUsers func)
 * push them into an array
 */
type TypeUsers []TypeUser

type TypeUserSources struct {
	User    TypeUser    `json:"user"`
	Sources TypeSources `json:"sources"`
}

type TypeUsersSources []TypeUserSources

/*** * * ***/

/**
 * ?  Get users according to posted source.
 */
/*
 * if single row result
 */
type TypeSourceUsers struct {
	Source TypeSource `json:"source"`
	Users  TypeUsers  `json:"users"`
}

/*
 * if there is more than one row
 * push them into an array
 */
type TypeSourcesUsers []TypeSourceUsers

// *** Delete source user *** //
type TypeDeleteSourceUser struct {
	UserId TypeUserId `json:"userId"`
	Source TypeSource `json:"sourceSerial"`
}
type TypeResUserId struct {
	UserId TypeUserId `json:"userId"`
}
// ========== ENVIRONMENT ==========
type TypeEnv struct {
    COMPOSE_PROFILES                       string
    TRANSMITTER_FALLOCATE_MB               string
    TRANSMITTER_FALLOCATE_ENABLE           string
    SAVE_LAST_HOUR_INTERVAL_PER_SERIAL_SEC string
    TTL_RECORDS_HISTORY_DAYS               string
    TTL_ALARMS_HISTORY_DAYS                string
}

// ========== CONTEXT ==========
type TypeCtx struct {
    GinEngine            *gin.Engine
    RedisConn            *redis.Client
    CassandraConn        *gocql.Session
    SqliteConn           *sql.DB
    SqliteMiniExtrasConn *sql.DB
    Serial               string
    Env                  TypeEnv
}