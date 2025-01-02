package usecases

import (
	"fmt"
	"strings"
)

func SpamDetector[T any](clientConnector IClientConnector[T], unseenMessagesGetter IUnseenMessagesGetter[T], spamMover ISpamMover[T]) {

	var ids []uint32
	messages := unseenMessagesGetter.Get(clientConnector)

	if messages == nil {
		return
	}

	fmt.Printf("Messages non lus : %d\n", len(messages))

	for _, msg := range messages {
		for _, addressOrigin := range msg.Mails {
			if strings.Contains(strings.ToLower(addressOrigin.PersonalName), "happy promos") {
				ids = append(ids, msg.Id)
			}
		}
		fmt.Printf("Sujet: %s\n", msg.Subject)
	}

	fmt.Printf("Spams détecté : %d\n", len(ids))

	spamMover.Move(clientConnector, ids)
}
