package odin

import (
	"bytes"
	"encoding/base64"

	uuid "github.com/satori/go.uuid"
)

func generateAuthorizationCode(cid, uid string) (code string) {
	buf := bytes.NewBufferString(cid)
	buf.WriteString(uid)

	token := uuid.NewV3(uuid.NewV1(), buf.String())

	code = base64.RawURLEncoding.EncodeToString(token.Bytes())
	return
}
