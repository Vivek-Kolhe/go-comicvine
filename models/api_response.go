package models

type ApiResponseBase struct {
	Error                string `json:"error"`
	Limit                int    `json:"limit"`
	Offset               int    `json:"offset"`
	NumberOfPageResults  int    `json:"number_of_page_results"`
	NumberOfTotalResults int    `json:"number_of_total_results"`
	StatusCode           int    `json:"status_code"`
	Version              string `json:"version"`
}

type ApiResponseSingleResult[T any] struct {
	ApiResponseBase
	Results *T `json:"results"`
}

type ApiResponseMultipleResult[T any] struct {
	ApiResponseBase
	Results []*T `json:"results"`
}
