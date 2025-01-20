package common

type User struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Phone_number string      `json:"phone_number"`
	Addresses    []Addresses `json:"addresses"`
}

type Addresses struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip_code string `json:"zip_code"`
	Country  string `json:"country"`
}
