package chat

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chazari-x/training-sandbox-parser/model"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
	regexps *regexps
}

type regexps struct {
	star           *regexp.Regexp
	user           *regexp.Regexp
	prefix         *regexp.Regexp
	message        *regexp.Regexp
	typeWarn       *regexp.Regexp
	typeWarnReason *regexp.Regexp
	typeWarnTime   *regexp.Regexp
	typeWarnUntil  *regexp.Regexp
	typeServer     *regexp.Regexp
	typeSale       *regexp.Regexp
	typePMForBot   *regexp.Regexp
	typePMFromBot  *regexp.Regexp
	typeSADS       *regexp.Regexp
	typeUser       *regexp.Regexp

	ChatTypeUser      *regexp.Regexp
	chatTypeSADS      *regexp.Regexp
	chatTypeWorld     *regexp.Regexp
	chatTypeGlobal    *regexp.Regexp
	chatTypeADS       *regexp.Regexp
	chatTypeASK       *regexp.Regexp
	chatTypeWarn      *regexp.Regexp
	chatTypeServer    *regexp.Regexp
	chatTypeSale      *regexp.Regexp
	chatTypePMForBot  *regexp.Regexp
	chatTypePMFromBot *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{
		regexps: &regexps{
			star:           regexp.MustCompile(`^(\{[a-fA-F0-9]{6}}.)`),
			user:           regexp.MustCompile(`([a-zA-Z0-9[\]()$@._=-]{3,})\((\d+)\):{[a-fA-F0-9]{6}}`),
			prefix:         regexp.MustCompile(` ({[a-fA-F0-9]{6}}.+){FFFFFF}`),
			message:        regexp.MustCompile(`\(\d+\):\{[a-fA-F0-9]{6}}(.+)$`),
			typeWarn:       regexp.MustCompile(`^(Модератор|Андроид) ([a-zA-Z0-9[\]()$@._=-]{3,}) ([а-я ]+) ([a-zA-Z0-9[\]()$@._=-]{3,}) ?н?а? ?(\d+ ч)?\..*$`),
			typeWarnReason: regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}.*\. Причина: (.+)$`),
			typeWarnTime:   regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,} на (\d+ .+)\..*$`),
			typeWarnUntil:  regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}. До (\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2})`),
			typeServer:     regexp.MustCompile(`^(\[SERVER]|SERVER): *(\{[a-fA-F0-9]{6}})?(.+)$`),
			typeSale:       regexp.MustCompile(`^(\[SALE]|SALE): *(\{[a-fA-F0-9]{6}})?(.+)$`),
			typePMForBot:   regexp.MustCompile(`^\(\( PM от ((([a-zA-Z0-9[\]()$@._=-]{3,}) \((\d+)\))|(Призрак)): (.+) \)\)$`),
			typePMFromBot:  regexp.MustCompile(`^\(\( PM к ([a-zA-Z0-9[\]()$@._ =-]{3,}) \((\d+)\): (.+) \)\)( \| AFK (\d+) сек.)?$`),
			typeSADS:       regexp.MustCompile(`^\[S-ADS #\d+] \{[a-fA-F0-9]{6}}([a-zA-Z0-9[\]()$@._=-]{3,})\((\d+)\):\{[a-fA-F0-9]{6}} *(.+)$`),
			typeUser:       regexp.MustCompile(`^([a-zA-Z0-9[\]()$@._=-]{3,}) \[(\d+)]$|^\*{2} ([a-zA-Z0-9[\]()$@._=-]{3,}) - UserID: (\d+) ?\|? ?A?F?K? ?(\d+)? ?с?е?к?\.?$`),

			ChatTypeUser:      regexp.MustCompile(`^([a-zA-Z0-9[\]()$@._=-]{3,}) \[(\d+)]$|^\*{2} ([a-zA-Z0-9[\]()$@._=-]{3,}) - UserID: (\d+) ?\|? ?A?F?K? ?(\d+)? ?с?е?к?\.?$`),
			chatTypeSADS:      regexp.MustCompile(`^\[S-ADS #\d+]`),
			chatTypeWorld:     regexp.MustCompile(`:\{91FF00} *(.+)$`),
			chatTypeGlobal:    regexp.MustCompile(`:\{(FFA500|F4831B)} *(.+)$`),
			chatTypeADS:       regexp.MustCompile(`^\[ADS]`),
			chatTypeASK:       regexp.MustCompile(`^\[ASK]`),
			chatTypeServer:    regexp.MustCompile(`^(\[SERVER]|SERVER):`),
			chatTypeSale:      regexp.MustCompile(`^(\[SALE]|SALE):`),
			chatTypeWarn:      regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}.*$`),
			chatTypePMForBot:  regexp.MustCompile(`^\(\( PM от `),
			chatTypePMFromBot: regexp.MustCompile(`^\(\( PM к `),
		},
	}
}

// Parse parses chat message and returns struct interface
func (p *Parser) Parse(text string) (interface{}, error) {
	if p.regexps.chatTypeADS.FindStringSubmatch(text) != nil {
		return p.ParseMessageADS(text)
	} else if p.regexps.chatTypeWorld.FindStringSubmatch(text) != nil {
		return p.ParseMessageWorld(text)
	} else if p.regexps.chatTypeGlobal.FindStringSubmatch(text) != nil {
		return p.ParseMessageGlobal(text)
	} else if p.regexps.chatTypeASK.FindStringSubmatch(text) != nil {
		return p.ParseMessageASK(text)
	} else if p.regexps.chatTypeWarn.FindStringSubmatch(text) != nil {
		return p.ParseMessageWarn(text)
	} else if p.regexps.chatTypeServer.FindStringSubmatch(text) != nil {
		return p.ParseMessageServer(text)
	} else if p.regexps.chatTypeSale.FindStringSubmatch(text) != nil {
		return p.ParseMessageSale(text)
	} else if p.regexps.chatTypePMForBot.FindStringSubmatch(text) != nil {
		return p.ParseMessagePMForBot(text)
	} else if p.regexps.chatTypePMFromBot.FindStringSubmatch(text) != nil {
		return p.ParseMessagePMFromBot(text)
	} else if p.regexps.chatTypeSADS.FindStringSubmatch(text) != nil {
		return p.ParseMessageSADS(text)
	} else if p.regexps.ChatTypeUser.FindStringSubmatch(text) != nil {
		return p.ParseMessageUser(text)
	} else {
		return p.ParseMessageAny(text)
	}
}

// ParseMessageUser parses user message and returns struct MessageUser
func (p *Parser) ParseMessageUser(text string) (*model.MessageUser, error) {
	var message = model.MessageUser{
		Type:      model.ChatMessageTypeUser,
		Timestamp: int(time.Now().UTC().Unix()),
	}

	matches := p.regexps.ChatTypeUser.FindStringSubmatch(text)
	if matches == nil || len(matches) < 4 || (matches[1] == "" && matches[3] == "") {
		return nil, fmt.Errorf("user not found")
	}

	if matches[1] != "" {
		message.Nick = matches[1]
		userID, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
		message.ID = userID
	} else {
		message.Nick = matches[3]
		userID, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}
		message.ID = userID
		if matches[5] != "" {
			afk, err := strconv.Atoi(matches[5])
			if err != nil {
				return nil, err
			}
			message.AFK = afk
		}
	}

	return &message, nil
}

// ParseMessageSADS parses SADS message and returns struct MessageSADS
func (p *Parser) ParseMessageSADS(text string) (*model.MessageSADS, error) {
	var message model.MessageSADS

	message.Type = model.ChatMessageTypeSADS
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeSADS.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
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

// ParseMessageAny parses any message and returns struct MessageAny
func (p *Parser) ParseMessageAny(text string) (*model.MessageAny, error) {
	var message model.MessageAny

	message.Type = model.ChatMessageTypeAny

	message.Message = text

	message.Timestamp = int(time.Now().UTC().Unix())

	return &message, nil
}

// ParseMessagePMFromBot parses PM from bot message and returns struct MessagePMFromBot
func (p *Parser) ParseMessagePMFromBot(text string) (*model.MessagePMFromBot, error) {
	var message model.MessagePMFromBot

	message.Type = model.ChatMessageTypePMFromBot
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typePMFromBot.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
	}

	message.Nick = strings.ReplaceAll(matches[1], " ", "_")

	userID, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	message.UserID = userID

	message.Message = strings.TrimSpace(matches[3])
	if strings.HasPrefix(message.Message, "code ") {
		return nil, nil
	}

	if len(matches) > 5 && matches[5] != "" {
		afk, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, err
		}

		message.AFK = afk
	}

	return &message, nil
}

// ParseMessagePMForBot parses PM for bot message and returns struct MessagePMForBot
func (p *Parser) ParseMessagePMForBot(text string) (*model.MessagePMForBot, error) {
	var message model.MessagePMForBot

	message.Type = model.ChatMessageTypePMForBot
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typePMForBot.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
	}

	if matches[3] != "" {
		message.Nick = matches[3]

		userID, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}

		message.UserID = userID
	} else if matches[5] != "" {
		message.UserID = -1
		message.Nick = matches[5]
	}

	message.Message = strings.TrimSpace(matches[6])

	return &message, nil
}

// ParseMessageSale parses sale message and returns struct MessageSale
func (p *Parser) ParseMessageSale(text string) (*model.MessageSale, error) {
	var message model.MessageSale

	message.Type = model.ChatMessageTypeSale
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeSale.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
	}

	if len(matches) < 3 {
		log.Error(matches)
		return nil, fmt.Errorf("message not found")
	}

	message.Message = strings.TrimSpace(matches[3])

	return &message, nil
}

// ParseMessageServer parses server message and returns struct MessageServer
func (p *Parser) ParseMessageServer(text string) (*model.MessageServer, error) {
	var message model.MessageServer

	message.Type = model.ChatMessageTypeServer
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeServer.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("message not found")
	}

	if len(matches) < 3 {
		log.Error(matches)
		return nil, fmt.Errorf("message not found")
	}

	message.Message = strings.TrimSpace(matches[3])

	return &message, nil
}

// ParseMessageWorld parses world message and returns struct MessageWorld
func (p *Parser) ParseMessageWorld(text string) (*model.MessageWorld, error) {
	var message model.MessageWorld

	message.Type = model.ChatMessageTypeWorld
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

// ParseMessageADS parses ADS message and returns struct MessageADS
func (p *Parser) ParseMessageADS(text string) (*model.MessageADS, error) {
	var message model.MessageADS

	message.Type = model.ChatMessageTypeAds
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

	return &message, nil
}

// ParseMessageASK parses ASK message and returns struct MessageASK
func (p *Parser) ParseMessageASK(text string) (*model.MessageASK, error) {
	var message model.MessageASK

	message.Type = model.ChatMessageTypeAsk
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

	return &message, nil
}

// ParseMessageWarn parses warn message and returns struct MessageWarn
func (p *Parser) ParseMessageWarn(text string) (*model.MessageWarn, error) {
	var message model.MessageWarn

	message.Type = model.ChatMessageTypeWarn
	message.Timestamp = int(time.Now().UTC().Unix())

	matches := p.regexps.typeWarn.FindStringSubmatch(text)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("user not found")
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

// ParseMessageGlobal parses global message and returns struct MessageGlobal
func (p *Parser) ParseMessageGlobal(text string) (*model.MessageGlobal, error) {
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
