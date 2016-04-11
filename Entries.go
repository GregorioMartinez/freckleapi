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
	data, err := c.service.client.run("GET", "entries", c.args)
	if err != nil && err != ErrNoMorePages {
		return nil, err
	}

	// @TODO This is ugly as fuck, but this gives us the last page error if it happened.
	oerr := err

	entries := make([]Entry, 0)

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return &entries, oerr
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

func (c *EntryListCall) Page(page int) *EntryListCall {
	c.args["page"] = page
	return c
}

func (c *EntryListCall) Sort(sort string) *EntryListCall {
	c.args["sort"] = sort
	return c
}
