# go-mc-scheduler

A simple task scheduler for Minecraft servers using RCON.

## Features

- Schedule tasks with cron syntax
- Execute server commands via RCON
- Multi-step jobs with delays

## Quick Start

### Configuration

Create `applcation.yml`:

```yaml
go_mc_scheduler:
  rcon:
    address: "localhost:25575"
    password: "your_password"
  scheduler:
    timezone: "UTC"
    
    jobs:
      - name: "Daily Restart"
        cron: "0 0 * * *"
        steps:
          - execute: "say Server restarting in 5 minutes"
          - wait: "5m"
          - execute: "save-all"
          - execute: "stop"
```

## Configuration

### RCON Setup

Enable RCON in your `server.properties`:

```properties
enable-rcon=true
rcon.port=25575
rcon.password=your_password
```

### Job Steps

- `execute: "command"` - Run a server command
- `wait: "5m"` - Wait (supports: `s` seconds, `m` minutes, `h` hours)

### Cron Format

```
* * * * *
│ │ │ │ │
│ │ │ │ └─ Day of week (0-6)
│ │ │ └─── Month (1-12)
│ │ └───── Day of month (1-31)
│ └─────── Hour (0-23)
└───────── Minute (0-59)
```

**Examples:**
- `0 0 * * *` - Daily at midnight
- `0 */6 * * *` - Every 6 hours
- `0 4 * * 1` - Every Monday at 4 AM

## Example: Auto-Restart

```yaml
jobs:
  - name: "Auto Restart"
    cron: "0 0 * * *"
    steps:
      - execute: "say §cServer restarting in 10 minutes!"
      - wait: "5m"
      - execute: "say §cServer restarting in 5 minutes!"
      - wait: "4m"
      - execute: "say §cServer restarting in 1 minute!"
      - wait: "55s"
      - execute: "say §cServer restarting in 5 seconds!"
      - wait: "5s"
      - execute: "save-all"
      - execute: "kick @a §cRestarting, reconnect in a moment"
      - execute: "stop"
```

> **Note:** Configure your server to auto-restart (Docker with `restart: always` or systemd)