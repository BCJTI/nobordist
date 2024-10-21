package nobordist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient("phx@phx.com",
		"OK278s.LLA23h",
		"eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo0NzYsImV4cCI6MTcyOTYxNzgyNX0.mtNxIpQyJaoB6sDIzZsGnPPAJrq8bwn3MSvOZ_6-QOg",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("phx@phx.com", "OK278s.LLA23h", true)

	// fmt.Println(Serialize(client))
	assert.NoError(t, err)
}
