package main

import (
	dataprovider "NoSpamGo/dataProvider"
	"fmt"
	"log"
	"strings"

	"github.com/emersion/go-imap"
)

func main() {

	arguments, err := dataprovider.EnvironmentVariableGetter()
	if err != nil {
		log.Fatal(err)
	}
	imapBox, err := dataprovider.ImapClientConnector(arguments.ImapUrl, arguments.Port, arguments.UserName, arguments.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer imapBox.Logout()

	messages := dataprovider.ImapClientUnseenMessagesGetter(imapBox)

	if messages == nil {
		return
	}

	spamUIDs := new(imap.SeqSet)

	for msg := range messages {
		enveloppe := msg.Envelope

		for _, addressOrigin := range enveloppe.From {
			fmt.Printf("SeqNum: %d\n", msg.SeqNum)
			fmt.Printf("PersonalName: %s\n", addressOrigin.PersonalName)
			fmt.Printf("MailboxName: %s\n", addressOrigin.MailboxName)
			fmt.Printf("HostName: %s\n", addressOrigin.HostName)
			fmt.Printf("Email complet: %s@%s\n", addressOrigin.MailboxName, addressOrigin.HostName)
			if strings.Contains(strings.ToLower(addressOrigin.PersonalName), "happy promos") {
				spamUIDs.AddNum(msg.SeqNum)
				fmt.Println("!! SPAM !! ")
			}
		}
		fmt.Printf("Sujet: %s\n", enveloppe.Subject)

		fmt.Println("------------------------")
	}

	if !spamUIDs.Empty() {
		if err := imapBox.Move(spamUIDs, "Junk"); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done!")
}
