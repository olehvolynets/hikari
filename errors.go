package hikari

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
	ErrMissing      = errors.New("value missing")
	ErrNil          = errors.New("is nil")
)

type HikariError struct {
	Err error
}

func (se HikariError) Error() string {
	if se.Err == nil {
		return "hikari: unspecified error"
	}

	return fmt.Sprintf("hikari: %s", se.Err.Error())
}

func Fatal(e error) {
	log.Fatal(HikariError{e})
}

func FatalMsg(msg string, e error) {
	// NOTE: do not wrap in HikariError here
	//  will be done inside Fatal().
	wrappedErr := fmt.Errorf("%s - %w", msg, e)
	Fatal(wrappedErr)
}
