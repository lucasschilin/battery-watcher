# 🔋 Battery Watcher (Go)

![Go](https://img.shields.io/badge/Go-1.24%2B-blue) ![Linux](https://img.shields.io/badge/OS-Linux-yellow) 

## 📌 Description

**Battery Watcher** is a lightweight background service written in Go that monitors laptop battery levels and AC power status on Linux systems and sends desktop notifications when predefined charge thresholds are reached.

The project is intentionally minimal, focused on reliability, low resource usage, and direct interaction with the Linux power subsystem.

---

## ⚙️ Prerequisites

This project relies on standard Linux interfaces and a single external dependency for desktop notifications.

### Required
- Linux (tested on **Pop!_OS / Ubuntu-based distros**)
- Go **1.20+**
- `systemd` (user services)
- Desktop environment with notification support (GNOME, etc.)

### Notification dependency
```bash
sudo apt install libnotify-bin
```

---

## 🧠 What does this project do?

- Reads the **battery charge percentage** directly from the Linux kernel
- Detects whether the **AC charger is connected**
- Sends desktop notifications when:
  - Battery reaches **30%** and the charger is disconnected
  - Battery reaches **80%** and the charger is connected
- Avoids notification spam by triggering alerts only on **state transitions**
- Runs as a **systemd user service**, starting automatically with the user session

---

## 🎯 Motivation

This project was created **for my own machine**.

My laptop (a Dell Inspiron from 2015) does **not support intelligent battery charge management**, such as limiting the charge to 80% via BIOS or firmware.

Because of that, I need to **manually connect and disconnect the charger** to preserve battery health.

This tool exists to:
- Notify me at the right moments
- Avoid constantly checking battery percentage
- Automate a repetitive, error-prone habit

---

## ▶️ How to run

### 1️⃣ Clone the repository
```bash
git clone https://github.com/lucasschilin/battery-watcher.git
cd battery-watcher
```

### 2️⃣ Build the binary
```bash
go build -o bin/battery-watcher
```

### 3️⃣ Move the binary to a user bin directory
```bash
mkdir -p ~/.local/bin
mv bin/battery-watcher ~/.local/bin/
```

---

## ⚙️ Running as a systemd user service

Create the service file:

```ini
[Unit]
Description=Battery Watcher (30% / 80%)
After=graphical-session.target

[Service]
Type=simple
ExecStart=/home/%u/.local/bin/battery-watcher
Restart=always
RestartSec=5
KillMode=process

[Install]
WantedBy=default.target
```

Enable it:

```bash
systemctl --user daemon-reload
systemctl --user enable --now battery-watcher.service
```

---

## 🔍 Monitoring and debugging

Check status:
```bash
systemctl --user status battery-watcher.service
```

View logs:
```bash
journalctl --user -u battery-watcher.service -f
```

---

## 🚧 Scope and future considerations

This project was **not designed to be a full-featured, cross-platform solution**.

It may evolve in the future, but that was not the intention.