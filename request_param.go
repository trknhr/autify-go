package autify

import (
	"net/url"
	"strconv"
)

// Page sets page number that a request should returns
func Page(page *int) RequestOption {
	return func(o *url.Values) {
		if page != nil {
			o.Set("page", strconv.Itoa(*page))
		}
	}
}

func PerPage(perPage *int) RequestOption {
	return func(o *url.Values) {
		if perPage != nil {
			o.Set("per_page", strconv.Itoa(*perPage))
		}
	}
}

func TestPlanID(testPlanID *int) RequestOption {
	return func(o *url.Values) {
		if testPlanID != nil {
			o.Set("test_plan_id", strconv.Itoa(*testPlanID))
		}
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
