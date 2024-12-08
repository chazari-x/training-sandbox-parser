package stats

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
	account        *regexp.Regexp
	vip            *regexp.Regexp
	socialCredits  *regexp.Regexp
	warns          *regexp.Regexp
	killsDeaths    *regexp.Regexp
	copChaseRating *regexp.Regexp
	punishment     *regexp.Regexp
	verification   *regexp.Regexp
	moderator      *regexp.Regexp
	achievement    *regexp.Regexp
}

func New() *Parser {
	return &Parser{
		regexps: &regexps{
			account:        regexp.MustCompile(`\{[a-zA-Z0-9]{6}}Статистика аккаунта:.*\{FFFFFF}(.+) #(\d+).*\n`),
			vip:            regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}\[VIP] (.+) *\n`),
			socialCredits:  regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Рейтинг Social Credits.*\{FFFFFF}(-?\d+\.\d+) *\n`),
			warns:          regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Предупреждения.*\{FFFFFF}(\d+) *\n`),
			killsDeaths:    regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Убийств/Смертей.*\{FFFFFF}(\d+)/(\d+) *\n`),
			copChaseRating: regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Рейтинг CopChase.*\{FFFFFF}(-?\d+) *\n`),
			punishment:     regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}(.*\{FFFFFF}.*-?.*\d{2,}:\d{2}:\d{2}\.?)\n`),
			verification:   regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Подтвержденный аккаунт:.*\{FFFFFF}(.+) *\n`),
			moderator:      regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}(Модератор.+) *\n`),
			achievement:    regexp.MustCompile(`\n\{[a-zA-Z0-9]{6}}Достижение:.*\{FFFFFF}(.+) *\n`),
		},
	}
}

// Parse parses the text and returns the account stats.
func (p *Parser) Parse(text string) (*model.AccountStats, error) {
	var accountStats model.AccountStats

	accountStats.Type = model.DialogAccountStatsType
	accountStats.Timestamp = time.Now().UTC().Unix()

	matches := p.regexps.account.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("account not found")
	}

	accountStats.AccountName = matches[1]

	accountID, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	accountStats.AccountID = accountID

	matches = p.regexps.vip.FindStringSubmatch(text)
	if matches != nil {
		accountStats.VIP = matches[1]
	}

	matches = p.regexps.socialCredits.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("social credits not found")
	}

	socialCredits, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return nil, err
	}

	accountStats.SocialCredits = socialCredits

	matches = p.regexps.warns.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("warns not found")
	}

	warns, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	accountStats.Warns = warns

	matches = p.regexps.killsDeaths.FindStringSubmatch(text)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("kills deaths not found")
	}

	deaths, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	kills, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	accountStats.Deaths = deaths

	accountStats.Kills = kills

	matches = p.regexps.copChaseRating.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("cop chase rating not found")
	}

	copChaseRating, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	accountStats.CopChaseRating = copChaseRating

	matches = p.regexps.punishment.FindStringSubmatch(text)
	if matches != nil {
		accountStats.Punishments = append(accountStats.Punishments, matches[1:]...)
	}

	matches = p.regexps.verification.FindStringSubmatch(text)
	if matches != nil {
		accountStats.Verification = matches[1]
	}

	matches = p.regexps.moderator.FindStringSubmatch(text)
	if matches != nil {
		accountStats.Moderator = true
	}

	matches = p.regexps.achievement.FindStringSubmatch(text)
	if matches != nil {
		accountStats.Achievement = matches[1]
	}

	return &accountStats, nil
}
