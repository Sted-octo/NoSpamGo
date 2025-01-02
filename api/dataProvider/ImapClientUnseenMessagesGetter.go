package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapClientUnseenMessagesGetter struct{}

func (o *ImapClientUnseenMessagesGetter) Get(clientConnector usecases.IClientConnector[*client.Client]) []domain.Message {

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}

	_, err := clientConnector.Get().Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	uids, err := clientConnector.Get().Search(criteria)
	if err != nil {
		log.Fatal(err)
	}

	if len(uids) == 0 {
		return nil
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(uids...)

	items := []imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags, imap.FetchBody}
	mails := make(chan *imap.Message)
	done := make(chan error, 1)

	var messages []domain.Message

	go func() {
		done <- clientConnector.Get().Fetch(seqSet, items, mails)
	}()

	for msg := range mails {
		enveloppe := msg.Envelope
		message := new(domain.Message)
		message.Subject = enveloppe.Subject
		message.Id = msg.SeqNum

		for _, addressOrigin := range enveloppe.From {
			address := new(domain.MailAddress)
			address.PersonalName = addressOrigin.PersonalName
			address.MailboxName = addressOrigin.MailboxName
			address.HostName = addressOrigin.HostName

			message.Mails = append(message.Mails, *address)
		}
		messages = append(messages, *message)
	}

	return messages
}
