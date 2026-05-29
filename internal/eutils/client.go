package eutils

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	database   string
	rettype    string
}

type options struct {
	database *string
	rettype  *string
}

type Option func(options *options) error

func WithNucleotide() Option {
	db := "nucleotide"
	return func(options *options) error {
		options.database = &db
		return nil
	}
}

func WithGenbank() Option {
	returnType := "genbank"
	return func(options *options) error {
		options.rettype = &returnType
		return nil
	}
}


// NewClient -
func NewClient(timeout time.Duration, opts ...Option) Client {
	var options options
	client := Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		database: "protein",
		rettype: "fasta",
	}
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return client
		}
	}
	if options.database != nil {
		client.database = *options.database
	}
	if options.rettype != nil {
		client.rettype = *options.rettype
	}
	return client

}
