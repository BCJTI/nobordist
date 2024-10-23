package nobordist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient("phx@phx.com",
		"OK278s.LLA23h",
		"eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo0NzYsImV4cCI6MTcyOTgwNDI2MX0.fwmlloCSwyZ1U8jspwqKo533qL8wHyGtu1F3vDkY2EM",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("phx@phx.com", "OK278s.LLA23h", true)

	// fmt.Println(Serialize(client))
	assert.NoError(t, err)
}
