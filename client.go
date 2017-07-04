package odin

type Client interface {
	ID() string
	SetID(string)
	Uri() string
	SetUri(string)
	Secret() string
	SetSecret(string)
	Custom() string
	SetCustom(string)
}

type client struct {
	ClientID    string
	RedirectUri string
	SecretValue string
	CustomData  string
}

func NewClient() *client {
	return NewClient()
}

func DefaultClient() *client {
	return new(client)
}

func (c *client) ID() string {
	return ""
}
func (c *client) SetID(string) {

}
func (c *client) Uri() string {
	return ""
}

func (c *client) SetUri(string) {

}
func (c *client) Secret() string {
	return ""
}

func (c *client) SetSecret(string) {

}
func (c *client) Custom() string {
	return ""
}
func (c *client) SetCustom(string) {

}
