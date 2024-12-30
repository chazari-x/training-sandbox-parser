package training_sandbox_parser

import (
	"regexp"

	"github.com/chazari-x/training-sandbox-parser/chat/ads"
	"github.com/chazari-x/training-sandbox-parser/chat/any"
	"github.com/chazari-x/training-sandbox-parser/chat/ask"
	"github.com/chazari-x/training-sandbox-parser/chat/global"
	"github.com/chazari-x/training-sandbox-parser/chat/pmForBot"
	"github.com/chazari-x/training-sandbox-parser/chat/pmFromBot"
	"github.com/chazari-x/training-sandbox-parser/chat/sads"
	"github.com/chazari-x/training-sandbox-parser/chat/sale"
	"github.com/chazari-x/training-sandbox-parser/chat/server"
	"github.com/chazari-x/training-sandbox-parser/chat/user"
	"github.com/chazari-x/training-sandbox-parser/chat/warn"
	"github.com/chazari-x/training-sandbox-parser/chat/world"
	"github.com/chazari-x/training-sandbox-parser/dialog/copchase"
	"github.com/chazari-x/training-sandbox-parser/dialog/list"
	"github.com/chazari-x/training-sandbox-parser/dialog/rules"
	"github.com/chazari-x/training-sandbox-parser/dialog/stats"
	"github.com/chazari-x/training-sandbox-parser/model"
)

type Parser struct {
	Chat    *chat    // Chat contains all chat message parsers for different types.
	Dialog  *dialog  // Dialog contains all dialog message parsers for different types.
	regexps *regexps // regexps contains all regular expressions used for parsing.
}

type chat struct {
	Ads       *ads.Parser       // Ads contains the parser for chat messages of type model.ChatMessageTypeAds.
	Any       *any.Parser       // Any contains the parser for chat messages of type model.ChatMessageTypeAny.
	Ask       *ask.Parser       // Ask contains the parser for chat messages of type model.ChatMessageTypeAsk.
	Global    *global.Parser    // Global contains the parser for chat messages of type model.ChatMessageTypeGlobal.
	PmForBot  *pmForBot.Parser  // PmForBot contains the parser for chat messages of type model.ChatMessageTypePMForBot.
	PmFromBot *pmFromBot.Parser // PmFromBot contains the parser for chat messages of type model.ChatMessageTypePMFromBot.
	SAds      *sads.Parser      // SAds contains the parser for chat messages of type model.ChatMessageTypeSADS.
	Sale      *sale.Parser      // Sale contains the parser for chat messages of type model.ChatMessageTypeSale.
	Server    *server.Parser    // Server contains the parser for chat messages of type model.ChatMessageTypeServer.
	User      *user.Parser      // User contains the parser for chat messages of type model.ChatMessageTypeUser.
	Warn      *warn.Parser      // Warn contains the parser for chat messages of type model.ChatMessageTypeWarn.
	World     *world.Parser     // World contains the parser for chat messages of type model.ChatMessageTypeWorld.
}

type dialog struct {
	Stats    *stats.Parser    // Stats contains the parser for dialog messages of type model.DialogAccountStatsType.
	List     *list.Parser     // List contains the parser for dialog messages of type model.DialogWorldsListType.
	CopChase *copchase.Parser // CopChase contains the parser for dialog messages of type model.DialogCopChaseListType.
	Rules    *rules.Parser    // Rules contains the parser for dialog messages of type model.DialogRulesType.
}

type regexps struct {
	dialog             *regexp.Regexp
	chat               *regexp.Regexp
	dialogTypeStats    *regexp.Regexp
	dialogTypeList     *regexp.Regexp
	dialogTypeCopChase *regexp.Regexp
	dialogTypeRules    *regexp.Regexp
	chatTypeUser       *regexp.Regexp
	chatTypeSADS       *regexp.Regexp
	chatTypeWorld      *regexp.Regexp
	chatTypeGlobal     *regexp.Regexp
	chatTypeADS        *regexp.Regexp
	chatTypeASK        *regexp.Regexp
	chatTypeWarn       *regexp.Regexp
	chatTypeServer     *regexp.Regexp
	chatTypeSale       *regexp.Regexp
	chatTypePMForBot   *regexp.Regexp
	chatTypePMFromBot  *regexp.Regexp
}

