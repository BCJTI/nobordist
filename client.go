package nobordist

import (
	"fmt"
)

type Client struct {
	Token   string
	Sandbox bool
	User    string
	Pass    string
}

func NewClientToken(user, pass string, sandbox bool) (*Client, error) {

	c := &Client{
		User:    user,
		Pass:    pass,
		Sandbox: sandbox,
	}

	model := new(TokenResponse)
	//{"email": "YOUR_EMAIL", "password": "YOUR_PASSWORD"}
	params := map[string]string{
		"email":    user,
		"password": pass,
	}

	var err error

	err = c.Post(tokenAuth, params, nil, model)

	if err == nil {
		if model.AuthToken != "" {
			c.Token = model.AuthToken
		}
	}

	return c, err
}

func NewClient(user, pass, token string, sandbox bool) *Client {
	return &Client{
		User:    user,
		Pass:    pass,
		Token:   token,
		Sandbox: sandbox,
	}
}

func (c *Client) GetToken() string {
	if c.Token != "" {
		return fmt.Sprintf("Bearer %s", c.Token)
	}
	return ""
}

func (c *Client) GetEndpoint() string {

	if c.Sandbox {
		return SandBoxUrl
	}

	return ProductionUrl

}
