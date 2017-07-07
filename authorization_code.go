package odin

import (
	"bytes"
	"encoding/base64"
	"net/url"
	"strings"

	"errors"

	uuid "github.com/satori/go.uuid"
)

type Authorization interface {
	Request() map[string]string
}

//response_type	Required. Must be set to code
//client_id	Required. The client identifier as assigned by the authorization server, when the client was registered.
//redirect_uri	Optional. The redirect URI registered by the client.
//scope	Optional. The possible scope of the request.
//state	Optional (recommended). Any client state that needs to be passed on to the client request URI.
type authorization struct {
	Request  map[string]string
	Response map[string]string
	Error    map[string]string
	client   Client
	user     User
	//clientCallback   ClientCallback
	//userCallback     UserCallback
	//validateCallback ValidateCallback
}

type ValidateCallback func(Authorization) error
type ClientCallback func(Authorization) Client
type UserCallback func(Authorization) User

var (
	vc ValidateCallback
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
	auth.Request = make(map[string]string)
	auth.Response = make(map[string]string)

	a, e := initialize(auth, values)
	if e != nil {
		return nil
	}
	return a
}

func initialize(authorization Authorization, values ...interface{}) (Authorization, error) {
	var e error
	vc = defaultValidate
	cc = defaultClient
	uc = defaultUser

	if values == nil {
		return authorization, nil
	}
	for _, val := range values {
		switch val.(type) {
		case url.Values:
			for _, v := range authorizationRequestList {
				val := val.(url.Values)
				if gv := val.Get(v); gv != "" {
					if gv == CN_REDIRECTURI {
						if qu, e := url.QueryUnescape(gv); e == nil {
							authorization.Request[v] = qu
							continue
						}
					}
					authorization.Request[v] = gv
				}
			}
		case ValidateCallback:
			SetValidateCallback(val.(ValidateCallback))
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
	for _, v := range authorizationRequestList {
		if values.Get(v) != "" {
			a.Request[v] = values.Get(v)
		}
	}
}

func SetValidateCallback(callback ValidateCallback) {
	vc = callback
}

func (a *authorization) SetValidateCallback(callback ValidateCallback) {
	SetValidateCallback(callback)
}

func (a *authorization) SetClientCallback(callback ClientCallback) {
	SetClientCallback(callback)
}

func SetClientCallback(callback ClientCallback) {
	cc = callback
}

func (a *authorization) SetUserCallback(callback UserCallback) {
	SetUserCallback(callback)
}

func SetUserCallback(callback UserCallback) {
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

func validateUser() {

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

	code := ""
	if v, b := a.GetRequest(CN_CLIENTID); b {
		code = generateAuthorizationCode(v, "")
	}
	a.Response[CN_CODE] = code
}

func (a *Authorization) ResponseUri(other ...map[string]string) string {
	s := []string{}
	a.MakeResponse()
	for k, v := range a.Response {
		s = append(s, strings.Join([]string{k, v}, "="))
	}

	//if size := len(other); size > 0 {
	//	for ; size > 0; size-- {
	//		for k, v := range other[size-1] {
	//			s = append(s, strings.Join([]string{k, v}, "="))
	//		}
	//	}
	//}

	return strings.Join(s, "&")
}

func generateAuthorizationCode(cid, uid string) (code string) {
	buf := bytes.NewBufferString(cid)
	buf.WriteString(uid)

	token := uuid.NewV3(uuid.NewV1(), buf.String())

	code = base64.RawURLEncoding.EncodeToString(token.Bytes())
	return
}

func defaultClient(auth Authorization) Client {
	c := NewClient()
	c.ClientID = "1234"
	c.SecretValue = "1234"
	c.RedirectUri = "localhost:8080"
	return c
}

func defaultValidate(auth Authorization) error {
	if auth.Request == nil {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	return nil
}

func defaultUser(auth Authorization) User {
	if auth.R == nil {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	return nil
}
