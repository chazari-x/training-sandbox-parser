package sale

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
	typeSale *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			typeSale: regexp.MustCompile(`^(\[SALE]|SALE): *(\{[a-fA-F0-9]{6}})?(.+)$`),
		},
	}
}

// Parse parses sale message and returns struct MessageSale
func (p *Parser) Parse(text string) (*model.MessageSale, error) {
	var message model.MessageSale

	message.Type = model.ChatMessageTypeSale
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeSale.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("sale not found")
	}

	if len(matches) < 3 {
		return nil, fmt.Errorf("message not found: %v", matches)
	}

	message.Message = strings.TrimSpace(matches[3])

	return &message, nil
}
