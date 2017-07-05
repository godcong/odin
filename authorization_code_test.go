package odin

import (
	"log"
	"net/url"
	"reflect"
	"testing"
)

var v = url.Values{
	"111":           []string{"111"},
	"response_type": []string{"222"},
}

func TestNewAuthorizationRequest(t *testing.T) {
	log.Println(NewAuthorization(v).Request)
	log.Println(NewAuthorization(nil).Request)

}

func TestAuthorization_ResponseUri(t *testing.T) {
	auth := NewAuthorization(v)
	log.Println(auth.ResponseUri())


}
