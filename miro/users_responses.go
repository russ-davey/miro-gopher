package miro

type CurrentUserMembership struct {
	BasicUserInfo
	Role string `json:"role"`
}
