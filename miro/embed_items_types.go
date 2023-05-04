package miro

import "time"

type EmbedItemData struct {
	ContentType  string `json:"contentType"`
	Description  string `json:"description"`
	Html         string `json:"html"`
	Mode         string `json:"mode"`
	PreviewUrl   string `json:"previewUrl"`
	ProviderName string `json:"providerName"`
	ProviderUrl  string `json:"providerUrl"`
	Title        string `json:"title"`
	Url          string `json:"url"`
}

type EmbedItem struct {
	ID         string          `json:"id"`
	Data       EmbedItemData   `json:"data"`
	Position   Position        `json:"position"`
	Geometry   Geometry        `json:"geometry"`
	CreatedAt  time.Time       `json:"createdAt"`
	CreatedBy  BasicEntityInfo `json:"createdBy"`
	ModifiedAt time.Time       `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo `json:"modifiedBy"`
	Parent     Parent          `json:"parent"`
	Links      Links           `json:"links"`
	Type       string          `json:"type"`
}

type Mode string

const (
	ModeInline Mode = "inline"
	ModeModal  Mode = "modal"
)

type SetEmbedItemData struct {
	URL        string `json:"url"`
	Mode       Mode   `json:"mode"`
	PreviewUrl string `json:"previewUrl"`
}

type SetEmbedItem struct {
	Data     SetEmbedItemData `json:"data"`
	Position PositionSet      `json:"position"`
	Geometry Geometry         `json:"geometry"`
	Parent   ParentSet        `json:"parent"`
}
