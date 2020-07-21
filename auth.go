package rmq

type ExternalAuthentication struct {
}

func (auth *ExternalAuthentication) Mechanism() string {
	return "EXTERNAL"
}

func (auth *ExternalAuthentication) Response() string {
	return ""
}
