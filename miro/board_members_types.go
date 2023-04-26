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

type RoleUpdate struct {
	Role Role `json:"role"`
}

type ShareBoardInvitation struct {
	// Emails Email IDs of the users you want to invite to the board. You can invite up to 20 members per call (required).
	Emails []string `json:"emails"`
	// Role of the board member.
	Role Role `json:"role,omitempty"`
	// Message The message that will be sent in the invitation email.
	Message string `json:"message,omitempty"`
}

type BoardInvitationResponse struct {
	Failed []struct {
		Email  string `json:"email,omitempty"`
		Reason string `json:"reason,omitempty"`
	} `json:"failed,omitempty"`
	Successful string `json:"successful,omitempty"`
}

type BoardMember struct {
	BasicEntityInfo
	Role  Role            `json:"role"`
	Links PaginationLinks `json:"links,omitempty"`
}

type ListBoardMembersResponse struct {
	Data   []*BoardMember   `json:"data"`
	Total  int              `json:"total"`
	Size   int              `json:"size"`
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
	Links  *PaginationLinks `json:"links"`
	Type   string           `json:"type"`
}

type BoardMemberSearchParams struct {
	// limit The maximum number of board members to retrieve. Default: 20.
	limit string `query:"limit,omitempty"`
	// offset The (zero-based) offset of the first item in the collection to return. Default: 0.
	offset string `query:"offset,omitempty"`
}
