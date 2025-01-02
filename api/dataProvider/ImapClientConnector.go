package dataprovider

import (
	"fmt"

	"github.com/emersion/go-imap/client"
)

type ImapClientConnector struct {
	ImapBox *client.Client
}

func (box *ImapClientConnector) Connect(imapUrl string, port int, userName string, password string) error {
	var err error
	box.ImapBox, err = client.DialTLS(fmt.Sprintf("%s:%d", imapUrl, port), nil)
	if err != nil {
		return err
	}

	if err := box.ImapBox.Login(userName, password); err != nil {
		return err
	}

	return nil
}

func (box *ImapClientConnector) Close() {
	if box.ImapBox != nil {
		box.ImapBox.Close()
	}
}

func (box *ImapClientConnector) Get() *client.Client {
	return box.ImapBox
}
