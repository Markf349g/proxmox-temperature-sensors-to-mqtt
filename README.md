This is an incomplete version of the repository. The repository will be usable once it updates the data. You will not see this message on the final product.

# 📡 Proxmox Temperature Sensors to MQTT 🌡️

### A program for sending temperature metrics of sensors from Proxmox over the MQTT protocol.

---

## 🚀 Features

- 🔍 **Collects temperature metrics** from Proxmox sensors.
- 📤 **Publishes sensor data** via MQTT for easy integration.
- 🛠️ **Configurable & efficient**, suitable for automation.
- ⚡ **Lightweight** and optimized for minimal resource usage.

---

## 🔧 Installation

### Prerequisites
- 🐧 **Linux system** (Debian-based recommended)
- 📦 **Go installed** (`>=1.18`)
- 🐝 **MQTT broker** (e.g., [Mosquitto](https://mosquitto.org/))

### Steps
```sh
git clone https://github.com/username/proxmox-temperature-sensors-to-mqtt.git
cd proxmox-temperature-sensors-to-mqtt
go build -o proxmox-mqtt
