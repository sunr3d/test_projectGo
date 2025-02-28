package link_service_impl

import "errors"

var (
	ErrLinkAlreadyExists = errors.New("link already exists")
	ErrWrongLinkFormat   = errors.New("wrong link format")
	ErrLinkNotFound      = errors.New("link not found")
)
