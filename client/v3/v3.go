package v3

import (
	"github.com/terraform-providers/terraform-provider-nutanix/client"
)

const (
	libraryVersion = "v3"
	absolutePath   = "api/nutanix/" + libraryVersion
	userAgent      = "nutanix/" + libraryVersion
)

// Client manages the V3 API
type Client struct {
	client *client.Client
	V3     Service
}

// NewV3Client return a client to operate V3 resources
func NewV3Client(credentials client.Credentials) (*Client, error) {
	var baseClient *client.Client

	// check if all required fields are present. Else create an empty client
	if credentials.Username != "" && credentials.Password != "" && credentials.Endpoint != "" {
		c, err := client.NewClient(&credentials, userAgent, absolutePath, false)
		if err != nil {
			return nil, err
		}
		baseClient = c
	} else {
		baseClient = &client.Client{UserAgent: userAgent, Error: "nutanix pc Client is missing. " +
			"Please provide required details - username, password & endpoint" +
			" in provider configuration."}
	}

	f := &Client{
		client: baseClient,
		V3: Operations{
			client: baseClient,
		},
	}

	// f.client.OnRequestCompleted(func(req *http.Request, resp *http.Response, v interface{}) {
	// 	if v != nil {
	// 		utils.PrintToJSON(v, "[Debug] FINISHED REQUEST")
	// 		// TBD: How to print responses before all requests.
	// 	}
	// })

	return f, nil
}
