package link_service_impl

import "errors"

var (
	ErrLinkAlreadyExists = errors.New("link already exists")
	ErrWrongInputFormat  = errors.New("wrong input format")
	ErrLinkNotFound      = errors.New("link not found")
)
