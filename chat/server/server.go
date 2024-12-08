package server

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
	regexps *regexps
}

type regexps struct {
	typeServer *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typeServer: regexp.MustCompile(`^(\[SERVER]|SERVER): *(\{[a-fA-F0-9]{6}})?(.+)$`),
		},
	}
}

// Parse parses server message and returns struct MessageServer
func (p *Parser) Parse(text string) (*model.MessageServer, error) {
	var message model.MessageServer

	message.Type = model.ChatMessageTypeServer
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeServer.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("server not found")
	}

	if len(matches) < 3 {
		return nil, fmt.Errorf("message not found: %v", matches)
	}

	message.Message = strings.TrimSpace(matches[3])

	return &message, nil
}
