package ask

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
	user    *regexp.Regexp
	message *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			user:    regexp.MustCompile(`([a-zA-Z0-9[\]()$@._=-]{3,})\((\d+)\):{[a-fA-F0-9]{6}}`),
			message: regexp.MustCompile(`\(\d+\):\{[a-fA-F0-9]{6}}(.+)$`),
		},
	}
}

// Parse parses ASK message and returns struct MessageASK
func (p *Parser) Parse(text string) (*model.MessageASK, error) {
	var message model.MessageASK

	message.Type = model.ChatMessageTypeAsk
	message.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.user.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("user not found")
	}

	message.Nick = matches[1]

	userID, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	message.UserID = userID

	matches = p.regexps.message.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
	}

	message.Message = strings.TrimSpace(matches[1])

	return &message, nil
}
