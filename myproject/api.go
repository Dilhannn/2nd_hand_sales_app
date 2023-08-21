// api.go

package main

type Api struct {
	Service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		Service: service,
	}
}
