package miro

import "fmt"

const (
	// EndpointOAUTHToken /oauth-token endpoint
	EndpointOAUTHToken = "oauth-token"
	// EndpointOAUTH /oauth endpoint
	EndpointOAUTH = "oauth"
)

type AccessTokenService struct {
	client     *Client
	APIVersion string
}

// Get information about an access token, such as the token type, scopes, team, user, token creation date and time, and the user who created the token.
func (a *AccessTokenService) Get() (*AccessToken, error) {
	response := &AccessToken{}

	url := fmt.Sprintf("%s/%s/%s", a.client.BaseURL, a.APIVersion, EndpointOAUTHToken)
	err := a.client.Get(url, response)

	return response, err
}

// Revoke Revoking an access token means that the access token will no longer work. When an access token is revoked,
// the refresh token is also revoked and no longer valid. This does not uninstall the application for the user.
func (a *AccessTokenService) Revoke(accessToken string) error {
	url := fmt.Sprintf("%s/%s/%s/revoke", a.client.BaseURL, a.APIVersion, EndpointOAUTH)

	return a.client.postNoContent(url, Parameter{
		"access_token": accessToken,
	})
}
