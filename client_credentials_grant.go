package odin

var (
	/*
		grant_type	Required. Must be set to client_credentials .
		scope	Optional. The scope of the authorization.
	*/
	clientCredentialsGrantRequestList = []string{}

	/*
		access_token property is the access token as assigned by the authorization server.
		token_type property is a type of token assigned by the authorization server.
		expires_in property is a number of seconds after which the access token expires, and is no longer valid. Expiration of access tokens is optional.
		A refresh token should not be included for this type of authorization request.
	*/
	clientCredentialsGrantRequestResponseList = []string{}
)
