package odin

import (
	"errors"
	"net/url"
)

type AuthorizationRequest map[string]string

type Authorization interface {
	GetRequest() AuthorizationRequest
	SetRequest(request AuthorizationRequest) Authorization
}

//response_type	Required. Must be set to code
//client_id	Required. The client identifier as assigned by the authorization server, when the client was registered.
//redirect_uri	Optional. The redirect URI registered by the client.
//scope	Optional. The possible scope of the request.
//state	Optional (recommended). Any client state that needs to be passed on to the client request URI.
type authorization struct {
	request  AuthorizationRequest
	response map[string]string
	error    map[string]string
	//clientCallback   ClientCallback
	//userCallback     UserCallback
	//validateCallback ValidateCallback
}

type ValidationFunc func(Authorization) error
type ClientCallback func(Authorization) Client
type UserCallback func(Authorization) User

var (
	vc ValidationFunc
	cc ClientCallback
	uc UserCallback
)

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

	//define an authorization initialize error
	AuthorizationInitializeError = errors.New("Authorization initialize error with some unknown type!")
)

func NewAuthorization(values ...interface{}) Authorization {
	auth := new(authorization)
	auth.request = make(map[string]string)
	auth.response = make(map[string]string)

	a, e := initialize(auth, values)
	if e != nil {
		return nil
	}
	return a
}

/**
convert the oauth2 request fields to AuthorizationRequest
*/
func ParseAuthorizationRequest(values url.Values) AuthorizationRequest {
	ar := make(AuthorizationRequest)
	for _, v := range authorizationRequestList {
		if gv := values.Get(v); gv != "" {
			if gv == CN_REDIRECTURI {
				if qu, e := url.QueryUnescape(gv); e == nil {
					ar[v] = qu
					continue
				}
			}
			ar[v] = gv
		}
	}
	return ar
}

func initialize(authorization Authorization, values ...interface{}) (Authorization, error) {
	var e error

	if values == nil {
		return authorization, nil
	}
	for _, val := range values {
		switch val.(type) {
		case url.Values:
			ParseAuthorizationRequest(val.(url.Values))
		case ValidationFunc:
			SetValidationFunc(val.(ValidationFunc))
		case ClientCallback:
			SetClientCallback(val.(ClientCallback))

		case Client:
		default:
			e = AuthorizationInitializeError

		}
	}
	return authorization, e
}

func (a *authorization) ParseRequest(values url.Values) {
	a.request = ParseAuthorizationRequest(values)
}

func (a *authorization) SetRequest(request AuthorizationRequest) Authorization {
	a.request = request
	return a
}

func (a *authorization) GetRequest() AuthorizationRequest {
	return a.request
}

func SetValidationFunc(callback ValidationFunc) {
	vc = callback
}

func (a *authorization) SetValidationFunc(callback ValidationFunc) {
	SetValidationFunc(callback)
}

func (a *authorization) SetClientCallback(callback ClientCallback) {
	SetClientCallback(callback)
}

func SetClientCallback(callback ClientCallback) {
	cc = callback
}

func (a *authorization) SetUserHandle(callback UserCallback) {
	SetUserHandle(callback)
}

func SetUserHandle(callback UserCallback) {
	uc = callback
}

func (a *authorization) Verification() error {
	if a != nil {
		return validateClient(a)
	}
	return ERROR_MAP[E_UNAUTHORIZED_CLIENT]
}

func validateClient(auth Authorization) error {
	if vc != nil {
		return vc(auth)
	}
	return nil
}
