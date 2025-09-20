package models

type GenericEntry struct {
	ApiDetailURL  *string `json:"api_detail_url"`
	ID            *int    `json:"id"`
	Name          *string `json:"name"`
	SiteDetailURL *string `json:"site_detail_url"`
}

type GenericIssue struct {
	GenericEntry
	IssueNumber *string `json:"issue_number"`
}

type Friend GenericEntry
type Enemy GenericEntry
type Creator GenericEntry
type Movie GenericEntry
type GenericOrigin GenericEntry
type GenericPublisher GenericEntry
type GenericPower GenericEntry
type GenericTeam GenericEntry
type GenericVolume GenericEntry
