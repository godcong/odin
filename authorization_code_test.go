package odin

import (
	"log"
	"net/url"
	"testing"
)

func TestNewAuthorizationRequest(t *testing.T) {
	v := url.Values{
		"111":           []string{"111"},
		"response_type": []string{"222"},
	}
	log.Println(NewAuthorization(v).Request)
	log.Println(NewAuthorization(nil).Request)

}
