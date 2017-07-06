package odin

type Client interface {
	ID() string
	//SetID(string)
	Uri() string
	//SetUri(string)
	Secret() string
	//SetSecret(string)
	Custom() interface{}
	//SetCustom(string)
}

type client struct {
	ClientID    string
	RedirectUri string
	SecretValue string
	CustomData  string
}

func NewClient() *client {
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
func (c *client) Custom() interface{} {
	return ""
}
func (c *client) SetCustom(interface{}) {

}
