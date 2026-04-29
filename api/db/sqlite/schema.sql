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

