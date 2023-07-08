package main

import (
	"strings"
	"time"
)

type Searcher interface { //интерфейс реализующий метод поиска
	Search(text string) bool
}
type Timestamp struct {
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}

type Document struct {
	ID     int
	UserID int
	Name   string
	Value  string
	Timestamp
}

func (d *Document) Search(text string) bool {
	return strings.Contains(d.Name, text) || strings.Contains(d.Value, text)
}

func NewDocument(userid int, name, value string) *Document {
	return &Document{
		Name:      name,
		UserID:    userid,
		Value:     value,
		Timestamp: NewTimestamp(),
	}
}

type Message struct {
	ID        int
	SenderID  int
	ReceverID int
	Text      string
	Timestamp
}

func (m *Message) Search(text string) bool {
	return strings.Contains(m.Text, text)
}
func NewMessage(senderId, receverId int, text string) *Message {
	return &Message{
		SenderID:  senderId,
		ReceverID: receverId,
		Text:      text,
		Timestamp: NewTimestamp(),
	}
}

func SearchInText(text string, value []Searcher) []Searcher {
	result := make([]Searcher, 0)
	for _, v := range value {
		if v.Search(text) {
			result = append(result, v)
		}
	}
	return result
}
