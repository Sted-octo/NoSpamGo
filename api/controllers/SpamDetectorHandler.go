package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/presenter"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/emersion/go-imap/client"
	"github.com/julienschmidt/httprouter"
)

func SpamDetectorHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var emails dataprovider.Emails
	if err := json.NewDecoder(r.Body).Decode(&emails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxWorkers := 5
	jobs := make(chan string, len(emails.Mails))
	results := make(chan presenter.EmailResult, len(emails.Mails))

	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	for _, email := range emails.Mails {
		jobs <- email
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var allResults []presenter.EmailResult
	for result := range results {
		allResults = append(allResults, result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}

func worker(wg *sync.WaitGroup, jobs <-chan string, results chan<- presenter.EmailResult) {
	defer wg.Done()

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer dbConnector.Close()

	var userByMailLoader usecases.IUserByMailLoader[*sql.DB] = new(dataprovider.UserByMailLoader)
	var unseenMessagesGetter usecases.IUnseenMessagesGetter[*client.Client] = new(dataprovider.ImapClientUnseenMessagesGetter)
	var spamMover usecases.ISpamMover[*client.Client] = new(dataprovider.ImapClientSpamMover)
	var clientConnector usecases.IClientConnector[*client.Client] = new(dataprovider.ImapClientConnector)
	var filtersGetter usecases.IFiltersGetter[*sql.DB] = new(dataprovider.FiltersGetter)
	var filterByNameForUserMailLoader usecases.IFilterByNameForUserMailLoader[*sql.DB] = new(dataprovider.FilterByNameForUserMailLoader)
	var filterSaver usecases.IFilterSaver[*sql.DB] = new(dataprovider.FilterSaver)

	for email := range jobs {
		result := presenter.EmailResult{
			Mail:              email,
			CountSpamDetected: 0,
		}

		user := userByMailLoader.Load(email, dbConnector)
		if user != nil {

			err := clientConnector.Connect(user.ImapServerUrl, user.ImapServerPort, user.ImapUsername, user.ImapPassword)
			if err != nil {
				log.Println(err)
				break
			}
			defer clientConnector.Close()

			result.CountSpamDetected = usecases.SpamDetector[*client.Client, *sql.DB](email, clientConnector, unseenMessagesGetter, spamMover, filtersGetter, dbConnector, filterSaver, filterByNameForUserMailLoader)
		}

		results <- result
	}
}
