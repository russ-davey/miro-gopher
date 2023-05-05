package miro

import "time"

type NoteShape string

const (
	NoteShapeSquare    NoteShape = "square"
	NoteShapeRectangle NoteShape = "rectangle"
)

type StickyNoteData struct {
	Content   string    `json:"content"`
	NoteShape NoteShape `json:"shape"`
}

type StickyNote struct {
	ID         string          `json:"id"`
	Data       StickyNoteData  `json:"data"`
	Style      Style           `json:"style"`
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

type NoteColor string

const (
	NoteColorGray        NoteColor = "gray"
	NoteColorLightYellow NoteColor = "light_yellow"
	NoteColorYellow      NoteColor = "yellow"
	NoteColorOrange      NoteColor = "orange"
	NoteColorLightGreen  NoteColor = "light_green"
	NoteColorGreen       NoteColor = "green"
	NoteColorDarkGreen   NoteColor = "dark_green"
	NoteColorCyan        NoteColor = "cyan"
	NoteColorLightPink   NoteColor = "light_pink"
	NoteColorPink        NoteColor = "pink"
	NoteColorViolet      NoteColor = "violet"
	NoteColorRed         NoteColor = "red"
	NoteColorLightBlue   NoteColor = "light_blue"
	NoteColorBlue        NoteColor = "blue"
	NoteColorDarkBlue    NoteColor = "dark_blue"
	NoteColorBlack       NoteColor = "black"
)

type StickyNoteStyle struct {
	NoteColor         NoteColor         `json:"fillColor"`
	TextAlign         TextAlign         `json:"textAlign"`
	TextAlignVertical TextAlignVertical `json:"textAlignVertical"`
}

type StickyNoteSet struct {
	Data     StickyNoteData  `json:"data"`
	Style    StickyNoteStyle `json:"style"`
	Position PositionSet     `json:"position"`
	Geometry GeometrySet     `json:"geometry"`
	Parent   ParentSet       `json:"parent"`
}
