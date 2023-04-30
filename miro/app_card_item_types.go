package miro

import "time"

type Fields struct {
	// IconShape The shape of the icon on the preview field.
	IconShape string `json:"iconShape"`
	// FillColor Hex value representing the color that fills the background area of the preview field, when it's displayed on the app card.
	FillColor string `json:"fillColor"`
	// IconURL A valid URL pointing to an image available online. The transport protocol must be HTTPS. Possible image file formats: JPG/JPEG, PNG, SVG.
	IconUrl string `json:"iconUrl"`
	// TextColor Hex value representing the color of the text string assigned to value.
	TextColor string `json:"textColor"`
	// Tooltip A short text displayed in a tooltip when clicking or hovering over the preview field.
	Tooltip string `json:"tooltip"`
	// Value The actual data value of the custom field. It can be any type of information that you want to convey.
	Value string `json:"value"`
}

type AppCardItemData struct {
	// Fields Array where each object represents a custom preview field. Preview fields are displayed on the bottom half of the app card in the compact view.
	Fields []Fields `json:"fields"`
	//Status indicating whether an app card is connected and in sync with the source. When the source for the app card is deleted, the status returns disabled.
	Status Status `json:"status"`
	// Title A short text header to identify the app card.
	Title string `json:"title"`
	// Description A short text description to add context about the app card.
	Description string `json:"description"`
}

type SetAppCardItem struct {
	// Data Contains app card item data, such as the title, description, or fields.
	Data AppCardItemData `json:"data"`
	// Style Contains information about the style of an app card item, such as the fill color.
	Style Style `json:"style"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and theorigin of the x and y coordinates.
	Position PositionSet `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry Geometry `json:"geometry"`
	// Parent Contains information about the parent this item attached to. Passing null for ID will attach widget to the canvas directly.
	Parent ParentSet `json:"parent"`
}

type Status string

const (
	StatusConnected    Status = "connected"
	StatusDisconnected Status = "disconnected"
	StatusDisabled     Status = "disabled"
)

type AppCardItem struct {
	ID   string `json:"id"`
	Data struct {
		Description string   `json:"description"`
		Fields      []Fields `json:"fields"`
		Owned       bool     `json:"owned"`
		Status      string   `json:"status"`
		Title       string   `json:"title"`
	} `json:"data"`
	Style      Style           `json:"style"`
	Position   Position        `json:"position"`
	Geometry   Geometry        `json:"geometry"`
	CreatedAt  time.Time       `json:"createdAt"`
	CreatedBy  BasicEntityInfo `json:"createdBy"`
	ModifiedAt time.Time       `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo `json:"modifiedBy"`
	Parent     *Parent         `json:"parent,omitempty"`
	Links      Links           `json:"links"`
	Type       string          `json:"type"`
}
