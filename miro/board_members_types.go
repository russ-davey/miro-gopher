package miro

type ShareBoardInvitation struct {
	Emails  []string `json:"emails"`
	Role    string   `json:"role,omitempty"`
	Message string   `json:"message,omitempty"`
}

type BoardInvitationResponse struct {
	Failed []struct {
		Email  string `json:"email,omitempty"`
		Reason string `json:"reason,omitempty"`
	} `json:"failed,omitempty"`
	Successful int64 `json:"successful,omitempty"`
}
