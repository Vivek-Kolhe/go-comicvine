package models

type Base struct {
	ApiDetailURL  *string `json:"api_detail_url"`
	ID            *int    `json:"id"`
	Name          *string `json:"name"`
	SiteDetailURL *string `json:"site_detail_url"`
}

type IssueBase struct {
	Base
	IssueNumber *string `json:"issue_number"`
}

type PublisherBase struct {
	Base
}
