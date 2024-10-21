package nobordist

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonCreation = `{
    "payments": [
        {
            "courrier_name": "Jack",
            "courrier_cnpj": "78978387",
            "barcode": "02920939089289292",
            "value": 333.21,
            "payment_type": "darf"
        },
        {
            "dir_number": "776534231",
            "customer_name": "Fulano Beltrano",
            "customer_cpf": "50975384040",
            "customer_cep": "05410700",
            "tracking_number": "NX123456789NX",
            "products_description": "cotton nike tshirts, small 200mA charger, plastic screen protector",
            "customer_state": "SP",
            "dir_date": "2024-04-23",
            "value": 77.42,
            "master_number": "8937899783",
            "payment_type": "icms"
        }
    ]
}`

func TestPostCreation(t *testing.T) {

	creation := CreationInput{}
	var bPack = []byte(jsonCreation)
	var err error
	err = json.Unmarshal(bPack, &creation)
	if err != nil {
		spew.Dump(err)
	}

	fmt.Println(Serialize(creation))

	packList, err := client.PostCreation(creation)

	fmt.Println(Serialize(packList))
	assert.NoError(t, err)
}
