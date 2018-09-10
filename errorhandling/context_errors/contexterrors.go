package context_errors

type contextor interface {
	Context() map[string]interface{}
}

func Context(err error) map[string]interface{} {
	var ctx map[string]interface{}
	if e, ok := err.(contextor); ok {
		ctx = e.Context()
	} else {
		ctx = make(map[string]interface{})
	}
	return ctx
}
