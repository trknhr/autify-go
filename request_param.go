package autify

import (
	"net/url"
	"strconv"
)

func Page(page int) RequestOption {
	return func(o *url.Values) {
		o.Set("page", strconv.Itoa(page))
	}
}

type RequestOption func(*url.Values)

func buildQuery(options ...RequestOption) string {
	res := url.Values{}
	for _, opt := range options {
		opt(&res)
	}

	return res.Encode()
}
