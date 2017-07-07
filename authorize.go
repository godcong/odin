package odin

type AuthorizeRequest map[string]string

type Authorize struct {
	ResponseType ResponseType
	Client
}

func NewAuthorize() *Authorize {
	return new(Authorize)
}

func (a *Authorize) SetClient(client Client) {
	a.Client = client

}

func ParseAuthorizeRequest() AuthorizeRequest {

}
