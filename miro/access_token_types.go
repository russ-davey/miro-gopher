package miro

type AccessToken struct {
	Type         string          `json:"type"`
	Team         BasicEntityInfo `json:"team"`
	CreatedBy    BasicEntityInfo `json:"createdBy"`
	Scopes       []string        `json:"scopes"`
	Organization BasicEntityInfo `json:"organization"`
	User         BasicEntityInfo `json:"user"`
}
