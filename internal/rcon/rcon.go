package rcon

import (
	"go-mc-scheduler/internal/config"
	"time"

	"github.com/gorcon/rcon"
	"github.com/rs/zerolog/log"
)

type Connection struct {
	connection *rcon.Conn
}

func NewConnection(connection *rcon.Conn) *Connection {
	return &Connection{connection: connection}
}

func Connect() (*Connection, error) {
	log.Info().Str("address", config.GetConfig().Rcon.Address).Msg("Connecting to RCON server")
	conn, err := rcon.Dial(
		config.GetConfig().Rcon.Address,
		config.GetConfig().Rcon.Password,
		rcon.SetDialTimeout(10*time.Second),
	)
	return NewConnection(conn), err
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
