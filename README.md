# DISTRIBUTED ALARM & TELEMETRY SYSTEM (DATS)
**MQTT · Go API · Redis · Cassandra · React**

---

## 🧠 DOCKER HOST ARCHITECTURE

### Core Services

- **Node-RED**
  - HTTP: `1880`
  - WebSocket: `9001`

- **Mosquitto (MQTT Broker)**
  - MQTT: `1883`

- **React Frontend**
  - HTTP: `3000`

- **NGINX Gateway**
  - HTTP: `80`
  - HTTPS: `443`

---

### Data Layer

- **Redis**
  - Cache layer
  - Port: `6379`

- **Cassandra**
  - Main distributed database
  - Port: `9042`

- **Go API**
  - Core backend service
  - HTTP: `8083`

- **SQLite**
  - Local fallback / mini-mode storage

---

### Network & Storage

- Network: `alarms-net`
- Persistent Volumes:
  - `cassandra-data`
  - `redis-data`
  - `sqlite-data`
  - `nodered-data`
  - `mosquitto-data`
  - `mosquitto-config`

---

## 🔄 DATA FLOW

### 1. Record Data Flow
