package dataprovider

import (
	"NoSpamGo/domain"
	"errors"
	"log"
	"os"
	"strconv"
)

func EnvironmentVariableGetter() (*domain.Argument, error) {
	username := os.Getenv("IMAP_USERNAME")

	password := os.Getenv("IMAP_PASSWORD")

	imapAddress := os.Getenv("IMAP_URL")

	imapPort := os.Getenv("IMAP_PORT")

	inputFiltersFile := os.Getenv("INPUT_FILE_FILTERS_LIST")

	if len(username) == 0 ||
		len(password) == 0 ||
		len(imapAddress) == 0 ||
		len(imapPort) == 0 ||
		len(inputFiltersFile) == 0 {
		log.Println("Mandatory env variables list :")
		log.Println("IMAP_USERNAME")
		log.Println("IMAP_PASSWORD")
		log.Println("IMAP_URL")
		log.Println("IMAP_PORT")
		log.Println("INPUT_FILE_FILTERS_LIST")
		return nil, errors.New("missing mandatory environment variables")
	}
	port, err := strconv.Atoi(imapPort)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	arguments := new(domain.Argument)
	arguments.UserName = username
	arguments.Password = password
	arguments.ImapUrl = imapAddress
	arguments.Port = port
	arguments.InputFilterFileName = inputFiltersFile

	return arguments, nil
}
