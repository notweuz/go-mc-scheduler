package rcon

import (
	"gorestart-minecraft/internal/config"

	"github.com/gorcon/rcon"
	"github.com/rs/zerolog/log"
)

type RCONClient struct {
	Rcon   *rcon.Conn
	Config *config.RconConfig
}

func SetupRCONClient(config *config.RconConfig) {
	Client = &RCONClient{
		Config: config,
	}
}

var Client *RCONClient

func (rc *RCONClient) Connect() error {
	log.Printf("Connecting to %s\n", rc.Config.Address)
	conn, err := rcon.Dial(rc.Config.Address, rc.Config.Password)
	rc.Rcon = conn
	return err
}

func (rc *RCONClient) Execute(command string) (string, error) {
	response, err := rc.Rcon.Execute(command)
	log.Info().Str("command", command).Str("response", response).Msg("Executed RCON command")
	return response, err
}

func (rc *RCONClient) Close() error {
	err := rc.Rcon.Close()
	log.Info().Msg("Closing RCON connection")
	return err
}
