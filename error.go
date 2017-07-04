package odin

import (
	"encoding/json"
	"errors"
)

const (
	E_INVALID_REQUEST           = iota //"invalid_request"
	E_UNAUTHORIZED_CLIENT              //"unauthorized_client"
	E_ACCESS_DENIED                    //"access_denied"
	E_UNSUPPORTED_RESPONSE_TYPE        //"unsupported_response_type"
	E_INVALID_SCOPE                    //"invalid_scope"
	E_SERVER_ERROR                     //"server_error"
	E_TEMPORARILY_UNAVAILABLE          //"temporarily_unavailable"
	E_UNSUPPORTED_GRANT_TYPE           //"unsupported_grant_type"
	E_INVALID_GRANT                    //"invalid_grant"
	E_INVALID_CLIENT                   //"invalid_client"
	/*----------------------------------------------------*/
	E_INVALID_NONE //no errors
)

var ERROR_MAP = map[int]error{
	E_INVALID_REQUEST:           errors.New("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed."),
	E_UNAUTHORIZED_CLIENT:       errors.New("The client is not authorized to request a token using this method."),
	E_ACCESS_DENIED:             errors.New("The resource owner or authorization server denied the request."),
	E_UNSUPPORTED_RESPONSE_TYPE: errors.New("The authorization server does not support obtaining a token using this method."),
	E_INVALID_SCOPE:             errors.New("The requested scope is invalid, unknown, or malformed."),
	E_SERVER_ERROR:              errors.New("The authorization server encountered an unexpected condition that prevented it from fulfilling the request."),
	E_TEMPORARILY_UNAVAILABLE:   errors.New("The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server."),
	E_UNSUPPORTED_GRANT_TYPE:    errors.New("The authorization grant type is not supported by the authorization server."),
	E_INVALID_GRANT:             errors.New("The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client."),
	E_INVALID_CLIENT:            errors.New("Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method)."),
	E_INVALID_NONE:              nil,
}


/*
For example, the authorization server redirects the user-agent by
sending the following HTTP response:

HTTP/1.1 302 Found
Location: https://client.example.com/cb?error=access_denied&state=xyz
*/
type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
	State            string `json:"state"`
}

func (e *ErrorResponse) Encode() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func Decode(t, s string) interface{} {
	return nil
}

func ErrorResponseDecode(s string) *ErrorResponse {
	v := ErrorResponse{}
	if err := json.Unmarshal(([]byte)(s), &v); err == nil {
		return &v
	}
	return nil
}
