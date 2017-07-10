package odin

type ResponseType string

const (
	Code  ResponseType = "code"
	Token ResponseType = "token"
)

var (
	ResponseTypeList = []ResponseType{
		Code,
		Token,
	}
)

type GrantType string

const (
	AuthorizationCode                GrantType = "authorization_code"
	ResourceOwnerPasswordCredentials GrantType = "password"
	ClientCredentials                GrantType = "client_credentials "
	//Implicit GrantType = "__implicit"
)

type TokenRequest struct {
}
type TokenResponse struct {
}

type ErrorResponseType string

//defined column name to const
const (
	CN_RESPONSETYPE      = "response_type"
	CN_CLIENTID          = "client_id"
	CN_REDIRECTURI       = "redirect_uri"
	CN_SCOPE             = "scope"
	CN_STATE             = "state"
	CN_CODE              = "code"
	CN_ERROR             = "error"
	CN_ERROR_DESCRIPTION = "error_description"
	CN_ERROR_URI         = "error_uri"
)

const (
	//invalid_request
	//The request is missing a required parameter, includes an
	//invalid parameter value, includes a parameter more than
	//once, or is otherwise malformed.
	ErrorTypeInvalidRequest ErrorResponseType = "invalid_request"
	//unauthorized_client
	//The client is not authorized to request an authorization
	//code using this method.
	ErrorTypeUnauthorizedClient ErrorResponseType = "unauthorized_client"
	//access_denied
	//The resource owner or authorization server denied the
	//request.
	ErrorTypeAccessDenied ErrorResponseType = "access_denied"
	//unsupported_response_type
	//The authorization server does not support obtaining an
	//authorization code using this method.
	ErrorTypeUnsupportedResponseType ErrorResponseType = "unsupported_response_type"
	//invalid_scope
	//The requested scope is invalid, unknown, or malformed.
	ErrorTypeInvalidScope ErrorResponseType = "invalid_scope"
	//server_error
	//The authorization server encountered an unexpected
	//condition that prevented it from fulfilling the request.
	//(This error code is needed because a 500 Internal Server
	//Error HTTP status code cannot be returned to the client
	//via an HTTP redirect.)
	ErrorTypeServerError ErrorResponseType = "server_error"
	//temporarily_unavailable
	//The authorization server is currently unable to handle
	//the request due to a temporary overloading or maintenance
	//of the server.  (This error code is needed because a 503
	//Service Unavailable HTTP status code cannot be returned
	//to the client via an HTTP redirect.)
	ErrorTypeTemporarilyUnavailable ErrorResponseType = "temporarily_unavailable"
)

type ErrorMessage struct {
	Code   int
	Detail string
}

//https://tools.ietf.org/html/rfc6749#section-4.1.2.1
func AuthorizationResponseError(e ErrorResponseType) ErrorMessage {
	errors := map[ErrorResponseType]ErrorMessage{
		ErrorTypeInvalidRequest:          ErrorMessage{400, "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed."},
		ErrorTypeUnauthorizedClient:      ErrorMessage{401, ""},
		ErrorTypeAccessDenied:            ErrorMessage{401, ""},
		ErrorTypeUnsupportedResponseType: ErrorMessage{400, ""},
		ErrorTypeInvalidScope:            ErrorMessage{400, ""},
		ErrorTypeServerError:             ErrorMessage{500, ""},
		ErrorTypeTemporarilyUnavailable:  ErrorMessage{503, ""},
	}

	if v, b := errors[e]; b == true {
		return v
	}

	return errors[ErrorTypeInvalidRequest]
}

//ErrInvalidRequest:          400,
//ErrUnauthorizedClient:      401,
//ErrAccessDenied:            403,
//ErrUnsupportedResponseType: 401,
//ErrInvalidScope:            400,
//ErrServerError:             500,
//ErrTemporarilyUnavailable:  503,
//ErrInvalidClient:           401,
//ErrInvalidGrant:            401,
//ErrUnsupportedGrantType:    401,
//
//
//invalid_request
//The request is missing a required parameter, includes an
//unsupported parameter value (other than grant type),
//repeats a parameter, includes multiple credentials,
//utilizes more than one mechanism for authenticating the
//client, or is otherwise malformed.
//
//invalid_client
//Client authentication failed (e.g., unknown client, no
//client authentication included, or unsupported
//authentication method).  The authorization server MAY
//return an HTTP 401 (Unauthorized) status code to indicate
//which HTTP authentication schemes are supported.  If the
//client attempted to authenticate via the "Authorization"
//request header field, the authorization server MUST
//respond with an HTTP 401 (Unauthorized) status code and
//include the "WWW-Authenticate" response header field
//matching the authentication scheme used by the client.
//
//invalid_grant
//The provided authorization grant (e.g., authorization
//code, resource owner credentials) or refresh token is
//invalid, expired, revoked, does not match the redirection
//URI used in the authorization request, or was issued to
//another client.
//
//unauthorized_client
//The authenticated client is not authorized to use this
//authorization grant type.
//
//unsupported_grant_type
//The authorization grant type is not supported by the
//authorization server.
//
//invalid_scope
//The requested scope is invalid, unknown, malformed, or
//exceeds the scope granted by the resource owner.

func (rt ResponseType) Verify() error {
	for _, v := range ResponseTypeList {
		if rt == v {
			return nil
		}
	}
	return ERROR_MAP[E_UNAUTHORIZED_CLIENT]

}
