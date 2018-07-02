package main

import "fmt"

var (
	locGotQuestion          = "OK now that we got a question, please send answer options to your poll."
	locStartCommand         = "/start"
	locEditCommand          = "/edit"
	locCreateNewPoll        = "create new poll"
	locInlineInsertPoll     = "insert poll into chat"
	locSharePoll            = "share poll"
	locNewQuestion          = "Great! Send a question for the new poll, please."
	locEditQuestion         = "Okay, just send the new question, please."
	locEditQuestionButton   = "change question"
	locAddOption            = "Alright, send options that you want to add to the poll, please."
	locAddOptionButton      = "add options"
	locGotEditQuestion      = "Thanks, the question was changed to \"%s\"."
	locNoMessageToEdit      = "Sorry, I could not find a poll to edit."
	locFinishedCreatingPoll = "Finished creating a new poll\n\nPreview:\n"
	locMainMenu             = "I can help you create, send and manage polls.\n\nWhat do you want to do?"
	locPollDoneButton       = "done"
	locToggleInactive       = "open poll"
	locToggleOpen           = "close poll"
	locToggleSingleChoice   = "set multiple choice"
	locToggleMultipleChoice = "set single choice"
	locAddedOption          = fmt.Sprintf(
		"You can add more options by sending messages each containing one option. If you are done, please push the %s button.\n\nPreview:\n",
		locPollDoneButton)
)
