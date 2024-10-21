package nobordist

import "time"

const getCreation = Version + "fiscal_payments"

type ConsultInput struct {
	ReferenceNumbers []string `json:"reference_number,omitempty"`
}

type ConsultFiscalParcel struct {
	Id                  int       `json:"id,omitempty"`
	DirNumber           string    `json:"dir_number,omitempty"`
	CustomerName        string    `json:"customer_name,omitempty"`
	CustomerCpf         string    `json:"customer_cpf,omitempty"`
	CustomerCep         string    `json:"customer_cep,omitempty"`
	TrackingNumber      string    `json:"tracking_number,omitempty"`
	ProductsDescription string    `json:"products_description,omitempty"`
	CustomerState       string    `json:"customer_state,omitempty"`
	SellerId            int       `json:"seller_id,omitempty"`
	DirDate             string    `json:"dir_date,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	PlatformId          int       `json:"platform_id,omitempty"`
	MasterNumber        string    `json:"master_number,omitempty"`
}

type ConsultData struct {
	Ids      []int             `json:"ids,omitempty"`
	Count    int               `json:"count,omitempty"`
	Elements []ConsultElements `json:"elements,omitempty"`
}

type ConsultElements struct {
	Id              int                  `json:"id,omitempty"`
	ReferenceNumber string               `json:"reference_number,omitempty"`
	Barcode         string               `json:"barcode,omitempty"`
	Authentication  interface{}          `json:"authentication,omitempty"`
	Status          string               `json:"status,omitempty"`
	PaymentType     string               `json:"payment_type,omitempty"`
	Value           float64              `json:"value,omitempty"`
	FiscalParcelId  int                  `json:"fiscal_parcel_id,omitempty"`
	PaymentDate     interface{}          `json:"payment_date,omitempty"`
	CreatedAt       time.Time            `json:"created_at,omitempty"`
	UpdatedAt       time.Time            `json:"updated_at,omitempty"`
	PlatformId      int                  `json:"platform_id,omitempty"`
	LogSent         bool                 `json:"log_sent,omitempty"`
	Messages        interface{}          `json:"messages,omitempty"`
	CourrierName    string               `json:"courrier_name,omitempty"`
	CourrierCnpj    string               `json:"courrier_cnpj,omitempty"`
	SellerId        int                  `json:"seller_id,omitempty"`
	SellerName      string               `json:"seller_name,omitempty"`
	FiscalParcel    *ConsultFiscalParcel `json:"fiscal_parcel,omitempty"`
}

type ConsultResponse struct {
	Data     ConsultData `json:"data"`
	Status   string      `json:"status"`
	Messages []string    `json:"messages"`
}

func (n ConsultInput) Error() string {
	panic("implement me")
}

func (c *Client) GetConsult(consultInput *ConsultInput) (*ConsultResponse, error) {
	creationResult := &ConsultResponse{}
	var err error

	err = c.Get(getCreation, consultInput, nil, creationResult)

	return creationResult, err
}
