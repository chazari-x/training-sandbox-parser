package training_sandbox_parser

import (
	"github.com/chazari-x/training-sandbox-parser/chat"
	"github.com/chazari-x/training-sandbox-parser/copchase"
	"github.com/chazari-x/training-sandbox-parser/list"
	"github.com/chazari-x/training-sandbox-parser/rules"
	"github.com/chazari-x/training-sandbox-parser/stats"
)

type Parser struct {
	Chat     *chat.Parser
	Stats    *stats.Parser
	List     *list.Parser
	CopChase *copchase.Parser
	Rules    *rules.Parser
}

func NewParser() *Parser {
	return &Parser{
		Chat:     chat.NewParser(),
		Stats:    stats.NewParser(),
		List:     list.NewParser(),
		CopChase: copchase.NewParser(),
		Rules:    rules.NewParser(),
	}
}
