package wraperrors

import (
	"fmt"
)

type authorizationError struct {
	operation string
	err       error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

func (e *authorizationError) Cause() error {
	return e.err
}
