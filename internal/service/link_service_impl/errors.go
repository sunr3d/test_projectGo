package link_service_impl

import "errors"

var (
	LinkAlreadyExists = errors.New("link already exists")
	LinkNotFound      = errors.New("link not found")
)
