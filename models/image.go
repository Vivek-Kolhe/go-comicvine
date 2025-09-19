package models

type Image struct {
	IconURL        *string `json:"icon_url"`
	MediumURL      *string `json:"medium_url"`
	ScreenURL      *string `json:"screen_url"`
	ScreenLargeURL *string `json:"screen_large_url"`
	SmallURL       *string `json:"small_url"`
	SuperURL       *string `json:"super_url"`
	ThumbURL       *string `json:"thumb_url"`
	TinyURL        *string `json:"tiny_url"`
	OriginalURL    *string `json:"original_url"`
	ImageTags      *string `json:"image_tags"`
}
