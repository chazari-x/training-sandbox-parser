package pmForBot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
	regexps *regexps
}

type regexps struct {
	typePMForBot *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typePMForBot: regexp.MustCompile(`^\(\( PM от ((([a-zA-Z0-9[\]()$@._ =-]{3,}) \((\d+)\))|(Призрак)): (.+) \)\)$`),
		},
	}
}

// Parse parses PM for bot message and returns struct MessagePMForBot
func (p *Parser) Parse(text string) (*model.MessagePMForBot, error) {
	var message model.MessagePMForBot

	message.Type = model.ChatMessageTypePMForBot
	message.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.typePMForBot.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("pm for bot not found")
	}

	if matches[3] != "" {
		message.Nick = matches[3]

		userID, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}

		message.UserID = userID
	} else if matches[5] != "" {
		message.UserID = -1
		message.Nick = matches[5]
	}

	message.Message = strings.TrimSpace(matches[6])

	return &message, nil
}
