package models

import "time"

type Message struct {
	ID        int64
	ChatID    int64
	Sender    string
	Text      string
	Timestamp time.Time
}
