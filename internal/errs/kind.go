package errs

type Kind uint8

const (
    Internal Kind = iota
    InvalidInput
    Unauthenticated
	MethodNotAllowed
	BadRequest
    Forbidden
    NotFound
    Conflict
)