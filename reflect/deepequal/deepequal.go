package deepequal

type bar struct {
	value string
}

type foo struct {
	name  string
	value int32
	arr   []string
	mp    map[string]bar
}
