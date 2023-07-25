package infinite

import (
	"fmt"
	"net/http"
)

type Error struct {
	Details  []string
	Response *ErrorResponse
}

func (e *Error) Error() string {
	var str string
	for _, detail := range e.Details {
		str = str + fmt.Sprintf("%s\n", detail)
	}
	return str
}

// create an error with enough info to be useful
func newError(resp *http.Response, slingErr error, apiErr *ErrorResponse) error {

	err := &Error{Response: apiErr}

	if slingErr != nil {
		err.Details = append(err.Details, slingErr.Error())
	}
	if resp != nil {
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			err.Details = append(err.Details, fmt.Sprintf(
				"request: %s; response: %s", resp.Request.URL, resp.Status,
			))
		}
	}
	if err.Error() == "" && len(err.Details) < 1 {
		return nil
	}

	return err
}
