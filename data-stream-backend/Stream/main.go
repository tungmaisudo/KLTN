package Stream

import (
	"errors"

	getstream "github.com/GetStream/stream-go"
)

// Client holds our connection to Stream
var Client *getstream.Client

// Connect -- connect to Stream, set our Client variable or report error
// params:
// apiKey    - string, Stream API key
// apiSecret - string, Stream API Secret
// apiRegion - string, Stream region (ie, "us-east", "eu-central"
func Connect(apiKey string, apiSecret string, apiRegion string, appID string) error {
	var err error
	if apiKey == "" || apiSecret == "" || apiRegion == "" {
		return errors.New("Config not complete")
	}

	Client, err = getstream.New(&getstream.Config{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Location:  apiRegion,
		AppID:     appID,
	})
	return err
}
