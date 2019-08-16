package news

import (
	"fmt"
)

type Query struct {
	Q string
	QInTitle string
	Sources string
	Domains string
	ExcludeDomains string
	From string
	To string
	Language string
	SortBy string
	PageSize int
	Page int
}
func (q *Query) ToMap() map[string]string {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 5 {
		q.PageSize = 5
	}
	return map[string]string{
		"q": q.Q,
		"qInTitle": q.QInTitle,
		"sources": q.Sources,
		"domains": q.Domains,
		"excludeDomains": q.ExcludeDomains,
		"from": q.From,
		"to": q.To,
		"language": q.Language,
		"sortBy": q.SortBy,
		"pageSize": fmt.Sprintf("%d", q.PageSize),
		"page": fmt.Sprintf("%d", q.Page),
	}
}