package model

// Type - тип обработанного сообщения
type Type string

const (
	Chat                     Type = "CHAT"                  // Chat - сообщение в чате
	Dialog                   Type = "DIALOG"                // Dialog - диалог
	ChatMessageTypeAds       Type = "CHAT_ADS"              // ChatMessageTypeAds - сообщение в рекламном чате
	ChatMessageTypeWorld     Type = "CHAT_WORLD"            // ChatMessageTypeWorld - сообщение в мире
	ChatMessageTypeGlobal    Type = "CHAT_GLOBAL"           // ChatMessageTypeGlobal - сообщение в глобальном чате
	ChatMessageTypeAsk       Type = "CHAT_ASK"              // ChatMessageTypeAsk - сообщение в чате вопросов
	ChatMessageTypeWarn      Type = "CHAT_WARN"             // ChatMessageTypeWarn - сообщение о выдаче наказания модератором
	ChatMessageTypeServer    Type = "CHAT_SERVER"           // ChatMessageTypeServer - сообщение от сервера
	ChatMessageTypeSale      Type = "CHAT_SALE"             // ChatMessageTypeSale - сообщение о скидках
	ChatMessageTypeAny       Type = "CHAT_ANY"              // ChatMessageTypeAny - сообщение в чате
	ChatMessageTypePMForBot  Type = "CHAT_PM_FOR_BOT"       // ChatMessageTypePMForBot - сообщение от игрока боту
	ChatMessageTypePMFromBot Type = "CHAT_PM_FROM_BOT"      // ChatMessageTypePMFromBot - сообщение от бота игроку
	ChatMessageTypeSADS      Type = "CHAT_SADS"             // ChatMessageTypeSADS - сообщение в SADS чате
	ChatMessageTypeUser      Type = "CHAT_USER"             // ChatMessageTypeUser - сообщение об игроке
	DialogAccountStatsType   Type = "DIALOG_ACCOUNT_STATS"  // DialogAccountStatsType - статистика аккаунта
	DialogWorldsListType     Type = "DIALOG_WORLDS_LIST"    // DialogWorldsListType - список миров
	DialogCopChaseListType   Type = "DIALOG_COP_CHASE_LIST" // DialogCopChaseListType - список лобби
	DialogRulesType          Type = "DIALOG_RULES"          // DialogRulesType - правила
)

// String - преобразование типа в строку
func (t Type) String() string {
	return string(t)
}

// Rules - структура правил
type Rules struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Rules     []Rule `json:"rules"`     // Правила
	Timestamp int64  `json:"timestamp"` // Время
}

// Rule - структура правила
type Rule struct {
	Point string `json:"point"` // Пункт
	Text  string `json:"text"`  // Текст
}

// WorldsList - структура списка миров
type WorldsList struct {
	Type      Type    `json:"type"`      // Тип сообщения
	Worlds    []World `json:"worlds"`    // Миры
	Timestamp int64   `json:"timestamp"` // Время
}

// World - структура мира
type World struct {
	Name    string `json:"name"`    // Название мира
	Players int    `json:"players"` // Игроков
	Static  bool   `json:"static"`  // Статичный
	SSMP    bool   `json:"ssmp"`    // S&SMP
}

// CopChaseList - структура списка лобби
type CopChaseList struct {
	Type      Type    `json:"type"`      // Тип сообщения
	Lobbies   []Lobby `json:"lobbies"`   // Лобби
	Timestamp int64   `json:"timestamp"` // Время
}

// Lobby - структура лобби
type Lobby struct {
	Number     int    `json:"number"`      // Номер лобби
	Status     string `json:"status"`      // Статус
	Time       string `json:"time"`        // Время
	Rating     string `json:"rating"`      // Рейтинг
	Players    int    `json:"players"`     // Игроки
	MaxPlayers int    `json:"max_players"` // Максимум игроков
}

// MessageUser - структура сообщения об игроке
type MessageUser struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	ID        int    `json:"id"`        // User
	AFK       int    `json:"afk"`       // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessagePMFromBot - структура сообщения от бота игроку
type MessagePMFromBot struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	AFK       int    `json:"afk"`       // AFK секунды
	Timestamp int64  `json:"timestamp"` // Время
}

// MessagePMForBot - структура сообщения от игрока боту
type MessagePMForBot struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageAny - структура сообщения в чате
type MessageAny struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageSale - структура сообщения о скидках
type MessageSale struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageServer - структура сообщения от сервера
type MessageServer struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageASK - структура сообщения в чате вопросов
type MessageASK struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageADS - структура сообщения в рекламном чате
type MessageADS struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageWorld - структура сообщения в мире
type MessageWorld struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	Star      string `json:"star"`      // Звезда
	Prefix    string `json:"prefix"`    // Префикс
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageGlobal - структура сообщения в глобальном чате
type MessageGlobal struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // UserID игрока
	Message   string `json:"message"`   // Текст сообщения
	Star      string `json:"star"`      // Звезда
	Prefix    string `json:"prefix"`    // Префикс
	Timestamp int64  `json:"timestamp"` // Время
}

// MessageWarn - структура сообщения о выдаче наказания модератором
type MessageWarn struct {
	Type       Type   `json:"type"`       // Тип сообщения
	Moderator  string `json:"moderator"`  // NickName модератора
	User       string `json:"user"`       // NickName игрока
	Reason     string `json:"reason"`     // Причина
	Punishment string `json:"punishment"` // Наказание (предупреждение, мут, бан, кик)
	Time       string `json:"time"`       // Минут
	Until      string `json:"until"`      // До
	Timestamp  int64  `json:"timestamp"`  // Время
}

// MessageSADS - структура сообщения в SADS чате
type MessageSADS struct {
	Type      Type   `json:"type"`      // Тип сообщения
	Nick      string `json:"nick"`      // NickName игрока
	UserID    int    `json:"id"`        // Id игрока
	Message   string `json:"message"`   // Текст сообщения
	Timestamp int64  `json:"timestamp"` // Время
}

// AccountStats - структура статистики аккаунта
type AccountStats struct {
	Type           Type          `json:"type"`             // Тип сообщения
	AccountID      int           `json:"account_id"`       // Id аккаунта
	AccountName    string        `json:"account_name"`     // Имя аккаунта
	Moderator      bool          `json:"moderator"`        // Статус модератора
	VIP            string        `json:"vip"`              // VIP статус
	Premium        int64         `json:"premium"`          // Premium статус
	SocialCredits  float64       `json:"social_credits"`   // Социальные кредиты
	Warns          int           `json:"warns"`            // Количество предупреждений
	BonusPoints    int           `json:"bonusPoints"`      // BonusPoints
	Kills          int           `json:"kills"`            // Количество убийств
	Deaths         int           `json:"deaths"`           // Количество смертей
	CopChaseRating int           `json:"cop_chase_rating"` // Рейтинг CopChase
	Punishments    []string      `json:"punishments"`      // Ограничения
	Verification   string        `json:"verification"`     // Подтверждение аккаунта
	Achievement    string        `json:"achievement"`      // Достижение
	Timestamp      int64         `json:"timestamp"`        // Время
	Descriptions   []Description `json:"descriptions"`     // Росписи модераторов
}

type Description struct {
	Date     string `json:"date"`     // Дата
	NickName string `json:"nickname"` // NickName игрока
	Text     string `json:"text"`     // Текст сообщения
}
