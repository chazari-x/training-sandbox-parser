package pmFromBot

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
	typePMFromBot *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typePMFromBot: regexp.MustCompile(`^\(\( PM к ([a-zA-Z0-9[\]()$@._ =-]{3,}) \((\d+)\): (.+) \)\)( \| AFK (\d+) сек.)?$`),
		},
	}
}

// Parse parses PM from bot message and returns struct MessagePMFromBot
func (p *Parser) Parse(text string) (*model.MessagePMFromBot, error) {
	var message model.MessagePMFromBot

	message.Type = model.ChatMessageTypePMFromBot
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typePMFromBot.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("pm from bot not found")
	}

	message.Nick = strings.ReplaceAll(matches[1], " ", "_")

	userID, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	message.UserID = userID

	message.Message = strings.TrimSpace(matches[3])

	if len(matches) > 5 && matches[5] != "" {
		afk, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, err
		}

		message.AFK = afk
	}

	return &message, nil
}
