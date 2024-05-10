package request

import (
	"compress/gzip"
	"errors"
	"net/http"
	"net/url"
)

var (
	ErrNoConfig   = errors.New("missing config")
	ErrMissingURL = errors.New("missing URL")
	ErrStatus     = errors.New("failed to upload: %w")
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Builder shares an HTTP client across request handlers.
type Builder struct {
	Client    Doer
	GzipLevel *int
	URL       string
}

func (b *Builder) With(params map[string]string, headers map[string]string) *Handler {
	values := make(url.Values, len(params))
	for k, v := range params {
		values.Add(k, v)
	}

	u, _ := url.Parse(b.URL)
	u.RawQuery = values.Encode()

	client := b.Client
	if client == nil {
		client = http.DefaultClient
	}

	gzipLevel := gzip.DefaultCompression
	if v := b.GzipLevel; v != nil {
		gzipLevel = *v
	}

	return &Handler{
		URL:       u.String(),
		Headers:   headers,
		GzipLevel: gzipLevel,
		Client:    client,
	}
}