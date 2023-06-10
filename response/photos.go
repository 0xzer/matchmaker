package response

type User struct {
	WidthPct   int     `json:"width_pct,omitempty"`
	XOffsetPct int     `json:"x_offset_pct,omitempty"`
	HeightPct  float64 `json:"height_pct,omitempty"`
	YOffsetPct int     `json:"y_offset_pct,omitempty"`
}
type Algo struct {
	WidthPct   float64 `json:"width_pct,omitempty"`
	XOffsetPct float64 `json:"x_offset_pct,omitempty"`
	HeightPct  float64 `json:"height_pct,omitempty"`
	YOffsetPct float64 `json:"y_offset_pct,omitempty"`
}
type Faces struct {
	Algo                  Algo    `json:"algo,omitempty"`
	BoundingBoxPercentage float64 `json:"bounding_box_percentage,omitempty"`
}
type CropInfo struct {
	User                User    `json:"user,omitempty"`
	Algo                Algo    `json:"algo,omitempty"`
	ProcessedByBullseye bool    `json:"processed_by_bullseye,omitempty"`
	UserCustomized      bool    `json:"user_customized,omitempty"`
	Faces               []Faces `json:"faces,omitempty"`
}
type ProcessedFiles struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}
type Assets struct {
	URL       string `json:"url,omitempty"`
	AssetType string `json:"asset_type,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}
type Photos struct {
	ID              string           `json:"id,omitempty"`
	CropInfo        CropInfo         `json:"crop_info,omitempty"`
	URL             string           `json:"url,omitempty"`
	ProcessedFiles  []ProcessedFiles `json:"processedFiles,omitempty"`
	ProcessedVideos []any            `json:"processedVideos,omitempty"`
	FileName        string           `json:"fileName,omitempty"`
	Extension       string           `json:"extension,omitempty"`
	Assets          []Assets         `json:"assets,omitempty"`
	MediaType       string           `json:"media_type,omitempty"`
}
type SelectedInterests struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type UserInterests struct {
	SelectedInterests []SelectedInterests `json:"selected_interests,omitempty"`
}