package odin

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
)

var (
	/*
	   client_id	Required. The client application's id.
	   client_secret	Required. The client application's client secret .
	   grant_type	Required. Must be set to authorization_code .
	   code	Required. The authorization code received by the authorization server.
	   redirect_uri	Required, if the request URI was included in the authorization request. Must be identical then.
	*/
	TokenRequestList = []string{}

	/*
		access_token property is the access token as assigned by the authorization server.
		token_type property is a type of token assigned by the authorization server.
		expires_in property is a number of seconds after which the access token expires, and is no longer valid. Expiration of access tokens is optional.
		refresh_token property contains a refresh token in case the access token can expire. The refresh token is used to obtain a new access token once the one returned in this response is no longer valid.
	*/
	TokenResponseList = []string{}
)

func AuthorizeGenerateToken(cid, uid string) (code string) {
	buf := bytes.NewBufferString(cid)
	buf.WriteString(uid)

	token := uuid.NewV3(uuid.NewV1(), buf.String())
	code = base64.URLEncoding.EncodeToString(token.Bytes())
	code = TOKEN_PREFIX + strings.ToUpper(strings.TrimRight(code, "="))
	return
}

func AccessGenerateToken(cid, uid string, nano int64, genRefresh bool) (access, refresh string) {
	buf := bytes.NewBufferString(cid)
	buf.WriteString(uid)
	buf.WriteString(strconv.FormatInt(nano, 10))

	access = base64.URLEncoding.EncodeToString(uuid.NewV3(uuid.NewV4(), buf.String()).Bytes())
	access = TOKEN_PREFIX + strings.ToUpper(strings.TrimRight(access, "="))
	if genRefresh {
		refresh = base64.URLEncoding.EncodeToString(uuid.NewV5(uuid.NewV4(), buf.String()).Bytes())
		refresh = TOKEN_PREFIX + strings.ToUpper(strings.TrimRight(refresh, "="))
	}
	return
}
