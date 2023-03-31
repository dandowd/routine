package common

type RequestService interface {
	ServiceEndpoints() []Endpoint
}

type Endpoint interface {
	Path() string
}
