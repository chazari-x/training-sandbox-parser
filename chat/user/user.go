package user

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
	regexps *regexps
}

type regexps struct {
	typeUser *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typeUser: regexp.MustCompile(`^([a-zA-Z0-9[\]()$@._=-]{3,}) \[(\d+)]$|^\*{2} ([a-zA-Z0-9[\]()$@._=-]{3,}) - UserID: (\d+) ?\|? ?A?F?K? ?(\d+)? ?с?е?к?\.?$`),
		},
	}
}

// Parse parses user message and returns struct MessageUser
func (p *Parser) Parse(text string) (*model.MessageUser, error) {
	var message = model.MessageUser{
		Type:      model.ChatMessageTypeUser,
		Timestamp: int(time.Now().UTC().Unix()),
	}

	matches := p.regexps.typeUser.FindStringSubmatch(text)
	if matches == nil || len(matches) < 4 || (matches[1] == "" && matches[3] == "") {
		return nil, fmt.Errorf("user not found")
	}

	var err error
	if matches[1] != "" {
		message.Nick = matches[1]
		message.ID, err = strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
	} else {
		message.Nick = matches[3]
		message.ID, err = strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}
		if matches[5] != "" {
			message.AFK, err = strconv.Atoi(matches[5])
			if err != nil {
				return nil, err
			}
		}
	}

	return &message, nil
}
