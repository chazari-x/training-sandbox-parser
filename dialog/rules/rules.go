package rules

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
	rule *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			rule: regexp.MustCompile(`((\{[a-zA-Z0-9]{6}})? *(\d.\d*) *(\{[a-zA-Z0-9]{6}})? *(.+)|.+)\n`),
		},
	}
}

// Parse parses the text and returns the list of rules.
func (p *Parser) Parse(text string) (*model.Rules, error) {
	var data = model.Rules{
		Type:      model.DialogRulesType,
		Timestamp: int(time.Now().UTC().Unix()),
	}

	matches := p.regexps.rule.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return nil, fmt.Errorf("rules not found")
	}

	for _, match := range matches {
		var rule model.Rule
		rule.Point = strings.TrimSpace(match[3])

		if match[5]+match[4]+match[3]+match[2] == "" {
			rule.Text = strings.TrimSpace(match[1])
		} else {
			rule.Text = strings.TrimSpace(match[5])
		}

		data.Rules = append(data.Rules, rule)
	}

	return &data, nil
}
