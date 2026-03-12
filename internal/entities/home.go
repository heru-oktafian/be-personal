package entities

// Struktur untuk menangkap setting dalam format JSON
type SiteSetting struct {
	Key       string `json:"key" db:"key"`
	Value     string `json:"value" db:"value"` // Di-parse sebagai JSON string di Go
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type Service struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	IconName    string `json:"icon_name" db:"icon_name"`
	OrderNum    int    `json:"order_num" db:"order_num"`
}

type Skill struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Category   string `json:"category" db:"category"`
	Percentage int    `json:"percentage" db:"percentage"`
	IconURL    string `json:"icon_url" db:"icon_url"`
	OrderNum   int    `json:"order_num" db:"order_num"`
}

// Data gabungan untuk response API Home
type HomeAggregatedData struct {
	Settings []SiteSetting `json:"settings"`
	Services []Service     `json:"services"`
	Skills   []Skill       `json:"skills"`
	Projects []Project     `json:"featured_projects"` // Menggunakan entity Project yang sudah ada
}
