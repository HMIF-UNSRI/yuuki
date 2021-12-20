package domain

const DefaultPaginationLimit = 3

type Pagination struct {
	Count    uint32 `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type PaginationParam struct {
	Limit    uint32
	CursorID uint32
}
