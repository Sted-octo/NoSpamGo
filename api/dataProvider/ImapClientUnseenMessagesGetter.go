package dataprovider

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func ImapClientUnseenMessagesGetter(imapBox *client.Client) chan *imap.Message {

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}

	_, err := imapBox.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	uids, err := imapBox.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}

	if len(uids) == 0 {
		return nil
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(uids...)

	items := []imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags, imap.FetchBody}
	messages := make(chan *imap.Message)
	done := make(chan error, 1)

	go func() {
		done <- imapBox.Fetch(seqSet, items, messages)
	}()

	return messages
}
