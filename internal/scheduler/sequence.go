package scheduler

import "time"

type RestartSequenceEntry struct {
	delay   time.Duration
	message string
}

var RestartSequence = []RestartSequenceEntry{
	{10 * time.Minute, "say §cServer will restart in 10 minutes!"},
	{5 * time.Minute, "say §cServer will restart in 5 minutes!"},
	{1 * time.Minute, "say §cServer will restart in 1 minute!"},
	{30 * time.Second, "say §cServer will restart in 30 seconds!"},
	{10 * time.Second, "say §cServer will restart in 10 seconds!"},
	{5 * time.Second, "say §cServer will restart in 5 seconds!"},
	{4 * time.Second, "say §cServer will restart in 4 seconds!"},
	{3 * time.Second, "say §cServer will restart in 3 seconds!"},
	{2 * time.Second, "say §cServer will restart in 2 seconds!"},
	{1 * time.Second, "say §cServer will restart in 1 second!"},
	{5 * time.Second, "say §cServer is restarting now!"},
	{0 * time.Second, "kick @a Server is restarting!"},
}
