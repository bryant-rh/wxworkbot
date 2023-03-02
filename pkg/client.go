package pkg

import (
	"fmt"
	"strings"

	"github.com/imroc/req/v3"
)

// GithubClient is the go client for GitHub API.
type CMClient struct {
	*req.Client
	isLogged bool
}

// NewGithubClient create a GitHub client.
func NewReqClient() *CMClient {
	c := req.C().
		// All API requests user the same agent.
		SetUserAgent("reqcli api client").
		//SetCommonHeader("Accept", "application/json").
		// All API requests use the same base URL.
		//SetBaseURL(baseurl).
		// EnableDump at the request level in request middleware which dump content into
		// memory (not print to stdout), we can record dump content only when unexpected
		// exception occurs, it is helpful to troubleshoot problems in production.
		OnBeforeRequest(func(c *req.Client, r *req.Request) error {
			if r.RetryAttempt > 0 { // Ignore on retry.
				return nil
			}
			r.EnableDump()
			return nil
		}).
		// Unmarshal all GitHub error response into struct.
		SetCommonErrorResult(&APIError{}).
		// Handle common exceptions in response middleware.
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if err, ok := resp.ErrorResult().(*APIError); ok {
				// Server returns an error message, convert it to human-readable go error.
				return err
			}
			// Corner case: neither an error response nor a success response,
			// dump content to help troubleshoot.
			if !resp.IsSuccessState() {
				return fmt.Errorf("bad response, raw dump:\n%s", resp.Dump())
			}
			return nil
		})

	return &CMClient{
		Client: c,
	}
}

// APIError represents the error message that GitHub API returns.
// GitHub API doc: https://docs.github.com/en/rest/overview/resources-in-the-rest-api#client-errors
type APIError struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Errors []struct {
		Resource string `json:"resource"`
		Field    string `json:"field"`
		Code     string `json:"code"`
	} `json:"errors,omitempty"`
}

// Error convert APIError to a human readable error and return.
func (e *APIError) Error() string {
	msg := fmt.Sprintf("API Error: code: [%d], Error: %s ", e.Code, e.Msg)
	// if e.DocumentationUrl != "" {
	// 	return fmt.Sprintf("%s (see doc %s)", msg, e.DocumentationUrl)
	// }
	if len(e.Errors) == 0 {
		return msg
	}
	errs := []string{}
	for _, err := range e.Errors {
		errs = append(errs, fmt.Sprintf("resource:%s field:%s code:%s", err.Resource, err.Field, err.Code))
	}
	return fmt.Sprintf("%s (%s)", msg, strings.Join(errs, " | "))
}

// LoginWithToken login with GitHub personal access token.
func (c *CMClient) LoginWithToken(BearerAuthToken string) *CMClient {
	// All API requests need this BearerAuthToken.
	c.SetCommonBearerAuthToken(BearerAuthToken)
	c.isLogged = true
	return c
}

// IsLogged return true is user is logged in, otherwise false.
func (c *CMClient) IsLogged() bool {
	return c.isLogged
}

// SetDebug enable debug if set to true, disable debug if set to false.
func (c *CMClient) SetDebug(enable bool) *CMClient {
	if enable {
		c.EnableDebugLog()
		c.EnableDumpAll()
	} else {
		c.DisableDebugLog()
		c.DisableDumpAll()
	}
	return c
}
