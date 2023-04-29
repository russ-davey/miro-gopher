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
	ID                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Team                  BasicEntityInfo        `json:"team"`
	Picture               *Picture               `json:"picture,omitempty"`
	Policy                Policy                 `json:"policy"`
	PermissionsPolicy     PermissionsPolicy      `json:"permissionsPolicy,omitempty"`
	SharingPolicy         SharingPolicy          `json:"sharingPolicy"`
	ViewLink              string                 `json:"viewLink"`
	Owner                 BasicEntityInfo        `json:"owner"`
	CurrentUserMembership *CurrentUserMembership `json:"currentUserMembership,omitempty"`
	CreatedAt             time.Time              `json:"createdAt"`
	CreatedBy             BasicEntityInfo        `json:"createdBy"`
	ModifiedAt            time.Time              `json:"modifiedAt"`
	ModifiedBy            BasicEntityInfo        `json:"modifiedBy"`
	Links                 Links                  `json:"links"`
	Type                  string                 `json:"type"`
	Project               *Project               `json:"project,omitempty"`
}

type ListBoards struct {
	Data   []*Board         `json:"data"`
	Total  int              `json:"total"`
	Size   int              `json:"size"`
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
	Links  *PaginationLinks `json:"links"`
	Type   string           `json:"type"`
}

type (
	Access       string
	InviteAccess string
)

const (
	AccessPrivate Access = "private"
	AccessView    Access = "view"
	AccessEdit    Access = "edit"
	AccessComment Access = "comment"

	InviteAccessNoAccess  InviteAccess = "no_access"
	InviteAccessViewer    InviteAccess = "viewer"
	InviteAccessCommenter InviteAccess = "commenter"
	InviteAccessEditor    InviteAccess = "editor"
)

type SharingPolicy struct {
	// Access Defines the public-level access to the board.
	// Valid options: private | view | edit | comment
	Access Access `json:"access,omitempty"`
	// InviteToAccountAndBoardLinkAccess Defines the user role when inviting a user via the invite to team and board link.
	// For Enterprise users, this parameter is always set to no_access regardless of the value that you assign for this parameter.
	// Valid options: viewer | commenter | editor | no_access
	InviteToAccountAndBoardLinkAccess InviteAccess `json:"inviteToAccountAndBoardLinkAccess,omitempty"`
	// OrganizationAccess Defines the organization-level access to the board. If the board is created for a team that does
	// not belong to an organization, the organizationAccess parameter is always set to the default value.
	// Warning: may result in a "One of the requested features is not supported. (4.0408)" error message if you don't have the necessary access level.
	// Valid options: private | view | edit | comment
	OrganizationAccess Access `json:"organizationAccess,omitempty"`
	// TeamAccess Defines the team-level access to the board.
	// Valid options: private | view | edit | comment
	TeamAccess Access `json:"teamAccess,omitempty"`
}

type CollabAccess string

type CopyAccess string

type SharingAccess string

const (
	CollabAccessAllEditors             CollabAccess = "all_editors"
	CollabAccessBoardOwnersAndCoOwners CollabAccess = "board_owners_and_coowners"

	CopyAccessAnyone      CopyAccess = "anyone"
	CopyAccessTeamMembers CopyAccess = "team_members"
	CopyAccessTeamEditors CopyAccess = "team_editors"
	CopyAccessBoardOwner  CopyAccess = "board_owner"

	SharingAccessTeamMemberWithEditingRights SharingAccess = "team_members_with_editing_rights"
	SharingAccessOwnersAndCoOwners           SharingAccess = "owner_and_coowners"
)

// PermissionsPolicy Defines the permissions policies for the board.
type PermissionsPolicy struct {
	// CollaborationToolsStartAccess Defines who can start or stop timer, voting, video chat, screen sharing, attention management.
	// Others will only be able to join. To change the value for the collaborationToolsStartAccess parameter, contact Miro Customer Support.
	// Valid options: all_editors | board_owners_and_coowners
	CollaborationToolsStartAccess CollabAccess `json:"collaborationToolsStartAccess,omitempty"`
	// CopyAccess Defines who can copy the board, copy objects, download images, and save the board as a template or PDF.
	// Valid options: anyone | team_members | team_editors | board_owner
	CopyAccess CopyAccess `json:"copyAccess,omitempty"`
	// CopyAccessLevel ...
	CopyAccessLevel string `json:"copyAccessLevel,omitempty"`
	// SharingAccess Defines who can change access and invite users to this board. To change the value for the sharingAccess
	// parameter, contact Miro Customer Support.
	// Valid options: team_members_with_editing_rights | board_owners_and_coowners
	SharingAccess SharingAccess `json:"sharingAccess,omitempty"`
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

type Sort string

const (
	// SortDefault If team_id is present, last_created. Otherwise, last_opened
	SortDefault Sort = "default"
	// SortLastModified sort by the date and time when the board was last modified
	SortLastModified Sort = "last_modified"
	// SortLastOpened sort by the date and time when the board was last opened
	SortLastOpened Sort = "last_opened"
	// SortLastCreated sort by the date and time when the board was created
	SortLastCreated Sort = "last_created"
	// SortAlphabetically sort by the board name (alphabetically)
	SortAlphabetically Sort = "alphabetically"
)

type BoardSearchParams struct {
	// TeamID The team_id for which you want to retrieve the list of boards. If this parameter is sent in the request,
	// the query and owner parameters are ignored.
	TeamID string `query:"team_id,omitempty"`
	// Query Retrieves a list of boards that contain the query string provided in the board content or board name.
	// For example, if you want to retrieve a list of boards that contain the word beta within the board itself (board content),
	// add beta as the query parameter value. You can use the query parameter with the owner parameter to narrow down the board search results.
	Query string `query:"query,omitempty"`
	// Owner Retrieves a list of boards that belong to a specific owner ID. You must pass the owner ID (for example,
	//3074457353169356300), not the owner name. You can use the 'owner' parameter with the query parameter to narrow
	//down the board search results. Note that if you pass the team_id in the same request, the owner parameter is ignored.
	Owner string `query:"owner,omitempty"`
	// Limit The maximum number of boards to retrieve.
	// Default: 20
	Limit string `query:"limit,omitempty"`
	// The (zero-based) offset of the first item in the collection to return.
	// Default: 0.
	Offset string `query:"offset,omitempty"`
	// Sort The order in which you want to view the result set.
	// Options last_created and alphabetically are applicable only when you search for boards by team.
	Sort Sort `query:"sort,omitempty"`
}

type CurrentUserMembership struct {
	BasicEntityInfo
	Role string `json:"role"`
}

type Picture struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageURL"`
}
