package model

type Protocol struct {
	ID          int64
	Name        string
	DefaultPort *int
	Description *string
}