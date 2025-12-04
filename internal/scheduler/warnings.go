package scheduler

import "time"

type RestartWarning struct {
	delay   time.Duration
	message string
}

var RestartWarnings = []RestartWarning{
	{10 * time.Minute, "§cServer will restart in 10 minutes!"},
	{5 * time.Minute, "§cServer will restart in 5 minutes!"},
	{1 * time.Minute, "§cServer will restart in 1 minute!"},
	{30 * time.Second, "§cServer will restart in 30 seconds!"},
	{10 * time.Second, "§cServer will restart in 10 seconds!"},
	{5 * time.Second, "§cServer will restart in 5 seconds!"},
	{4 * time.Second, "§cServer will restart in 4 seconds!"},
	{3 * time.Second, "§cServer will restart in 3 seconds!"},
	{2 * time.Second, "§cServer will restart in 2 seconds!"},
	{1 * time.Second, "§cServer will restart in 1 second!"},
	{0 * time.Second, "§cServer is restarting now!"},
}
