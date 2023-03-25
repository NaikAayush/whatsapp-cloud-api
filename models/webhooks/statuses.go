package models

import "time"

type Status struct {
	Conversation Conversation `json:"conversation"`
	Errors       []Error      `json:"errors"`
	ID           string       `json:"id"`
	Pricing      Pricing      `json:"pricing"`
	RecipientID  string       `json:"recipient_id"`
	Status       string       `json:"status"`
	Timestamp    time.Time    `json:"timestamp"`
}

type Conversation struct {
	ID                  string     `json:"id"`
	Origin              Origin     `json:"origin"`
	ExpirationTimestamp *time.Time `json:"expiration_timestamp,omitempty"`
}

type Origin struct {
	Type               string `json:"type"`
	BusinessInitiated  bool   `json:"business_initiated"`
	CustomerInitiated  bool   `json:"customer_initiated"`
	ReferralConversion bool   `json:"referral_conversion"`
}

type Pricing struct {
	Billable     bool   `json:"billable"`
	Category     string `json:"category"`
	PricingModel string `json:"pricing_model"`
}
