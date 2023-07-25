package infinite

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	// HTTP client used to communicate with the OM API.
	client *http.Client

	// Base URL for API requests.
	baseURL string

	// User agent for client
	userAgent string

	// Auth token - passed to org manager as:
	// Authorization: bearer <access_token>
	token string

	// the services this client can work with
	Version *VersionService
	Events  *EventsService
	Venues  *VenuesService
}

type service struct {
	sling *sling.Sling
}

func New(optionFuncs ...optionFunc) *Client {

	client := &Client{
		client:    http.DefaultClient,
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
	}

	for _, option := range optionFuncs {
		option(client)
	}

	s := sling.New().Client(client.client).Base(client.baseURL).Set("Accept", "application/json").Set("User-agent", client.userAgent)

	if client.token != "" {
		s = s.Set("Authorization", "Bearer "+client.token)
	}

	client.Version = &VersionService{s.New().Path("version")}
	client.Events = &EventsService{s.New().Path("v1/events/")}
	client.Venues = &VenuesService{s.New().Path("v1/venues/")}

	return client

}

type optionFunc func(*Client)

// SetAuthToken is a client option for setting the auth token used accessing Org Manager
func SetAuthToken(token string) optionFunc {
	return func(c *Client) {
		c.token = token
	}
}

// SetHttpClient is a client option for setting the *http.Client
func SetHttpClient(httpClient *http.Client) optionFunc {
	return func(c *Client) {
		c.client = httpClient
	}
}

// SetBaseURL is a client option for setting the base url
func SetBaseURL(baseURL string) optionFunc {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) optionFunc {
	return func(c *Client) {
		c.userAgent = fmt.Sprintf("%s %s", ua, c.userAgent)
	}
}
