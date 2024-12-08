package sads

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
	typeSADS *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typeSADS: regexp.MustCompile(`^\[S-ADS #\d+] \{[a-fA-F0-9]{6}}([a-zA-Z0-9[\]()$@._=-]{3,})\((\d+)\):\{[a-fA-F0-9]{6}} *(.+)$`),
		},
	}
}

// Parse parses SADS message and returns struct MessageSADS
func (p *Parser) Parse(text string) (*model.MessageSADS, error) {
	var message model.MessageSADS

	message.Type = model.ChatMessageTypeSADS
	message.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.typeSADS.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("sads not found")
	}

	message.Nick = matches[1]

	userID, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	message.UserID = userID

	message.Message = strings.TrimSpace(matches[3])

	return &message, nil
}
