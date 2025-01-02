package dataprovider

import (
	"NoSpamGo/usecases"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapClientSpamMover struct{}

func (o *ImapClientSpamMover) Move(clientConnector usecases.IClientConnector[*client.Client], ids []uint32) {

	spamUIDs := new(imap.SeqSet)

	for _, id := range ids {
		spamUIDs.AddNum(id)
	}

	if !spamUIDs.Empty() {
		if err := clientConnector.Get().Move(spamUIDs, "Junk"); err != nil {
			log.Fatal(err)
		}
	}

}
