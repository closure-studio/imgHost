package consts

import (
	"errors"
)

const (
	FILE_NOT_PROVIDED     = "file not provided"
	UNABLE_TO_READ_FILE   = "unable to read file"
	UNABLE_TO_UPLOAD_FILE = "unable to upload file"
	UPLOAD_SUCCESS        = "upload success"
)

// define error instances
var (
	ErrClientIDEmpty = errors.New("clientID cannot be empty")
	ErrInValidClientId = errors.New("invalid clientID")
)
