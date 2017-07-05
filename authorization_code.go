package odin

import (
	"net/url"
	"strings"
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
	Error    map[string]string
	Callback ValidateCallback
}

type ValidateCallback func(auth Authorization) error

var (
	/*
	   response_type	Required. Must be set to code
	   client_id	Required. The client identifier as assigned by the authorization server, when the client was registered.
	   redirect_uri	Optional. The redirect URI registered by the client.
	   scope	Optional. The possible scope of the request.
	   state	Optional (recommended). Any client state that needs to be passed on to the client request URI.
	*/
	authorizationRequestList = []string{
		CN_RESPONSETYPE,
		CN_CLIENTID,
		CN_REDIRECTURI,
		CN_SCOPE,
		CN_STATE,
	}
	/*
	   code	Required. The authorization code.
	   state	Required, if present in request. The same value as sent by the client in the state parameter, if any.
	*/
	authorizationResponseList = []string{
		CN_CODE,
		CN_STATE,
	}

	authorizationErrorResponse = ErrorResponse{}
)

func NewAuthorization(values url.Values) *Authorization {
	auth := new(Authorization)
	auth.Request = make(map[string]string)
	auth.Response = make(map[string]string)
	auth.Callback = defaultCallback
	if values == nil {
		return auth
	}

	for _, v := range authorizationRequestList {
		if values.Get(v) != "" {
			auth.Request[v] = values.Get(v)
		}
	}

	return auth
}

func (a *Authorization) ParseRequest(values url.Values) {
	for _, v := range authorizationRequestList {
		if values.Get(v) != "" {
			a.Request[v] = values.Get(v)
		}
	}
}

func (a *Authorization) SetCallback(c ValidateCallback) *Authorization {
	a.Callback = c
	return a
}

func defaultCallback(auth Authorization) error {
	if auth.Request == nil {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	return nil
}

func (a *Authorization) Validate() error {
	if a != nil {
		return validateClient(a)
	}
	return ERROR_MAP[E_UNAUTHORIZED_CLIENT]
}

func validateClient(auth *Authorization) error {
	if c := auth.Callback; c != nil {
		return c(*auth)
	}

	return nil
}

func (a *Authorization) GetRequest(s string) (v string, b bool) {
	if a != nil {
		v, b = (a.Request)[s]
	}
	return
}

func (a *Authorization) MakeResponse() {
	if v, b := a.GetRequest(CN_STATE); b {
		a.Response[CN_STATE] = v
	}

}

func (a *Authorization) ResponseUri() string {
	s := []string{}
	for k, v := range a.Response {
		s = append(s, strings.Join([]string{k, v}, "="))
	}

	return strings.Join(s, "&")
}

