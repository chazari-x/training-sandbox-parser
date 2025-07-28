package sale

import (
	"fmt"
	"regexp"
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
	regexps *regexps
}

type regexps struct {
	message *regexp.Regexp
	content *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			content: regexp.MustCompile(`^"(.+)": ((\{[a-fA-F0-9]{6}}(\d+р\.)\{[a-fA-F0-9]{6}} \| \{[a-fA-F0-9]{6}}(-?\d+%)\{[a-fA-F0-9]{6}} \| Прошлая цена: \{[a-fA-F0-9]{6}}(\d+р\.))|(\{[a-fA-F0-9]{6}}Скидка: (-?\d+%)))$`),
			message: regexp.MustCompile(`^\[SALE]: \{[a-fA-F0-9]{6}} *(.+)$`),
		},
	}
}

// Parse parses sale message and returns struct MessageSale
func (p *Parser) Parse(text string) (*model.MessageSale, error) {
	var message model.MessageSale

	message.Type = model.ChatMessageTypeSale
	message.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.message.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("sale not found in message: %v", text)
	}

	if len(matches) < 2 {
		return nil, fmt.Errorf("message not found: %v", text)
	}

	message.Message = matches[1]

	matches = p.regexps.content.FindStringSubmatch(message.Message)
	if matches == nil && message.Message == "" {
		if message.Message == "" {
			return nil, fmt.Errorf("sale content not found in message: %v", text)
		}
		return &message, nil
	}

	if len(matches) < 9 {
		if message.Message == "" {
			return nil, fmt.Errorf("sale content not found in message: %v", text)
		}
		return &message, nil
	}

	message.Message = matches[1]

	if message.Message == "" {
		return nil, fmt.Errorf("product name not found in sale message: %v", text)
	}

	message.Price = matches[4]
	message.Discount = matches[5]
	message.OldPrice = matches[6]

	if message.Discount == "" {
		message.Discount = matches[8]
	}

	return &message, nil
}
