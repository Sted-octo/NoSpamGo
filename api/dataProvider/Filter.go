package dataprovider

import (
	"NoSpamGo/domain"
	"database/sql"
)

type Filter struct {
	Mail                       string
	FilterName                 sql.NullString
	FilterNumberOfSpamDetected sql.NullInt64
}

func (f *Filter) ToDomain() *domain.Filter {

	return &domain.Filter{
		Name:                 valueOrEmpty(f.FilterName),
		NumberOfSpamDetected: int(valueOrZero(f.FilterNumberOfSpamDetected)),
	}
}
