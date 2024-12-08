package any

import (
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

// Parse parses any message and returns struct MessageAny
func (p *Parser) Parse(text string) (*model.MessageAny, error) {
	var message model.MessageAny

	message.Type = model.ChatMessageTypeAny

	message.Message = text

	message.Timestamp = time.Now().UTC().Unix()

	return &message, nil
}
