package dto

type CustomerResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Zipcode   string `json:"zip_code"`
	BirthDate string `json:"birth_date"`
	Status    string `json:"status"`
}
