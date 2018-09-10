package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/xujihui1985/learninggo/errorhandling/context_errors"
	"github.com/xujihui1985/learninggo/errorhandling/sentinel_errors"
)

func main() {
	queryUser()
	queryUserSecondAttempt()
	queryUserLastAttempt()
	useContextError()
}

/**
很显然这样写非常啰嗦
**/
func queryUser() {
	user, err := errorhandling.GetUser(123)
	switch err := err.(type) {
	case nil:
		fmt.Printf("we got user %d\n", user.ID)
	case *errorhandling.UserNotFoundError:
		fmt.Printf("User %d not found\n", err.GetID())
	default:
		fmt.Println("unknown error")
	}
}

/**
这样写比之前好了很多，但调用的地方耦合了异常类型
**/
func IsUserNotFound(err error) bool {
	if _, ok := err.(*errorhandling.UserNotFoundError); ok {
		return true
	}
	return false
}

func queryUserSecondAttempt() {
	user, err := errorhandling.GetUser(123)
	if err == nil {
		fmt.Printf("we got user %d\n", user.ID)
	}
	if IsUserNotFound(err) {
		fmt.Printf("User %d not found\n", 123)
	}
}

type userNotFound interface {
	UserNotFound() (bool, int32)
}

func IsUserNotFoundLastAttempt(err error) (bool, int32) {
	if e, ok := err.(userNotFound); ok {
		return e.UserNotFound()
	}
	return false, -1
}

func queryUserLastAttempt() {
	user, err := errorhandling.GetUser(123)
	if err == nil {
		fmt.Printf("we got user %d\n", user.ID)
		return
	}
	if ok, id := IsUserNotFoundLastAttempt(err); ok {
		fmt.Printf("User %d not found\n", id)
		return
	}
	fmt.Println("unknown error")
}

func useContextError() {
	user, err := errorhandling.GetUser(123)
	if err != nil {
		logrus.WithFields(logrus.Fields(context_errors.Context(err))).Error(err)
		return
	}
	fmt.Printf("we got user %d\n", user.ID)
}
