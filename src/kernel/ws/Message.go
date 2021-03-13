package ws

const (
	MESSAGE_TEXT = iota
	MESSAGE_IMAGE
	MESSAGE_AUDIO
)

type Message struct {
	Form    string
	To      string
	Content string
	Type    int
}
