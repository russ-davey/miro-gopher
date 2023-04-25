package miro

type CurrentUserMembership struct {
	BasicEntityInfo
	Role string `json:"role"`
}
