package dataprovider

import (
	"fmt"

	"github.com/emersion/go-imap/client"
)

func ImapClientConnector(imapUrl string, port int, userName string, password string) (*client.Client, error) {
	imapBox, err := client.DialTLS(fmt.Sprintf("%s:%d", imapUrl, port), nil)
	if err != nil {
		return nil, err
	}

	if err := imapBox.Login(userName, password); err != nil {
		return nil, err
	}

	return imapBox, nil

}
