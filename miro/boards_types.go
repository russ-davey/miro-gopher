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
	Team                  BasicEntityInfo       `json:"team"`
	Picture               Picture               `json:"picture,omitempty"`
	Policy                Policy                `json:"policy"`
	ViewLink              string                `json:"viewLink"`
	Owner                 BasicEntityInfo       `json:"owner"`
	CurrentUserMembership CurrentUserMembership `json:"currentUserMembership"`
	CreatedAt             time.Time             `json:"createdAt"`
	CreatedBy             BasicEntityInfo       `json:"createdBy"`
	ModifiedAt            time.Time             `json:"modifiedAt"`
	ModifiedBy            BasicEntityInfo       `json:"modifiedBy"`
	Links                 BoardLinks            `json:"links"`
	Type                  string                `json:"type"`
	Project               Project               `json:"project,omitempty"`
}

type ListBoards struct {
	Data   []*Board         `json:"data"`
	Links  *ListBoardsLinks `json:"links"`
	Type   string           `json:"type"`
	Total  int              `json:"total"`
	Size   int              `json:"size"`
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
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
	Access string `json:"access,omitempty"`
	// InviteToAccountAndBoardLinkAccess Defines the user role when inviting a user via the invite to team and board link.
	// For Enterprise users, this parameter is always set to no_access regardless of the value that you assign for this parameter.
	// Valid options: viewer | commenter | editor | no_access
	InviteToAccountAndBoardLinkAccess string `json:"inviteToAccountAndBoardLinkAccess,omitempty"`
	// OrganizationAccess Defines the organization-level access to the board. If the board is created for a team that does
	// not belong to an organization, the organizationAccess parameter is always set to the default value.
	// Warning: may result in a "One of the requested features is not supported. (4.0408)" error message if you don't have the necessary access level.
	// Valid options: private | view | edit | comment
	OrganizationAccess string `json:"organizationAccess,omitempty"`
	// TeamAccess Defines the team-level access to the board.
	// Valid options: private | view | edit | comment
	TeamAccess string `json:"teamAccess,omitempty"`
}

// PermissionsPolicy Defines the permissions policies for the board.
type PermissionsPolicy struct {
	// CollaborationToolsStartAccess Defines who can start or stop timer, voting, video chat, screen sharing, attention management.
	// Others will only be able to join. To change the value for the collaborationToolsStartAccess parameter, contact Miro Customer Support.
	// Valid options: all_editors | board_owners_and_coowners
	CollaborationToolsStartAccess string `json:"collaborationToolsStartAccess,omitempty"`
	// CopyAccess Defines who can copy the board, copy objects, download images, and save the board as a template or PDF.
	// Valid options: anyone | team_members | team_editors | board_owner
	CopyAccess string `json:"copyAccess,omitempty"`
	// CopyAccessLevel ...
	CopyAccessLevel string `json:"copyAccessLevel,omitempty"`
	// SharingAccess Defines who can change access and invite users to this board. To change the value for the sharingAccess
	// parameter, contact Miro Customer Support.
	// Valid options: team_members_with_editing_rights | board_owners_and_coowners
	SharingAccess string `json:"sharingAccess,omitempty"`
}

type Policy struct {
	// PermissionsPolicy Defines the permissions policies for the board.
	PermissionsPolicy `json:"permissionsPolicy,omitempty"`
	// SharingPolicy Defines the public-level, organization-level, and team-level access for the board. The access level
	// that a user gets depends on the highest level of access that results from considering the public-level, team-level,
	// organization-level, and direct sharing access.
	SharingPolicy `json:"sharingPolicy,omitempty"`
}

type Project struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
