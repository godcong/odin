package odin

import (
	"log"
	"net/url"
	"testing"
)

var v = url.Values{
	"client_id":     []string{"111"},
	"response_type": []string{"222"},
}

func TestNewAuthorizationRequest(t *testing.T) {
	log.Println(NewAuthorization(v).Request)
	log.Println(NewAuthorization(nil).Request)

}

func TestAuthorization_ResponseUri(t *testing.T) {
	auth := NewAuthorization(v)
	log.Println(auth.ResponseUri(
		map[string]string{
			"1234": "aaaaa",
			"2234": "bbbbb",
		},
		map[string]string{
			"3234": "aaaaa",
			"4234": "bbbbb",
		},
	))

}
