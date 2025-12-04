package rcon

import (
	"go-mc-scheduler/internal/config"

	"github.com/gorcon/rcon"
	"github.com/rs/zerolog/log"
)

type Connection struct {
	connection *rcon.Conn
}

func Connect() (*Connection, error) {
	log.Printf("Connecting to %s\n", config.GetConfig().Rcon.Address)
	conn, err := rcon.Dial(config.GetConfig().Rcon.Address, config.GetConfig().Rcon.Password)
	return &Connection{connection: conn}, err
}

func (c *Connection) Execute(command string) (string, error) {
	response, err := c.connection.Execute(command)
	log.Info().Str("command", command).Str("response", response).Msg("Executed RCON command")
	return response, err
}

func (c *Connection) Close() error {
	err := c.connection.Close()
	log.Info().Msg("Closing RCON connection")
	return err
}
