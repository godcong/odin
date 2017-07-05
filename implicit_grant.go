package odin

var (
	/*
	   response_type	Required. Must be set to token .
	   client_id	Required. The client identifier as assigned by the authorization server, when the client was registered.
	   redirect_uri	Optional. The redirect URI registered by the client.
	   scope	Optional. The possible scope of the request.
	   state	Optional (recommended). Any client state that needs to be passed on to the client request URI.
	*/
	implicitGrantRequestList = []string{}
	/*
	   access_token	Required. The access token assigned by the authorization server.
	   token_type	Required. The type of the token
	   expires_in	Recommended. A number of seconds after which the access token expires.
	   scope	Optional. The scope of the access token.
	   state	Required, if present in the autorization request. Must be same value as state parameter in request.
	*/
	implicitGrantResponseList = []string{}

	implicitGrantErrorResponse = ErrorResponse{}
)
