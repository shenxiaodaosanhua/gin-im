package ws

const (
	MESSAGE_TEXT = iota
	MESSAGE_IMAGE
	MESSAGE_AUDIO
)

type Message struct {
	Form    string `json:"form"`
	To      string `json:"to"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}
