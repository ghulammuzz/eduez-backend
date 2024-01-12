// helper/pagination.go

package helper

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
}

func GetPaginationLinks(baseURL string, page, pageSize, totalRows int) (nextLink, prevLink string) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	if offset+pageSize < totalRows {
		nextPage := page + 1
		nextLink = constructPaginationLink(baseURL, nextPage, pageSize)
	}

	if offset > 0 {
		prevPage := page - 1
		prevLink = constructPaginationLink(baseURL, prevPage, pageSize)
	}

	return nextLink, prevLink
}

func constructPaginationLink(baseURL string, page, pageSize int) string {
	u, _ := url.Parse(baseURL)
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("page_size", strconv.Itoa(pageSize))
	u.RawQuery = q.Encode()
	return u.String()
}

func CalculateTotalPages(totalRows, pageSize int) int {
	if pageSize <= 0 {
		pageSize = 10
	}
	return (totalRows + pageSize - 1) / pageSize
}
