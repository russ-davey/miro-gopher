package miro

type OEmbed struct {
	HTML            string `json:"html"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Version         string `json:"version"`
	Type            string `json:"type"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
}

type OEmbedFormat string

const (
	OEmbedFormatJSON OEmbedFormat = "json"
	OEmbedFormatXML  OEmbedFormat = "xml"
)

type OEmbedParams struct {
	// Format Specifies the return format of the response. It complies with the oEmbed standard.
	// Allowed formats: either "json", or "xml".
	Format OEmbedFormat `query:"format"`
	// Referrer The URL pointing to the source of the request.
	// Service providers such as Embedly use it to forward the initial site that triggered the oEmbed request.
	Referrer string `query:"referrer"`
	// MaxWidth The maximum width available to the embed, in pixels.
	MaxWidth int32 `query:"maxwidth"`
	// MaxHeight The maximum height available to the embed, in pixels.
	MaxHeight int32 `query:"maxheight"`
}
