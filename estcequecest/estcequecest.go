package estcequecest

import (
	"fmt"
	"time"
)

// types for the json configuration data
type estcequecestData struct {
	Title          string         `json:"title"`
	Defaultmessage string         `json:"default"`
	Messages       []messagesData `json:"messages"`
}
type messagesData struct {
	Interval string            `json:"interval"`
	Timeline map[string]string `json:"timeline"`
}

// object that holds everything
type Estcequecest struct {
	title          string
	defaultMessage string
	messages       []*messages
}

func NewEstcequecest(data estcequecestData) *Estcequecest {
	e := new(Estcequecest)
	e.title = data.Title
	e.defaultMessage = data.Defaultmessage
	e.messages = make([]*messages, 0)
	for _, mess := range data.Messages {
		e.messages = append(e.messages, newMessages(mess))
	}
	return e
}

func (e *Estcequecest) String() string {
	s := fmt.Sprintf("Title:'%s', default:'%s'\n", e.title, e.defaultMessage)
	for _, mess := range e.messages {
		s += mess.String()
	}
	return s
}

func (e *Estcequecest) getMessageAtTime(t time.Time) (title string, message string) {
	for _, message := range e.messages {
		if mess, ok := message.getMessageAtTime(t); ok && mess != "" {
			return e.title, mess
		}
	}
	if e.defaultMessage != "" {
		return e.title, e.defaultMessage
	}
	return "", ""
}
