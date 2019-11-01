// Documents
// https://api.slack.com/methods/chat.postMessage

package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const (
	ENDPOINT = "https://slack.com/api/chat.postMessage"
)

type RequestParam struct {
	Channel  string `json:"channel"`
	Text     string `json:"text"`
	LinkName bool   `json:"link_names"`
	Username string `json:"username"`
}

type Client struct {
	httpClient *http.Client
	username   string
	token      string // ex: xxxx-xxxxxxxxx-xxxx
	// see: https://api.slack.com/custom-integrations/legacy-tokens
}

func NewClient(username string) (*Client, error) {
	f, err := os.OpenFile(filepath.Join(".", "token", "slack_token.json"), os.O_RDONLY, 0666)
	if err != nil {
		return &Client{}, nil
	}
	token, err := ioutil.ReadAll(f)

	c := &Client{
		httpClient: &http.Client{},
		username:   username,
		token:      string(token),
	}
	return c, nil
}

func (c *Client) PostMessage(test, channel string) (string, error) {
	p := RequestParam{
		Channel:  channel,
		Text:     test,
		LinkName: true,
		Username: c.username,
	}
	return c.postMessage(p)
}

func (c *Client) postMessage(requestParam RequestParam) (string, error) {
	u, err := url.ParseRequestURI(ENDPOINT)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(requestParam)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New("failed to send message, got status code: " + resp.Status + ":" + string(body))
	}

	return fmt.Sprintf("success to send folloing message to slack: \n%s", requestParam.Text), nil
}
