package models

type Value struct {
	Contacts         []Contact `json:"contacts"`
	Errors           []Error   `json:"errors"`
	MessagingProduct string    `json:"messaging_product"`
	Messages         []Message `json:"messages"`
	Metadata         Metadata  `json:"metadata"`
	Statuses         []Status  `json:"statuses"`
}

type Contact struct {
	WaID    string  `json:"wa_id"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"name"`
}
type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}
