package odin

type server struct {
	Authorization
}

type Server interface {
	SetValidationFunc(ValidationFunc)
	SetUserHandle(UserCallback)
	GetUser() User
}

func NewServer() Server {
	s := new(server)
	s.SetValidationFunc(defaultVelidation)
	s.SetUserHandle(defaultUser)
	return s
}

func (s *server) GetUser() User {
	if uc != nil {
		return uc(s)
	}
	return nil
}

func (s *server) SetValidationFunc(validationFunc ValidationFunc) {
	SetValidationFunc(validationFunc)
}

func (s *server) SetUserHandle(callback UserCallback) {
	SetUserHandle(callback)
}

func defaultClient(authorization Authorization) Client {
	c := NewClient()
	c.ClientID = "1234"
	c.SecretValue = "1234"
	c.RedirectUri = "localhost:8080"
	return c
}

func defaultVelidation(authorization Authorization) error {
	req := authorization.GetRequest()
	if req == nil {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	c, cb := req[CN_CLIENTID]
	r, rb := req[CN_REDIRECTURI]
	if !cb || !rb || c == "" || r == "" {
		return ERROR_MAP[E_INVALID_REQUEST]
	}

	rt, rtb := req[CN_RESPONSETYPE]
	if !rtb || rt == "" {
		return ERROR_MAP[E_UNSUPPORTED_RESPONSE_TYPE]
	} else if e := ResponseType(rt).Verify(); e != nil {
		return e
	}

	return nil
}

func defaultUser(authorization Authorization) User {
	u := new(user)
	u.id = "user1234"
	return u
}

