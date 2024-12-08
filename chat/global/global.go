package global

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
	star    *regexp.Regexp
	user    *regexp.Regexp
	prefix  *regexp.Regexp
	message *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			star:    regexp.MustCompile(`^(\{[a-fA-F0-9]{6}}.)`),
			user:    regexp.MustCompile(`([a-zA-Z0-9[\]()$@._=-]{3,})\((\d+)\):{[a-fA-F0-9]{6}}`),
			prefix:  regexp.MustCompile(` ({[a-fA-F0-9]{6}}.+){FFFFFF}`),
			message: regexp.MustCompile(`\(\d+\):\{[a-fA-F0-9]{6}}(.+)$`),
		},
	}
}

// Parse parses global message and returns struct MessageGlobal
func (p *Parser) Parse(text string) (*model.MessageGlobal, error) {
	var message model.MessageGlobal

	message.Type = model.ChatMessageTypeGlobal
	message.Timestamp = int(time.Now().UTC().Unix())

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

	matches = p.regexps.star.FindStringSubmatch(text)
	if matches != nil {
		message.Star = matches[1]
	}

	matches = p.regexps.prefix.FindStringSubmatch(text)
	if matches != nil {
		message.Prefix = matches[1]
	}

	return &message, nil
}
