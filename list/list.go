package list

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
	world *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{
		regexps: &regexps{
			world: regexp.MustCompile(`(.+)\s+(\d+)\s+(\[S]|\[ ?S&SMP ?])?\s*\n`),
		},
	}
}

// Parse parses the text and returns the list of worlds.
func (p *Parser) Parse(text string) (*model.WorldsList, error) {
	var data = model.WorldsList{
		Type:      model.WorldsListType,
		Timestamp: int(time.Now().UTC().Unix()),
	}

	matches := p.regexps.world.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return nil, fmt.Errorf("worlds not found")
	}

	for _, match := range matches {
		var world model.World
		world.Name = strings.TrimSpace(match[1])
		world.Players, _ = strconv.Atoi(match[2])
		if len(match) < 4 {
			world.Static = false
		} else {
			world.Static = match[3] == "[S]"
			world.SSMP = strings.Contains(match[3], "S&SMP")
		}
		data.Worlds = append(data.Worlds, world)
	}

	return &data, nil
}
