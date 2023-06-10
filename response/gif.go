package response

type GifResponse struct {
	Locale  string    `json:"locale,omitempty"`
	Results []GifResults `json:"results,omitempty"`
	Next    string    `json:"next,omitempty"`
}
type Tinygif struct {
	Dims     []int   `json:"dims,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	URL      string  `json:"url,omitempty"`
	Size     int     `json:"size,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}
type Mp4 struct {
	URL      string  `json:"url,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Size     int     `json:"size,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}
type WebpTransparent struct {
	Size     int     `json:"size,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	URL      string  `json:"url,omitempty"`
}
type Webm struct {
	Duration float64 `json:"duration,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	URL      string  `json:"url,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Size     int     `json:"size,omitempty"`
}
type Tinymp4 struct {
	Size     int     `json:"size,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	URL      string  `json:"url,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}
type Loopedmp4 struct {
	URL      string  `json:"url,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Size     int     `json:"size,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Preview  string  `json:"preview,omitempty"`
}
type Nanomp4 struct {
	Dims     []int   `json:"dims,omitempty"`
	Size     int     `json:"size,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	URL      string  `json:"url,omitempty"`
}
type Nanogif struct {
	Size     int     `json:"size,omitempty"`
	URL      string  `json:"url,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}
type Nanowebm struct {
	Duration float64 `json:"duration,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	URL      string  `json:"url,omitempty"`
	Size     int     `json:"size,omitempty"`
}
type NanowebpTransparent struct {
	Size     int     `json:"size,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	URL      string  `json:"url,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
}
type TinywebpTransparent struct {
	Size     int     `json:"size,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	URL      string  `json:"url,omitempty"`
}
type Mediumgif struct {
	Size     int     `json:"size,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	URL      string  `json:"url,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}
type Tinywebm struct {
	URL      string  `json:"url,omitempty"`
	Dims     []int   `json:"dims,omitempty"`
	Size     int     `json:"size,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Preview  string  `json:"preview,omitempty"`
}
type Gif struct {
	Dims     []int   `json:"dims,omitempty"`
	Preview  string  `json:"preview,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Size     int     `json:"size,omitempty"`
	URL      string  `json:"url,omitempty"`
}
type Media struct {
	Tinygif             Tinygif             `json:"tinygif,omitempty"`
	Mp4                 Mp4                 `json:"mp4,omitempty"`
	WebpTransparent     WebpTransparent     `json:"webp_transparent,omitempty"`
	Webm                Webm                `json:"webm,omitempty"`
	Tinymp4             Tinymp4             `json:"tinymp4,omitempty"`
	Loopedmp4           Loopedmp4           `json:"loopedmp4,omitempty"`
	Nanomp4             Nanomp4             `json:"nanomp4,omitempty"`
	Nanogif             Nanogif             `json:"nanogif,omitempty"`
	Nanowebm            Nanowebm            `json:"nanowebm,omitempty"`
	NanowebpTransparent NanowebpTransparent `json:"nanowebp_transparent,omitempty"`
	TinywebpTransparent TinywebpTransparent `json:"tinywebp_transparent,omitempty"`
	Mediumgif           Mediumgif           `json:"mediumgif,omitempty"`
	Tinywebm            Tinywebm            `json:"tinywebm,omitempty"`
	Gif                 Gif                 `json:"gif,omitempty"`
}
type GifResults struct {
	ID                 string  `json:"id,omitempty"`
	Title              string  `json:"title,omitempty"`
	ContentDescription string  `json:"content_description,omitempty"`
	ContentRating      string  `json:"content_rating,omitempty"`
	H1Title            string  `json:"h1_title,omitempty"`
	Media              []Media `json:"media,omitempty"`
	BgColor            string  `json:"bg_color,omitempty"`
	Created            float64 `json:"created,omitempty"`
	Itemurl            string  `json:"itemurl,omitempty"`
	URL                string  `json:"url,omitempty"`
	Tags               []any   `json:"tags,omitempty"`
	Flags              []any   `json:"flags,omitempty"`
	Shares             int     `json:"shares,omitempty"`
	Hasaudio           bool    `json:"hasaudio,omitempty"`
	Hascaption         bool    `json:"hascaption,omitempty"`
	SourceID           string  `json:"source_id,omitempty"`
	Composite          any     `json:"composite,omitempty"`
}

type GifSearchResponse struct {
	Results []GifResults `json:"results,omitempty"`
	Next    string    `json:"next,omitempty"`
}