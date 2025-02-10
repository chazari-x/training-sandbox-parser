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

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			lobby: regexp.MustCompile(`#(\d{1,2})\s+(\{[a-zA-Z0-9]{6}}\s*)?((П\W+|В\W+)(\{[a-zA-Z0-9]{6}}))?\s*(\d{1,2}:\d{2})?\s*(\{[a-zA-Z0-9]{6}}\s*(\{[a-zA-Z0-9]{6}})?)?(\d+(\s-\s\d+)?)\s*(\{[a-zA-Z0-9]{6}}\s*)?(\{[a-zA-Z0-9]{6}}\s*)?(\d)\s/\s(\d)\s*(\{[a-zA-Z0-9]{6}}\s*)?\n`),
		},
	}
}

// Parse parses the text and returns the list of lobbies.
func (p *Parser) Parse(text string) (*model.CopChaseList, error) {
	data := model.CopChaseList{
		Type:      model.DialogCopChaseListType,
		Timestamp: time.Now().UTC().Unix(),
	}

	matches := p.regexps.lobby.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return nil, fmt.Errorf("lobbies not found")
	}

	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		players, _ := strconv.Atoi(match[13])
		maxPlayers, _ := strconv.Atoi(match[14])

		world := model.Lobby{
			Number:     number,
			Status:     strings.TrimSpace(match[4]),
			Time:       strings.TrimSpace(match[6]),
			Rating:     strings.TrimSpace(match[9]),
			Players:    players,
			MaxPlayers: maxPlayers,
		}

		data.Lobbies = append(data.Lobbies, world)
	}

	return &data, nil
}
