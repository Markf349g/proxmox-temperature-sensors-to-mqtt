# 📡 Proxmox Temperature Sensors to MQTT 🌡️

### 🚀 **Easily Send Proxmox Sensor Temperatures via MQTT!**

> **Note**: This is an unfinished version of the repository. It will become fully functional after data updates. This note will be removed in the final version.

---

## ✨ **Features**

- 🔍 **Collect temperature data** from Proxmox sensors.
- 📤 **Publish data** via MQTT for seamless integration.
- 🛠️ **Flexible configuration**, perfect for automation.
- ⚡ **Written in pure Go**, producing portable static binaries.

---

## 🛠️ **Installation**

### 📋 **Requirements**
- 🌐 **Internet connection**
- 🖥️ **Proxmox server**
- 📦 **Willingness to follow instructions**

### 💻 **Operating System Options**
Go is highly portable:
- Works on any OS supported by Go.
- Build the program for any supported OS.
- For Proxmox, an **LXC container** is recommended.

### 📥 **Clone the Repository**
Grab the project from GitHub:
```sh
git clone https://github.com/Markf349g/proxmox-temperature-sensors-to-mqtt.git
```

### 📦 **Install dependencies**
```sh
go mod tidy
```

### 🏗️ **Build the Application**
Compile it with this command:

Local Compile:
```sh
go build -v ./...
```
Cross Compile:
- **On Unix-like systems (Linux, macOS, etc.):**

  ```sh
  env GOOS=<OS> GOARCH=<ARCH> go build -v ./...
  ```

- **On Windows (Command Prompt):**

  ```sh
  set GOOS=<OS> && set GOARCH=<ARCH> && go build -v ./...
  ```

- **In PowerShell (any OS):**

  ```powershell
  $env:GOOS=<OS>; $env:GOARCH=<ARCH>; go build -v ./...
  ```

**Important:** Replace **<OS>** and **<ARCH>** with actual values. For example:
- For Linux on AMD64: `GOOS=linux GOARCH=amd64`
- For Windows on ARM64: `GOOS=windows GOARCH=arm64`

For a full list of supported values, check the [Go documentation](https://go.dev/doc/install/source#environment).

---

## ⚙️ **Configuration**

Customize the config file for your setup. Here’s an example with explanations:

### 📝 **Configuration Example**
```json
{
    "SSH": {
        "host": "127.0.0.1:22",
        "_host": "Required: Specify the SSH server host [127.0.0.1:22]",
        "username": "root",
        "_username": "Required: Specify the SSH username [root]",
        "password": "12345",
        "_password": "Required: Specify the SSH password [12345]",
        "delay": "1000",
        "_delay": "Required: Specify delay between errors in ms [1000]"
    },
    "DATA": {
        "prefix": "/homeassistant/sensor/proxmox_system",
        "_prefix": "Optional: Specify the data prefix [/homeassistant/sensor/proxmox_system]",
        "delay": "1",
        "_delay": "Required: Specify delay between data sends in minutes [1]",
        "difference": "5",
        "_difference": "Required: Specify temperature difference [5]"
    },
    "MQTT": {
        "host": "tcp://127.0.0.1:1883",
        "_host": "Required: Specify the MQTT server host [tcp://127.0.0.1:1883]",
        "client-id": "MQTT-Telemetry[proxmox-system]",
        "_client-id": "Required: Specify the MQTT client ID [MQTT-Telemetry[proxmox-system]]",
        "username": "user",
        "_username": "Required: Specify the MQTT username [user]",
        "password": "12345",
        "_password": "Required: Specify the MQTT password [12345]",
        "delay": "1000",
        "_delay": "Required: Specify delay before disconnect in ms [1000]"
    }
}
```

1. Open the config file.
2. Replace the default values with your own settings.
3. Save the file before running the program.

---

## 🚀 **Running the Program**

### 🐧 **Unix Systems**
```sh
./proxmox-temperature-sensors-to-mqtt
```

### 🖼️ **Windows 10+**
```sh
proxmox-temperature-sensors-to-mqtt
```

### ⚡ **Run Directly with Go**
```sh
go run .
```

---

### 📝 **StartUp**
### Windows 10+:
1. Right-click the executable file. 
2. In the menu, select *New* → *Shortcut*.
3. Copy the shortcut to `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup`
4. ***Reboot the system***

### Linux:
#### Cron(easy, **not recommended**)
1. Open Cron in edit mode
```sh
crontab -e
```
2. Add this command in the end of the file
```sh
@reboot </my/path/to/app>
```
**Important:** Replace **</my/path/to/app>** with actual value. For example:
- For *proxmox-temperature-sensors-to-mqtt* in */home/root/proxmox-temperature-sensors-to-mqtt*: `/home/user/app/proxmox-temperature-sensors-to-mqtt`
- For *whale* in */turtle/dragon/unicorn*: `/turtle/dragon/unicorn/whale`
3. Save and close the file
- For **GNU nano**: 
Save:  `Ctrl+S`
Close: `Ctrl+X`
- For **Vim**: 
Save:  `:w`
Close: `:q`
4. ***Reboot the system***

#### Systemd(difficult, **recommended**)
1. Create and edit the file */etc/systemd/system/proxmox-temperature-sensors-to-mqtt.service*
```sh
sudo nano /etc/systemd/system/proxmox-temperature-sensors-to-mqtt.service
```
2. Edit the file
```
[Unit]
Description=Proxmox Temperature Sensors to MQTT
After=network.target

[Service]
ExecStart=</my/path/to/app>
Restart=always
User=<user>
WorkingDirectory=</my/path/to>

[Install]
WantedBy=multi-user.target
```
**Important:** Replace **</my/path/to/app>**, **</my/path/to>**, and **<user>** with actual value. For example:
##### For **</my/path/to/app>**(path to the application):
- For *proxmox-temperature-sensors-to-mqtt* in */home/root/proxmox-temperature-sensors-to-mqtt*: `/home/user/app/proxmox-temperature-sensors-to-mqtt`
- For *whale* in */turtle/dragon/unicorn*: `/turtle/dragon/unicorn/whale`
##### For **</my/path/to>**(parent directory of the application):
- For *proxmox-temperature-sensors-to-mqtt* in */home/root/proxmox-temperature-sensors-to-mqtt*: `/home/user/app`
- For *whale* in */turtle/dragon/unicorn*: `/turtle/dragon/unicorn`
##### For **<user>**(your user in the system):
- For *root*: `root`
- For *user*: `user`
3. Save and close the file
- For **GNU nano**: 
Save:  `Ctrl+S`
Close: `Ctrl+X`
- For **Vim**: 
Save:  `:w`
Close: `:q`
4. ***Reboot the system***
---