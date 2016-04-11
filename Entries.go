package freckleapi

import (
	"encoding/json"
)

type EntryResponse struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	ApprovedAt                  string  `json:"approved_at"`
	ApprovedBy                  User    `json:"approved_by"`
	ApprovedURL                 string  `json:"approved_url"`
	Billable                    bool    `json:"billable"`
	CreatedAt                   string  `json:"created_at"`
	Date                        string  `json:"date"`
	Description                 string  `json:"description"`
	ID                          float64 `json:"id"`
	Import                      Import  `json:"import"`
	Invoice                     Invoice `json:"invoice"`
	InvoicedAt                  string  `json:"invoiced_at"`
	InvoicedOutsideOfFreckleURL string  `json:"invoiced_outside_of_freckle_url"`
	Minutes                     float64 `json:"minutes"`
	Project                     Project `json:"project"`
	SourceURL                   string  `json:"source_url"`
	Tags                        []Tag   `json:"tags"`
	UnapprovedURL               string  `json:"unapproved_url"`
	UpdatedAt                   string  `json:"updated_at"`
	URL                         string  `json:"url"`
	User                        User    `json:"user"`
}

type Project struct {
	Billable         bool    `json:"billable"`
	BillingIncrement float64 `json:"billing_increment"`
	Color            string  `json:"color"`
	Enabled          bool    `json:"enabled"`
	ID               float64 `json:"id"`
	Name             string  `json:"name"`
	URL              string  `json:"url"`
}

type Invoice struct {
	ID          float64 `json:"id"`
	InvoiceDate string  `json:"invoice_date"`
	Reference   string  `json:"reference"`
	State       string  `json:"state"`
	TotalAmount float64 `json:"total_amount"`
	URL         string  `json:"url"`
}

type Import struct {
	ID  float64 `json:"id"`
	URL string  `json:"url"`
}

type Tag struct {
	Billable      bool    `json:"billable"`
	FormattedName string  `json:"formatted_name"`
	ID            float64 `json:"id"`
	Name          string  `json:"name"`
	URL           string  `json:"url"`
}

type EntryService struct {
	client *Client
}

func NewEntryService(client *Client) *EntryService {
	return &EntryService{client}
}

type EntryListCall struct {
	service *EntryService
	args    map[string]interface{}
}

func (s *EntryService) List() *EntryListCall {
	return &EntryListCall{
		service: s,
		args:    make(map[string]interface{}),
	}
}

// Should return a url, not a string
func (c *EntryListCall) Do() (*[]Entry, error) {
	data, err := c.service.client.run("GET", "entries", nil)
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, 0)

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return &entries, nil
}

func (c *EntryListCall) Users(user_ids []float64) *EntryListCall {
	c.args["user_ids"] = user_ids
	return c
}

//Optional string of a date in ISO 8061 format YYYY-MM-DD
func (c *EntryListCall) From(from string) *EntryListCall {
	c.args["from"] = from
	return c
}

func (c *EntryListCall) To(to string) *EntryListCall {
	c.args["to"] = to
	return c
}
