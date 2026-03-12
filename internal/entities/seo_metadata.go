package entities

type SeoMetadata struct {
	ID            string `json:"id"`
	ReferenceID   string `json:"reference_id"`   // ID dari Project atau Page
	ReferenceType string `json:"reference_type"` // 'Project' atau 'Page'
	MetaTitle     string `json:"meta_title"`
	MetaDesc      string `json:"meta_desc"`
	OgImageURL    string `json:"og_image_url"`
	AltText       string `json:"alt_text"`
}
