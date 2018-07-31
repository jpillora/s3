package s3

import (
	"errors"
	"net/http"
)

func NewRequest(method string, options ...Option) (*http.Request, error) {
	input := newInput(method, options)

	if err := input.validate(); err != nil {
		return nil, err
	}

	if err := input.buildClient(); err != nil {
		return nil, err
	}

	aws := input.buildAWSRequest()
	return aws.HTTPRequest, aws.Sign()
}

const (
	GET = "GET"
	PUT = "PUT"
)

var (
	ErrInvalidRequestMethod = errors.New("Invalid method.")
	ErrBucketMissing        = errors.New("Bucket is required.")
	ErrKeyMissing           = errors.New("Key is required.")
	ErrContentMissing       = errors.New("Content is required.")
)