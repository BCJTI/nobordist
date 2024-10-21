package nobordist

const tokenAuth = "authenticate"

type TokenResponse struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Profile        string `json:"profile,omitempty"`
	WhatsappNumber string `json:"whatsapp_number,omitempty"`
	GroupName      string `json:"group_name,omitempty"`
	Agreement      bool   `json:"agreement,omitempty"`
	AuthToken      string `json:"auth_token,omitempty"`
	CompanyName    string `json:"company_name,omitempty"`
	Language       string `json:"language,omitempty"`
	IsSuper        bool   `json:"is_super,omitempty"`
	Platform       string `json:"platform,omitempty"`
	PlatformId     int    `json:"platform_id,omitempty"`
	CanOverpack    bool   `json:"can_overpack,omitempty"`
}

func (n TokenResponse) Error() string {
	panic("implement me")
}
