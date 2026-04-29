-- transmitter_serials
CREATE TABLE IF NOT EXISTS transmitter_serials (
                                                   modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                   transmitter_serial VARCHAR(255) NOT NULL,
    UNIQUE (transmitter_serial)
    );

-- devices
CREATE TABLE IF NOT EXISTS devices (
                                       modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       transmitter_serial VARCHAR(255) NOT NULL,
    device VARCHAR(255) NOT NULL,
    UNIQUE (transmitter_serial, device)
    );

-- tags
CREATE TABLE IF NOT EXISTS tags (
                                    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    transmitter_serial VARCHAR(255) NOT NULL,
    tag VARCHAR(255) NOT NULL,
    UNIQUE (transmitter_serial, tag)
    );

-- records_batches
CREATE TABLE IF NOT EXISTS records_batches (
                                               modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                               transmitter_serial VARCHAR(255) NOT NULL,
    hour_unix_sec BIGINT NOT NULL,
    tag VARCHAR(255) NOT NULL,
    lat DOUBLE,
    `long` DOUBLE,
    data LONGBLOB NOT NULL,
    frequency_ms INT NOT NULL,
    UNIQUE (transmitter_serial, hour_unix_sec, tag)
    );

-- records_batches_no_repeat
CREATE TABLE IF NOT EXISTS records_batches_no_repeat (
                                                         modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                         transmitter_serial VARCHAR(255) NOT NULL,
    hour_unix_sec BIGINT NOT NULL,
    tag VARCHAR(255) NOT NULL,
    lat DOUBLE,
    `long` DOUBLE,
    data LONGBLOB NOT NULL,
    frequency_ms INT NOT NULL,
    UNIQUE (transmitter_serial, hour_unix_sec, tag)
    );

-- devices_alarm_groups
CREATE TABLE IF NOT EXISTS devices_alarm_groups (
                                                    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                    transmitter_serial VARCHAR(255) NOT NULL,
    device VARCHAR(255) NOT NULL,
    devices_alarm_group VARCHAR(255) NOT NULL,
    UNIQUE (transmitter_serial, device, devices_alarm_group)
    );

-- alarms
CREATE TABLE IF NOT EXISTS alarms (
                                      id INT AUTO_INCREMENT PRIMARY KEY,
                                      modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      transmitter_serial VARCHAR(255) NOT NULL,
    device VARCHAR(255) NOT NULL,
    devices_alarm_group VARCHAR(255) NOT NULL,
    alarm VARCHAR(255) NOT NULL,
    UNIQUE (
               transmitter_serial,
               device,
               devices_alarm_group,
               alarm
           )
    );

-- alarms_events_history
CREATE TABLE IF NOT EXISTS alarms_events_history (
                                                     modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                     transmitter_serial VARCHAR(255) NOT NULL,
    device VARCHAR(255) NOT NULL,
    devices_alarm_group VARCHAR(255) NOT NULL,
    alarm VARCHAR(255) NOT NULL,
    message TEXT,
    message_alarm_lemma TEXT,
    on_ts_unix_sec BIGINT NOT NULL,
    ack_ts_unix_sec BIGINT,
    off_ts_unix_sec BIGINT,
    UNIQUE (
               transmitter_serial,
               device,
               devices_alarm_group,
               alarm,
               on_ts_unix_sec
           )
    );