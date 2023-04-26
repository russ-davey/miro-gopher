package miro

type Role string

const (
	RoleViewer    Role = "viewer"
	RoleCommenter Role = "commenter"
	RoleEditor    Role = "editor"
	RoleCoOwner   Role = "coowner"
	RoleOwner     Role = "owner"
	RoleGuest     Role = "guest"
)

type ShareBoardInvitation struct {
	Emails  []string `json:"emails"`
	Role    Role     `json:"role,omitempty"`
	Message string   `json:"message,omitempty"`
}

type BoardInvitationResponse struct {
	Failed []struct {
		Email  string `json:"email,omitempty"`
		Reason string `json:"reason,omitempty"`
	} `json:"failed,omitempty"`
	Successful int64 `json:"successful,omitempty"`
}
