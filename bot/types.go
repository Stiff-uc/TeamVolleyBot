package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Store is an interface for the persistent storage
// should allow easier swapping of databases
type Store interface {
	Close()
	AddMsgToPoll(pollid int, messageid int, chatid int64) error
	AddInlineMsgToPoll(pollid int, inlinemessageid string) error
	RemoveInlineMsg(inlinemessageid string) error
	GetPoll(pollid int) (*poll, error)
	GetUser(userid int, chatID int64) (*tgbotapi.User, error)
	GetPollsByUser(userid int) ([]*poll, error)
	GetPollID(messageid int) (int, error)
	GetPollNewer(pollid int, userid int) (*poll, error)
	GetPollOlder(pollid int, userid int) (*poll, error)
	GetAllPollMsg(pollid int) ([]pollident, error)
	GetAllPollInlineMsg(pollid int) ([]pollident, error)
	GetState(userid int) (state int, pollid int, chatID int64, err error)
	SaveState(userid int, pollid int, state int, chatID int64) error
	SaveUser(*tgbotapi.User, int64) error
	SavePoll(*poll) (int, error)
	SaveOptions([]option) error
	SaveAnswer(*poll, answer) (unvoted bool, err error)
	SaveChat(*tgbotapi.Chat) error
	EnterChat(*tgbotapi.Chat, int) error
	LeaveChat(*tgbotapi.Chat) error
	GetUserChatIds(int) ([]chat, error)
	GetChatUsers(int64) ([]user, error)
	GetPlayer(int, int64) (user, error)
	SavePlayer(user) error
}

type answer struct {
	ID       int
	PollID   int
	UserID   int
	OptionID int
}

type option struct {
	ID     int
	PollID int
	Text   string
	Ctr    int
}

type chat struct {
	ID          int64
	Title       string
	Status      string
	AdminUserID int
}

type poll struct {
	ID        int
	MessageID int
	UserID    int
	ChatID    int64
	Question  string
	Inactive  int
	Private   int
	Type      int
	Options   []option
	Answers   []answer
}
type user struct {
	ID           int
	ChatID       int64
	FirstName    string
	LastName     string
	IsBot        string
	UserName     string
	Priority     int
	Tag          string
	NameOverride string
}

func isInactive(poll *poll) bool {
	return poll.Inactive == inactive
}

func isMultipleChoice(poll *poll) bool {
	return poll.Type == multipleChoice
}
