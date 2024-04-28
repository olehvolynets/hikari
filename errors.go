package sylphy

import (
	"fmt"
	"log"
)

type SylphyError struct {
	Err error
}

func (se SylphyError) Error() string {
	if se.Err == nil {
		return "sylphy: unspecified error"
	}

	return fmt.Sprintf("sylphy: %s", se.Err.Error())
}

func Fatal(e error) {
	log.Fatal(SylphyError{e})
}

func FatalMsg(msg string, e error) {
	// NOTE: do not wrap in SylphyError here
	//  will be done inside Fatal().
	wrappedErr := fmt.Errorf("%s - %w", msg, e)
	Fatal(wrappedErr)
}
