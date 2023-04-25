package miro

// BasicEntityInfo info type for different entities (i.e. users, teams & organizations)
type BasicEntityInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
