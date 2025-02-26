package link_service_impl

import "errors"

var (
	LinkAlreadyExists = errors.New("link already exists or wrong input")
	LinkNotFound      = errors.New("link not found")
)
