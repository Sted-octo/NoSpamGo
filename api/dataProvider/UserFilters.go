package dataprovider

import "NoSpamGo/domain"

type UserFilters struct {
	Mail    string
	Filters []domain.Filter
}
