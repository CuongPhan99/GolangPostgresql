package model

type Accounts struct {
	Id              int    `json:"id" pg:"type:serial"`
	Logo            string `json:"logo"`
	CompanyName     string `json:"company_name"`
	AddressCompany  string `json:"address_company"`
	Phone           string `json:"phone"`
	Name            string `json:"name"`
	CoPartner       string `json:"co_partner"`
	IntroduceName   string `json:"seller_name"`
	CheckNumber     string `json:"check_number"`
	OriginerImprint string `json:"originer_imprint"`
}
