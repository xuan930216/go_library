package mcp

type MailchimpCredential struct {
	Country string `json:"country"`
	Key     string `json:"key"`
	ListID  string `json:"List_id"`
	USZone  string `json:"us_zone"`
}

type MailchimpCredentials struct {
	Accounts []*MailchimpCredential `json:"accounts"`
}

func LoadCredentials() *MailchimpCredentials {
	return nil
}
