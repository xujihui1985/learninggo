package errorhandling

import (
	"fmt"
)

var (
	ctxUserID = "userID"

	emptyUserID = int32(-1)
)

type UserNotFoundError struct {
	ctx map[string]interface{}
}

func NewUserNotFoundError(userID int32) *UserNotFoundError {
	err := &UserNotFoundError{}
	err.SetID(userID)
	return err
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("can not found user %d", e.GetID())
}

func (e *UserNotFoundError) UserNotFound() (bool, int32) {
	return true, e.GetID()
}

func (e *UserNotFoundError) Context() map[string]interface{} {
	if e.ctx == nil {
		e.ctx = make(map[string]interface{})
	}
	return e.ctx
}

func (e *UserNotFoundError) GetID() int32 {
	ctx := e.Context()
	if userID, ok := ctx[ctxUserID]; ok {
		return userID.(int32)
	}
	return emptyUserID
}

func (e *UserNotFoundError) SetID(id int32) {
	ctx := e.Context()
	ctx[ctxUserID] = id
}
