package main

const createNewPollQuery = "createNewPoll"
const createPollQuery = "createpoll"
const pollDoneQuery = "polldone"
const selectChat = "chat"
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
	addPlayerHint
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
