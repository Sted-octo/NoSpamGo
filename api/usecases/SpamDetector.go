package usecases

import (
	"NoSpamGo/domain"
	"NoSpamGo/tools"
	"strings"
)

func SpamDetector[T any, S any](email string,
	clientConnector IClientConnector[T],
	unseenMessagesGetter IUnseenMessagesGetter[T],
	spamMover ISpamMover[T],
	filtersGetter IFiltersGetter[S],
	dbConnector IDatabaseConnector[S],
	filterSaver IFilterSaver[S],
	filterByNameForUserMailLoader IFilterByNameForUserMailLoader[S]) {

	var ids []uint32
	var filtersUsed []string
	messages := unseenMessagesGetter.Get(clientConnector)

	if messages == nil {
		return
	}
	filters := filtersGetter.Get(email, dbConnector)
	filterMap := make(map[string]*domain.Filter)
	for _, filter := range filters {
		filterMap[filter.Name] = &filter
	}

	for _, msg := range messages {
		subjectLower := strings.ToLower(msg.Subject)
		personnalNameLower := ""
		hostNameLower := ""
		mailboxName := ""
		for _, addressOrigin := range msg.Mails {
			if addressOrigin.PersonalName != "" {
				personnalNameLower = tools.Concat(personnalNameLower, strings.ToLower(addressOrigin.PersonalName))
			}
			if addressOrigin.HostName != "" {
				hostNameLower = tools.Concat(personnalNameLower, strings.ToLower(addressOrigin.HostName))
			}
			if addressOrigin.MailboxName != "" {
				mailboxName = tools.Concat(personnalNameLower, strings.ToLower(addressOrigin.MailboxName))
			}
		}

		for _, filter := range filters {
			if (subjectLower != "" && strings.Contains(subjectLower, filter.Name)) ||
				(personnalNameLower != "" && strings.Contains(personnalNameLower, filter.Name)) ||
				(hostNameLower != "" && strings.Contains(hostNameLower, filter.Name)) ||
				(mailboxName != "" && strings.Contains(mailboxName, filter.Name)) {
				ids = append(ids, msg.Id)
				filterMap[filter.Name].NumberOfSpamDetected++
				filtersUsed = append(filtersUsed, filter.Name)
				break
			}
		}
	}

	for _, filterName := range filtersUsed {
		filterSaver.Save(email, *filterMap[filterName], dbConnector, filterByNameForUserMailLoader)
	}

	if len(ids) > 0 {
		spamMover.Move(clientConnector, ids)
	}
}
