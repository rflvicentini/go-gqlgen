package net

import (
	"net/http"
	"time"
)

type Options func(cli *http.Client)

func Client(options ...Options) http.Client {
	cli := http.Client{
		Timeout: time.Second * 5,
	}

	for _, opt := range options {
		opt(&cli)
	}

	return cli
}

func WithClientTimeout(t time.Duration) Options {
	return func(cli *http.Client) {
		cli.Timeout = t
	}
}

func WithRoundTripper(roundTriper http.RoundTripper) Options {
	return func(cli *http.Client) {
		cli.Transport = roundTriper
	}
}