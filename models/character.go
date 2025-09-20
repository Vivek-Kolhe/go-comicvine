package models

type CharacterBase struct {
	Aliases                 *string        `json:"aliases"`
	ApiDetailURL            *string        `json:"api_detail_url"`
	Birth                   *string        `json:"birth"`
	CountOfIssueAppearances *int           `json:"count_of_issue_appearances"`
	DateAdded               *string        `json:"date_added"`
	DateLastUpdated         *string        `json:"date_last_updated"`
	Deck                    *string        `json:"deck"`
	Description             *string        `json:"description"`
	FirstAppearedInIssue    *IssueBase     `json:"first_appeared_in_issue"`
	Gender                  *int           `json:"gender"`
	ID                      *int           `json:"id"`
	Image                   *Image         `json:"image"`
	Name                    *string        `json:"name"`
	Origin                  *Origin        `json:"origin"`
	Publisher               *PublisherBase `json:"publisher"`
	RealName                *string        `json:"real_name"`
	SiteDetailURL           *string        `json:"site_detail_url"`
}
