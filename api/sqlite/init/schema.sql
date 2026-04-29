--- - Users - ---
--- Users : Users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    username TEXT,
    password TEXT,
    --
    enabled_bool INTEGER
);

--- Users : Pin (1-1)
CREATE TABLE IF NOT EXISTS users_pin (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    user_id INTEGER,
    --
    pin TEXT,
    pin_enabled_bool INTEGER
);

--- - Sources - ---
-- Sources : Sources (look-up)
CREATE TABLE IF NOT EXISTS sources (serial TEXT PRIMARY KEY NOT NULL);

--- Permissions : User : Source (1-1)
CREATE TABLE IF NOT EXISTS permissions_user_source (
    modified TEXT DEFAULT (datetime('now')),
    --
    user_id INTEGER NOT NULL,
    source_serial TEXT NOT NULL,
    --
    read INTEGER DEFAULT 0 NOT NULL,
    write INTEGER DEFAULT 0 NOT NULL,
    ---
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (source_serial) REFERENCES sources (serial),
    ---
    PRIMARY KEY (user_id, source_serial)
);

CREATE TABLE IF NOT EXISTS assets (
    filename TEXT NOT NULL,
    --
    created TEXT DEFAULT (datetime('now')),
    --
    creator_user_id INTEGER NOT NULL,
    ---
    FOREIGN KEY (creator_user_id) REFERENCES users (id),
    ---
    PRIMARY KEY (filename)
);

CREATE TABLE IF NOT EXISTS permissions_user_asset (
    user_id INTEGER NOT NULL,
    --
    modified TEXT DEFAULT (datetime('now')), -- added after dev-rel-15
    --
    assets_filename TEXT NOT NULL, -- typo, should had been "asset_filename"
    ---
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (assets_filename) REFERENCES assets (filename),
    ---
    PRIMARY KEY (user_id, assets_filename)
);

--- - Meta - ---
-- Scope :
--  If both entity_id and entity_table are NULL, it indicates a 'global metadata' state.
--  If only entity_id is NULL, it indicates a 'global' state for the specified entity_table.
--  If only entity_table is NULL, it indicates a 'global' state for the specified entity_id.
--- Meta : Meta
CREATE TABLE IF NOT EXISTS meta (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    entity_id INTEGER,
    entity_table TEXT,
    --
    name TEXT,
    value TEXT
);

--- - Denormalized - ---
--- - Denormalized : Alarms - ---
--- - Denormalized : Alarms : Events - ---
--- - Denormalized : Alarms : Events : Live - ---
CREATE TABLE IF NOT EXISTS denormalized_alarms_events_live (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    source_serial TEXT, -- transmitter_serial
    device TEXT NOT NULL,
    devices_alarm_group TEXT NOT NULL,
    --
    alarm TEXT NOT NULL,
    --
    message TEXT,
    message_alarm_lemma TEXT,
    --
    on_ts TIMESTAMP NOT NULL,
    ack_ts TIMESTAMP,
    --
    UNIQUE (source_serial, device, devices_alarm_group, alarm)
);

CREATE TABLE IF NOT EXISTS denormalized_alarms_events_live_last_modified (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    source_serial TEXT NOT NULL, -- transmitter_serial
    device TEXT,
    devices_alarm_group TEXT,
    --
    alarm TEXT,
    --
    last_modified INTEGER,
    --
    UNIQUE (source_serial, device, devices_alarm_group, alarm)
);

CREATE INDEX IF NOT EXISTS
    idx_denormalized_alarms_events_live_last_modified_by_alarm
ON
    denormalized_alarms_events_live_last_modified
    (
        source_serial,
        device,
        devices_alarm_group,
        alarm,
        last_modified DESC
    )
;

CREATE INDEX IF NOT EXISTS
    idx_denormalized_alarms_events_live_last_modified_by_device
ON
    denormalized_alarms_events_live_last_modified
    (
        source_serial,
        device,
        last_modified DESC
    )
;

CREATE INDEX IF NOT EXISTS
    idx_denormalized_alarms_events_live_last_modified_by_serial
ON
    denormalized_alarms_events_live_last_modified
    (
        source_serial,
        last_modified DESC
    )
;

CREATE TABLE IF NOT EXISTS denormalized_devices_alarm_groups_nicenames (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    source_serial INTEGER, -- transmitter_serial
    device TEXT NOT NULL,
    devices_alarm_group TEXT,
    --
    nicename TEXT
);
