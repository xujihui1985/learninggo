package errorhandling

type User struct {
	ID   int32
	Name string
}

func GetUser(id int32) (*User, error) {
	return nil, NewUserNotFoundError(id)
}
