package nobordist

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const getConsult = `{"reference_number": ["3a9485e0-1cfc-46b8-9378-d4ac980c4be0", "c6aa4460-7bec-4519-a3ac-4d92c8217a66"]}`

func TestGetConsult(t *testing.T) {

	cons := &ConsultInput{}
	var bPack = []byte(getConsult)
	var err error
	err = json.Unmarshal(bPack, cons)
	if err != nil {
		spew.Dump(err)
	}

	fmt.Println(Serialize(cons))

	packList, err := client.GetConsult(cons)

	fmt.Println(Serialize(packList))
	assert.NoError(t, err)
}
