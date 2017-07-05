//Resource Owner Password Credentials Grant
package odin

var (
	/*
		grant_type	Required. Must be set to password
		username	Required. The username of the resource owner, UTF-8 encoded.
		password	Required. The password of the resource owner, UTF-8 encoded.
		scope		Optional. The scope of the authorization.
	*/
	rOPCGrantRequestList = []string{}
	/*
	   access_token property is the access token as assigned by the authorization server.
	   token_type property is a type of token assigned by the authorization server.
	   expires_in property is a number of seconds after which the access token expires, and is no longer valid. Expiration of access tokens is optional.
	   refresh_token property contains a refresh token in case the access token can expire. The refresh token is used to obtain a new access token once the one returned in this response is no longer valid.
	*/
	rOPCGrantResponseList = []string{}
)
