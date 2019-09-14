package models

type Command struct {
	Timestamp int64
	Type      string
	Nickname  string
	Code      string
	Msg       string
	Value     int
}
