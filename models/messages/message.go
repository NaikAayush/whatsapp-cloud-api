package models

type Message struct {
	Audio            *MediaObject       `json:"audio,omitempty"`
	Contacts         *ContactsObject    `json:"contacts,omitempty"`
	Context          *ContextObject     `json:"context,omitempty"`
	Document         *MediaObject       `json:"document,omitempty"`
	HSM              interface{}        `json:"hsm,omitempty"`
	Image            *MediaObject       `json:"image,omitempty"`
	Interactive      *InteractiveObject `json:"interactive,omitempty"`
	Location         *LocationObject    `json:"location,omitempty"`
	MessagingProduct string             `json:"messaging_product,omitempty"`
	PreviewURL       bool               `json:"preview_url,omitempty"`
	RecipientType    string             `json:"recipient_type,omitempty"`
	Status           string             `json:"status,omitempty"`
	Sticker          *MediaObject       `json:"sticker,omitempty"`
	Template         *TemplateObject    `json:"template,omitempty"`
	Text             *TextObject        `json:"text,omitempty"`
	To               string             `json:"to"`
	Type             string             `json:"type,omitempty"`
}

type MediaObject struct {
	MediaURL string `json:"media_url,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

type ContactsObject struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type ContextObject struct {
	MessageID string `json:"message_id"`
}

type LocationObject struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type InteractiveObject struct {
	Header interface{} `json:"header,omitempty"`
	Body   interface{} `json:"body,omitempty"`
	Footer interface{} `json:"footer,omitempty"`
	Action interface{} `json:"action,omitempty"`
}

type TemplateObject struct {
	TemplateName string                 `json:"template_name"`
	Components   map[string]interface{} `json:"components"`
}

type TextObject struct {
	Text       string `json:"text"`
	PreviewURL bool   `json:"preview_url,omitempty"`
}
