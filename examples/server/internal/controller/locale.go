package controller

import (
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var Locale = &locale{}

type locale struct {
}

func (c *locale) Locales(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

func (c *locale) Save(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}
