package odin

import (
	"net/url"
)

type ResponseCode string

//response_type	Required. Must be set to code
//client_id	Required. The client identifier as assigned by the authorization server, when the client was registered.
//redirect_uri	Optional. The redirect URI registered by the client.
//scope	Optional. The possible scope of the request.
//state	Optional (recommended). Any client state that needs to be passed on to the client request URI.
type Authorization struct {
	Request  map[string]string
	Response map[string]string
	Callback ValidateCallback
}

type ValidateCallback func(auth Authorization) error

//code	Required. The authorization code.
//state	Required, if present in request. The same value as sent by the client in the state parameter, if any.
//https://tools.ietf.org/html/rfc6749#section-4.1.2
/*
For example, the authorization server redirects the user-agent by
sending the following HTTP response:

  HTTP/1.1 302 Found
  Location: https://client.example.com/cb?code=SplxlOBeZQQYbYS6WxSbIA
            &state=xyz
*/

const (
	A_ResponseType = "response_type"
	A_ClientID     = "client_id"
	A_RedirectUri  = "redirect_uri"
	A_Scope        = "scope"
	A_State        = "state"
	A_Code         = "code"
)

var (
	AuthorizationRequestList = []string{
		A_ResponseType,
		A_ClientID,
		A_RedirectUri,
		A_Scope,
		A_State,
	}
)

func NewAuthorization(values url.Values) *Authorization {
	auth := new(Authorization)
	auth.Request = make(map[string]string)
	auth.Response = make(map[string]string)
	auth.Callback = DefaultCallback
	if values == nil {
		return auth
	}

	for _, v := range AuthorizationRequestList {
		if values.Get(v) != "" {
			auth.Request[v] = values.Get(v)
		}
	}

	return auth
}

func (a *Authorization) SetCallback(c ValidateCallback) *Authorization {
	a.Callback = c
	return a
}

func DefaultCallback(auth Authorization) error {
	if auth.Request == nil {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	return nil
}

func (a *Authorization) Validate() error {
	return validateClient(a)
}

func validateClient(auth *Authorization) error {
	if c := auth.Callback; c != nil {
		return c(*auth)
	}
	return nil
}

func (a *Authorization) Get(s string) (v string, b bool) {
	if a != nil {
		v, b = (a.Request)[s]
	}
	return
}

type AuthorizationResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