// New creates a new instance of the Parser struct with initialized sub-parsers and regular expressions.
//
// Returns:
// - A pointer to the newly created Parser instance.
func New() *Parser {
	return &Parser{
		Chat: &chat{
			Ads:       ads.New(),
			Any:       any.New(),
			Ask:       ask.New(),
			Global:    global.New(),
			PmForBot:  pmForBot.New(),
			PmFromBot: pmFromBot.New(),
			SAds:      sads.New(),
			Sale:      sale.New(),
			Server:    server.New(),
			User:      user.New(),
			Warn:      warn.New(),
			World:     world.New(),
		},

		Dialog: &dialog{
			Stats:    stats.New(),
			List:     list.New(),
			CopChase: copchase.New(),
			Rules:    rules.New(),
		},

		regexps: &regexps{
			dialog:             regexp.MustCompile(`^\[DIALOG]: ID: (\d+) \| STYLE: (\d+) \| TITLE: (.*) ?\| BTN1: (.*) ?\| BTN2: (.*) ?\| TEXT:`),
			chat:               regexp.MustCompile(`^\[CHAT]: (.*)`),
			dialogTypeList:     regexp.MustCompile(`\{FFFFFF}Название мира.+\{FFFFFF}Онлайн.*`),
			dialogTypeStats:    regexp.MustCompile(`\{[a-zA-Z0-9]{6}}Статистика аккаунта: +\{FFFFFF}.+ +#\d+`),
			dialogTypeCopChase: regexp.MustCompile(`\{FFFFFF}#Лобби.+\{FFFFFF}Статус.+\{FFFFFF}Рейтинг:.+\d.+\{FFFFFF}Онлайн.*\n`),
			dialogTypeRules:    regexp.MustCompile(`^(\{[a-zA-Z0-9]{6}})? *(\d.\d*) *(\{[a-zA-Z0-9]{6}})? *(.+)\n`),
			chatTypeUser:       regexp.MustCompile(`^([a-zA-Z0-9[\]()$@._=-]{3,}) \[(\d+)]$|^\*{2} ([a-zA-Z0-9[\]()$@._=-]{3,}) - UserID: (\d+) ?\|? ?A?F?K? ?(\d+)? ?с?е?к?\.?$`),
			chatTypeSADS:       regexp.MustCompile(`^\[S-ADS #\d+]`),
			chatTypeWorld:      regexp.MustCompile(`:\{91FF00} *(.+)$`),
			chatTypeGlobal:     regexp.MustCompile(`\(\d+\):\{[a-zA-Z0-9]{6}} *(.+)$`),
			chatTypeADS:        regexp.MustCompile(`^\[ADS]`),
			chatTypeASK:        regexp.MustCompile(`^\[ASK]`),
			chatTypeServer:     regexp.MustCompile(`^(\[SERVER]|SERVER):`),
			chatTypeSale:       regexp.MustCompile(`^(\[SALE]|SALE):`),
			chatTypeWarn:       regexp.MustCompile(`^(Модератор|Андроид) [a-zA-Z0-9[\]()$@._=-]{3,} [а-я ]+ [a-zA-Z0-9[\]()$@._=-]{3,}.*$`),
			chatTypePMForBot:   regexp.MustCompile(`^\(\( PM от `),
			chatTypePMFromBot:  regexp.MustCompile(`^\(\( PM к `),
		},
	}
}

// Type determines the type of the given text by matching it against predefined regular expressions.
// It returns a slice of strings containing the matched groups and the type of the text.
//
// Parameters:
// - text: The input string to be parsed.
//
// Returns:
// - A slice of strings containing the matched groups if a match is found, otherwise nil.
// - A model.Type indicating the type of the text if a match is found, otherwise an empty string.
func (p *Parser) Type(text string) ([]string, model.Type) {
	if matches := p.regexps.chat.FindStringSubmatch(text); matches != nil {
		return matches, model.Chat
	} else if matches = p.regexps.dialog.FindStringSubmatch(text); matches != nil {
		return matches, model.Dialog
	}

	return nil, ""
}

// DialogType determines the specific type of dialog by matching the given text against predefined regular expressions.
// It returns a model.Type indicating the type of the dialog.
//
// Parameters:
// - text: The input string to be parsed.
//
// Returns:
// - A model.Type indicating the type of the dialog if a match is found, otherwise an empty string.
func (p *Parser) DialogType(text string) model.Type {
	if matches := p.regexps.dialogTypeStats.FindStringSubmatch(text); matches != nil {
		return model.DialogAccountStatsType
	} else if matches = p.regexps.dialogTypeList.FindStringSubmatch(text); matches != nil {
		return model.DialogWorldsListType
	} else if matches = p.regexps.dialogTypeCopChase.FindStringSubmatch(text); matches != nil {
		return model.DialogCopChaseListType
	} else if matches = p.regexps.dialogTypeRules.FindStringSubmatch(text); matches != nil {
		return model.DialogRulesType
	}

	return ""
}

// ChatType determines the type of chat message by matching the given text against predefined regular expressions.
// It returns a model.Type indicating the type of the chat message.
//
// Parameters:
// - text: The input string to be parsed.
//
// Returns:
// - A model.Type indicating the type of the chat message if a match is found, otherwise model.ChatMessageTypeAny.
func (p *Parser) ChatType(text string) model.Type {
	switch {
	case p.regexps.chatTypeADS.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeAds
	case p.regexps.chatTypeWorld.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeWorld
	case p.regexps.chatTypeASK.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeAsk
	case p.regexps.chatTypeWarn.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeWarn
	case p.regexps.chatTypeServer.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeServer
	case p.regexps.chatTypeSale.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeSale
	case p.regexps.chatTypePMForBot.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypePMForBot
	case p.regexps.chatTypePMFromBot.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypePMFromBot
	case p.regexps.chatTypeSADS.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeSADS
	case p.regexps.chatTypeUser.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeUser
	case p.regexps.chatTypeGlobal.FindStringSubmatch(text) != nil:
		return model.ChatMessageTypeGlobal
	default:
		return model.ChatMessageTypeAny
	}
}
