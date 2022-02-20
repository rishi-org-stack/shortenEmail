package mail

import "os"
import "errors"

type gmailClient struct {
	clientID     string
	clientSecret string
	Url          string
}

func NewGmailClient(url string) (*gmailClient, error) {

	if url == "" {
		return nil, errors.New("NewgmailClient: empty url provided")
	}

	clientID := os.Getenv("CLIENT_ID")

	if clientID == "" {
		return nil, errors.New("NewgmailClient: no client id found")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientSecret == "" {
		return nil, errors.New("NewgmailClient: no client secret found")
	}

	return &gmailClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		Url:          url,
	}, nil
}
