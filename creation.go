package nobordist

const postCreation = Version + "fiscal_payments"

type CreationInput struct {
	Payments []CreationPayments `json:"payments"`
}

type CreationPayments struct {
	CourrierName        string  `json:"courrier_name,omitempty"`
	CourrierCnpj        string  `json:"courrier_cnpj,omitempty"`
	Barcode             string  `json:"barcode,omitempty"`
	Value               float64 `json:"value"`
	PaymentType         string  `json:"payment_type"`
	DirNumber           string  `json:"dir_number,omitempty"`
	CustomerName        string  `json:"customer_name,omitempty"`
	CustomerCpf         string  `json:"customer_cpf,omitempty"`
	CustomerCep         string  `json:"customer_cep,omitempty"`
	TrackingNumber      string  `json:"tracking_number,omitempty"`
	ProductsDescription string  `json:"products_description,omitempty"`
	CustomerState       string  `json:"customer_state,omitempty"`
	DirDate             string  `json:"dir_date,omitempty"`
	MasterNumber        string  `json:"master_number,omitempty"`
}

type CreationErrors struct {
	Barcode             []string `json:"barcode,omitempty"`
	CourrierName        []string `json:"courrier_name,omitempty"`
	ProductsDescription []string `json:"products_description,omitempty"`
	CustomerCpf         []string `json:"customer_cpf,omitempty"`
}

type CreationData struct {
	PaymentType         string         `json:"payment_type"`
	Value               float64        `json:"value"`
	CourrierName        string         `json:"courrier_name,omitempty"`
	CourrierCnpj        string         `json:"courrier_cnpj,omitempty"`
	Barcode             string         `json:"barcode,omitempty"`
	CreationStatus      string         `json:"creation_status"`
	ReferenceNumber     string         `json:"reference_number,omitempty"`
	DirNumber           string         `json:"dir_number,omitempty"`
	CustomerName        string         `json:"customer_name,omitempty"`
	CustomerCpf         string         `json:"customer_cpf,omitempty"`
	CustomerCep         string         `json:"customer_cep,omitempty"`
	MasterNumber        string         `json:"master_number,omitempty"`
	CustomerState       string         `json:"customer_state,omitempty"`
	TrackingNumber      string         `json:"tracking_number,omitempty"`
	ProductsDescription string         `json:"products_description,omitempty"`
	DirDate             string         `json:"dir_date,omitempty"`
	Errors              CreationErrors `json:"errors,omitempty"`
}

type CreationResult struct {
	Data     []CreationData `json:"data"`
	Status   string         `json:"status"`
	Messages []string       `json:"messages"`
}

func (n CreationInput) Error() string {
	panic("implement me")
}

func (c *Client) PostCreation(createInput CreationInput) (*CreationResult, error) {
	creationResult := &CreationResult{}
	err := c.Post(postCreation, createInput, nil, creationResult)
	return creationResult, err
}
