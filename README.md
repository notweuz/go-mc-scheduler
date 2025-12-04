# go-mc-restart
RCON & cron based Minecraft server auto-restart tool written in Go.

## Features
- Automatically restarts Minecraft server at specified intervals.
- Sends warning messages to players before restarting.
- Gracefully stops the server to prevent data loss.
- Configurable settings via a yaml file.

## Future Plans
- [ ] Complete rewrite project and rename it to go-mc-scheduler.
- [ ] Add more scheduling options (e.g., daily, weekly restarts).
- [ ] Add custom schedule tasks