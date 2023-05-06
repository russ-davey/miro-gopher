package miro

type TagColor string

type Tag struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	FillColor string `json:"fillColor"`
	Links     Links  `json:"links"`
	Type      string `json:"type"`
}

type ListTags struct {
	Tags []Tag `json:"tags"`
}

type TagSet struct {
	// Title Text of the tag. Case-sensitive. Must be unique. (required)
	Title string `json:"title"`
	// FillColor Fill color for the tag.
	FillColor TagColor `json:"fillColor"`
}

type TagSearchParams struct {
	// Limit The maximum number of items that can be returned for a single request.
	// Default: 20. Minimum 1. Maximum 50.
	Limit string `query:"limit,omitempty"`
	// Offset The displacement of the first item in the collection to return.
	// Default: 0.
	Offset string `query:"offset,omitempty"`
}

type ListBoardTags struct {
	Data   []Tag           `json:"data"`
	Total  int             `json:"total"`
	Size   int             `json:"size"`
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
	Links  PaginationLinks `json:"links"`
	Type   string          `json:"type"`
}

const (
	TagColorRed        TagColor = "red"
	TagColorLightGreen TagColor = "light_green"
	TagColorCyan       TagColor = "cyan"
	TagColorYellow     TagColor = "yellow"
	TagColorMagenta    TagColor = "magenta"
	TagColorGreen      TagColor = "green"
	TagColorBlue       TagColor = "blue"
	TagColorGray       TagColor = "gray"
	TagColorViolet     TagColor = "violet"
	TagColorDarkGreen  TagColor = "dark_green"
	TagColorDarkBlue   TagColor = "dark_blue"
	TagColorBlack      TagColor = "black"
)
