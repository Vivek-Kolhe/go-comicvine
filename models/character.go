package models

type CharacterBase struct {
	Aliases                 *string           `json:"aliases"`
	ApiDetailURL            *string           `json:"api_detail_url"`
	Birth                   *string           `json:"birth"`
	CountOfIssueAppearances *int              `json:"count_of_issue_appearances"`
	DateAdded               *string           `json:"date_added"`
	DateLastUpdated         *string           `json:"date_last_updated"`
	Deck                    *string           `json:"deck"`
	Description             *string           `json:"description"`
	FirstAppearedInIssue    *GenericIssue     `json:"first_appeared_in_issue"`
	Gender                  *int              `json:"gender"`
	ID                      *int              `json:"id"`
	Image                   *Image            `json:"image"`
	Name                    *string           `json:"name"`
	Origin                  *GenericOrigin    `json:"origin"`
	Publisher               *GenericPublisher `json:"publisher"`
	RealName                *string           `json:"real_name"`
	SiteDetailURL           *string           `json:"site_detail_url"`
}

type Character struct {
	CharacterBase
	CharacterFriends []*Friend          `json:"character_friends"`
	CharacterEnemies []*Enemy           `json:"character_enemies"`
	Creators         []*Creator         `json:"creators"`
	IssueCredits     []*GenericIssue    `json:"issue_credits"`
	IssuesDiedIn     []*GenericIssue    `json:"issues_died_in"`
	Movies           []*Movie           `json:"movies"`
	Powers           []*GenericPower    `json:"powers"`
	StoryArcCredits  []*GenericStoryArc `json:"story_arc_credits"`
	Teams            []*GenericTeam     `json:"teams"`
	TeamEnemies      []*Enemy           `json:"team_enemies"`
	TeamFriends      []*Friend          `json:"team_friends"`
	VolumeCredits    []*GenericVolume   `json:"volume_credits"`
}
