package teamroles

import "errors"

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	Error500Status  = errors.New("internal server error")
)
