package miro

const (
	// endpointOAUTHToken /oauth-token endpoint
	endpointOAUTHToken = "oauth-token"
	// endpointOAUTH /oauth endpoint
	endpointOAUTH = "oauth"
)

type AccessTokenService struct {
	client     *Client
	apiVersion string
}

// Get information about an access token, such as the token type, scopes, team, user, token creation date and time, and the user who created the token.
func (a *AccessTokenService) Get() (*AccessToken, error) {
	response := &AccessToken{}

	if url, err := constructURL(a.client.BaseURL, a.apiVersion, endpointOAUTHToken); err != nil {
		return response, err
	} else {
		err = a.client.Get(url, response)
		return response, err
	}
}

// Revoke Revoking an access token means that the access token will no longer work. When an access token is revoked,
// the refresh token is also revoked and no longer valid. This does not uninstall the application for the user.
func (a *AccessTokenService) Revoke(accessToken string) error {
	if url, err := constructURL(a.client.BaseURL, a.apiVersion, endpointOAUTH, "revoke"); err != nil {
		return err
	} else {
		err = a.client.postNoContent(url, Parameter{
			"access_token": accessToken,
		})
		return err
	}
}
