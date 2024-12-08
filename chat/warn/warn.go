package warn

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
	typeWarn       *regexp.Regexp
	typeWarnReason *regexp.Regexp
	typeWarnTime   *regexp.Regexp
	typeWarnUntil  *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typeWarn:       regexp.MustCompile(`^(Модератор|Андроид) ([a-zA-Z0-9[\]()$@._=-]{3,}) ([а-я ]+) ([a-zA-Z0-9[\]()$@._=-]{3,}) ?н?а? ?(\d+ ч)?\..*$`),
			typeWarnReason: regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}.*\. Причина: (.+)$`),
			typeWarnTime:   regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,} на (\d+ .+)\..*$`),
			typeWarnUntil:  regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}. До (\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2})`),
		},
	}
}

// Parse parses warn message and returns struct MessageWarn
func (p *Parser) Parse(text string) (*model.MessageWarn, error) {
	var message model.MessageWarn

	message.Type = model.ChatMessageTypeWarn
	message.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.typeWarn.FindStringSubmatch(text)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("warn not found")
	}

	message.Moderator = matches[2]

	message.Punishment = matches[3]

	message.User = matches[4]

	if message.Moderator == "" || message.User == "" || message.Punishment == "" {
		return nil, fmt.Errorf("parse error")
	}

	matches = p.regexps.typeWarnTime.FindStringSubmatch(text)
	if matches != nil && len(matches) > 2 {
		message.Time = matches[2]
	}

	matches = p.regexps.typeWarnReason.FindStringSubmatch(text)
	if matches != nil && len(matches) > 2 {
		message.Reason = strings.TrimSpace(matches[2])
	}

	matches = p.regexps.typeWarnUntil.FindStringSubmatch(text)
	if matches != nil && len(matches) > 2 {
		message.Until = matches[2]
	}

	return &message, nil
}
