package copchase

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
	lobby *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{
		regexps: &regexps{
			lobby: regexp.MustCompile(`#(\d{1,2})\s+(\{[a-zA-Z0-9]{6}}\s*)?((П\W+|В\W+)(\{[a-zA-Z0-9]{6}}))?\s*(\d{1,2}:\d{2})?\s*(\{[a-zA-Z0-9]{6}}\s*)?(\d+(\s-\s\d+)?)\s*(\{[a-zA-Z0-9]{6}}\s*)?(\{[a-zA-Z0-9]{6}}\s*)?(\d)\s/\s(\d)\s*(\{[a-zA-Z0-9]{6}}\s*)?\n`),
		},
	}
}

// Parse parses the text and returns the list of lobbies.
func (p *Parser) Parse(text string) (*model.CopChaseList, error) {
	var data = model.CopChaseList{
		Type:      model.CopChaseListType,
		Timestamp: int(time.Now().UTC().Unix()),
	}

	matches := p.regexps.lobby.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return nil, fmt.Errorf("lobbies not found")
	}

	for _, match := range matches {
		var world model.Lobby
		world.Number, _ = strconv.Atoi(match[1])
		world.Status = strings.TrimSpace(match[4])
		world.Time = strings.TrimSpace(match[6])
		world.Rating = strings.TrimSpace(match[8])
		world.Players, _ = strconv.Atoi(match[12])
		world.MaxPlayers, _ = strconv.Atoi(match[13])

		data.Lobbies = append(data.Lobbies, world)
	}

	return &data, nil
}
