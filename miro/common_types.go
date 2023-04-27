package miro

// BasicEntityInfo info type for different entities (i.e. users, teams & organizations)
type BasicEntityInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type PaginationLinks struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Self  string `json:"self,omitempty"`
}
