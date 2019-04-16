package main

const createNewPollQuery = "createNewPoll"
const createPollQuery = "createpoll"
const pollDoneQuery = "polldone"
const selectChat = "chat"
const selectPlayer = "player"
const updateTag = "playertag"
const updatePriority = "playerpriority"
const updatePlayerName = "playername"
const (
	ohHi = iota
	waitingForQuestion
	waitingForOption
	pollDone
	editPoll
	editQuestion
	addOption
	listChats
	listPlayers
	waitingForPlayerSettingSelect
	waitingForTag
	waitingForName
	waitingForPriority
)

const (
	open = iota
	inactive
)

const (
	standard = iota
	multipleChoice
)

const (
	typeSkillsVote = iota
	typeGame
)

var maxNumberOfUsersListed = 100
var maxPollsInlineQuery = 5
var maxPlayersInTeams = 18
