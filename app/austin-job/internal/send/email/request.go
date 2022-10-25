package email

import (
	"strings"
)

// Message implements send.Message
type Request struct {
	to      string
	data    string
	subject string
}

func NewMessage(to, subject, data string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		data:    data,
	}
}

func (m *Request) Subject() string {
	return m.subject
}

func (m *Request) To() string {
	return m.to
}

func (m *Request) Data() string {
	return m.data
}

func (m *Request) Content() []byte {
	return []byte(m.data)
}

func (m *Request) TextType(data string) string {
	if strings.Contains(data, "<html>") {
		return "text/html"
	}
	return "text/plain"
}
