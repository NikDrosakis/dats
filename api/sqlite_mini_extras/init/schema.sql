--- - Mini - ---
---
--- transmitter serials
CREATE TABLE IF NOT EXISTS transmitter_serials (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    ---
    UNIQUE (transmitter_serial)
);

--- devices
CREATE TABLE IF NOT EXISTS devices (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    --
    device TEXT NOT NULL,
    ---
    UNIQUE (transmitter_serial, device)
);

--- tags
CREATE TABLE IF NOT EXISTS tags (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    --
    tag TEXT NOT NULL,
    ---
    UNIQUE (transmitter_serial, tag)
);

--- records batches
CREATE TABLE IF NOT EXISTS records_batches (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    --
    -- Keep "hours"... hours.
    --  E.g., '2025-01-15 14:00:00.000000+0000',
    --  not '2025-01-15 14:34:45.524000+0000'
    hour_unix_sec INTEGER NOT NULL,
    --
    tag TEXT NOT NULL,
    --
    lat REAL, -- "8-byte IEEE float" => float64 / double
    long REAL, -- "8-byte IEEE float" => float64 / double
    --
    data BLOB NOT NULL,
    --
    frequency_ms INT NOT NULL,
    ---
    UNIQUE (transmitter_serial, hour_unix_sec, tag)
);

CREATE TABLE IF NOT EXISTS records_batches_no_repeat (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    --
    -- Keep "hours"... hours.
    --  E.g., '2025-01-15 14:00:00.000000+0000',
    --  not '2025-01-15 14:34:45.524000+0000'
    hour_unix_sec INTEGER NOT NULL,
    --
    tag TEXT NOT NULL,
    --
    lat REAL, -- "8-byte IEEE float" => float64 / double
    long REAL, -- "8-byte IEEE float" => float64 / double
    --
    data BLOB NOT NULL,
    --
    frequency_ms INT NOT NULL,
    ---
    UNIQUE (transmitter_serial, hour_unix_sec, tag)
);

--- devices : alarm groups
CREATE TABLE IF NOT EXISTS devices_alarm_groups (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    device TEXT NOT NULL,
    --
    devices_alarm_group TEXT NOT NULL,
    ---
    UNIQUE (transmitter_serial, device, devices_alarm_group)
);

--- devices : alarms
CREATE TABLE IF NOT EXISTS alarms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    device TEXT NOT NULL,
    devices_alarm_group TEXT NOT NULL,
    --
    alarm TEXT NOT NULL,
    ---
    UNIQUE (
        transmitter_serial,
        device,
        devices_alarm_group,
        --
        alarm
    )
);

--- devices : alarms : events
CREATE TABLE IF NOT EXISTS alarms_events_history (
    modified TEXT DEFAULT (datetime('now')),
    --
    transmitter_serial TEXT NOT NULL,
    device TEXT NOT NULL,
    devices_alarm_group TEXT NOT NULL,
    --
    alarm TEXT NOT NULL,
    --
    message TEXT,
    message_alarm_lemma TEXT,
    --
    on_ts_unix_sec INTEGER NOT NULL,
    ack_ts_unix_sec INTEGER,
    off_ts_unix_sec INTEGER,
    ---
    UNIQUE (
        transmitter_serial,
        device,
        devices_alarm_group,
        --
        alarm,
        --
        on_ts_unix_sec
    )
);
