package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/usecases"
	"log"
	"net/http"

	"github.com/emersion/go-imap/client"
	"github.com/julienschmidt/httprouter"
)

func SpamDetectorHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	arguments, err := dataprovider.EnvironmentVariableGetter()
	if err != nil {
		log.Fatal(err)
	}

	var unseenMessagesGetter usecases.IUnseenMessagesGetter[*client.Client] = new(dataprovider.ImapClientUnseenMessagesGetter)
	var spamMover usecases.ISpamMover[*client.Client] = new(dataprovider.ImapClientSpamMover)
	var clientConnector usecases.IClientConnector[*client.Client] = new(dataprovider.ImapClientConnector)

	err = clientConnector.Connect(arguments.ImapUrl, arguments.Port, arguments.UserName, arguments.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnector.Close()

	usecases.SpamDetector(clientConnector, unseenMessagesGetter, spamMover)
}
