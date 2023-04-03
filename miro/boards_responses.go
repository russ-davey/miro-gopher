package miro

import "time"

type CreateBoard struct {
	// Description of the board.
	Description string `json:"description"`
	// Name for the board.
	Name string `json:"name"`
	// Policy Defines the permissions policies and sharing policies for the board.
	Policy Policy `json:"policy"`
	// TeamID Unique identifier (ID) of the team where the board must be placed.
	TeamID string `json:"teamId"`
}

type Board struct {
	ID                    string                `json:"id"`
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	Team                  BasicTeamInfo         `json:"team"`
	Picture               Picture               `json:"picture"`
	Policy                Policy                `json:"policy"`
	ViewLink              string                `json:"viewLink"`
	Owner                 BasicUserInfo         `json:"owner"`
	CurrentUserMembership CurrentUserMembership `json:"currentUserMembership"`
	CreatedAt             time.Time             `json:"createdAt"`
	CreatedBy             BasicUserInfo         `json:"createdBy"`
	ModifiedAt            time.Time             `json:"modifiedAt"`
	ModifiedBy            BasicUserInfo         `json:"modifiedBy"`
	Links                 BoardLinks            `json:"links"`
	Type                  string                `json:"type"`
	Project               Project               `json:"project,omitempty"`
}

type ListBoards struct {
	Data   []*BoardData     `json:"data"`
	Links  *ListBoardsLinks `json:"links"`
	Type   string           `json:"type"`
	Total  int              `json:"total"`
	Size   int              `json:"size"`
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
}

type BoardData struct {
	ID                    string                `json:"id"`
	Type                  string                `json:"type"`
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	Links                 BoardLinks            `json:"links"`
	CreatedAt             time.Time             `json:"createdAt"`
	CreatedBy             BasicUserInfo         `json:"createdBy"`
	CurrentUserMembership CurrentUserMembership `json:"currentUserMembership,omitempty"`
	ModifiedAt            time.Time             `json:"modifiedAt"`
	ModifiedBy            BasicUserInfo         `json:"modifiedBy"`
	Owner                 BasicUserInfo         `json:"owner"`
	PermissionsPolicy     PermissionsPolicy     `json:"permissionsPolicy"`
	Policy                Policy                `json:"policy"`
	Project               Project               `json:"project,omitempty"`
	SharingPolicy         SharingPolicy         `json:"sharingPolicy"`
	Team                  BasicTeamInfo         `json:"team"`
	ViewLink              string                `json:"viewLink"`
}

type BoardLinks struct {
	Related string `json:"related"`
	Self    string `json:"self"`
}

type ListBoardsLinks struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Self  string `json:"self,omitempty"`
}

type SharingPolicy struct {
	// Access Defines the public-level access to the board.
	// Valid options: private | view | edit | comment
	Access string `json:"access"`
	// InviteToAccountAndBoardLinkAccess Defines the user role when inviting a user via the invite to team and board link.
	// For Enterprise users, this parameter is always set to no_access regardless of the value that you assign for this parameter.
	// Valid options: viewer | commenter | editor | no_access
	InviteToAccountAndBoardLinkAccess string `json:"inviteToAccountAndBoardLinkAccess"`
	// OrganizationAccess Defines the organization-level access to the board. If the board is created for a team that does
	// not belong to an organization, the organizationAccess parameter is always set to the default value.
	// Valid options: private | view | edit | comment
	OrganizationAccess string `json:"organizationAccess"`
	// TeamAccess Defines the team-level access to the board.
	// Valid options: private | view | edit | comment
	TeamAccess string `json:"teamAccess"`
}

// PermissionsPolicy Defines the permissions policies for the board.
type PermissionsPolicy struct {
	// CollaborationToolsStartAccess Defines who can start or stop timer, voting, video chat, screen sharing, attention management.
	// Others will only be able to join. To change the value for the collaborationToolsStartAccess parameter, contact Miro Customer Support.
	// Valid options: all_editors | board_owners_and_coowners
	CollaborationToolsStartAccess string `json:"collaborationToolsStartAccess"`
	// CopyAccess Defines who can copy the board, copy objects, download images, and save the board as a template or PDF.
	// Valid options: anyone | team_members | team_editors | board_owner
	CopyAccess string `json:"copyAccess"`
	// CopyAccessLevel ...
	CopyAccessLevel string `json:"copyAccessLevel"`
	// SharingAccess Defines who can change access and invite users to this board. To change the value for the sharingAccess
	// parameter, contact Miro Customer Support.
	// Valid options: team_members_with_editing_rights | board_owners_and_coowners
	SharingAccess string `json:"sharingAccess"`
}

type Policy struct {
	// PermissionsPolicy Defines the permissions policies for the board.
	PermissionsPolicy `json:"permissionsPolicy"`
	// SharingPolicy Defines the public-level, organization-level, and team-level access for the board. The access level
	// that a user gets depends on the highest level of access that results from considering the public-level, team-level,
	// organization-level, and direct sharing access.
	SharingPolicy `json:"sharingPolicy"`
}

type Project struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
