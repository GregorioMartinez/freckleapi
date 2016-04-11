package freckleapi

import (
	"encoding/json"
)

type Project struct {
	Billable            bool      `json:"billable"`
	BillableMinutes     float64   `json:"billable_minutes"`
	BillingIncrement    float64   `json:"billing_increment"`
	BudgetMinutes       float64   `json:"budgeted_minutes"`
	Color               string    `json:"color"`
	CreatedAt           string    `json:"created_at"`
	Enabled             bool      `json:"enabled"`
	Entries             float64   `json:"entries"`
	EntriesURL          string    `json:"entries_url"`
	Expenses            float64   `json:"expenses"`
	ExpensesURL         string    `json:"expenses_url"`
	Group               Group     `json:"group"`
	ID                  float64   `json:"id"`
	Import              Import    `json:"import"`
	InvoicedMinutes     float64   `json:"invoiced_minutes"`
	Invoices            []Invoice `json:"invoices"`
	Minutes             float64   `json:"minutes"`
	Name                string    `json:"name"`
	Participants        []User    `json:"participants"`
	ProjectArchiveURL   string    `json:"project_archive_url"`
	ProjectMergeURL     string    `json:"project_merge_url"`
	ProjectUnarchiveURL string    `json:"project_unarchive_url"`
	RemainingMinutes    float64   `json:"remaining_minutes"`
	UnbillableMinutes   float64   `json:"unbillable_minutes"`
	UpdatedAt           string    `json:"updated_at"`
	URL                 string    `json:"url"`
}

type ProjectService struct {
	client *Client
}

func NewProjectService(client *Client) *ProjectService {
	return &ProjectService{client}
}

type ProjectListCall struct {
	service *ProjectService
	args    map[string]interface{}
}

func (s *ProjectService) List() *ProjectListCall {
	return &ProjectListCall{
		service: s,
		args:    make(map[string]interface{}),
	}
}

func (c *ProjectListCall) Do() (*[]Project, error) {
	data, err := c.service.client.run("GET", "projects", c.args)
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0)

	err = json.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (c *ProjectListCall) Page(page int) *ProjectListCall {
	c.args["page"] = page
	return c
}

func (c *ProjectListCall) Enabled(b bool) *ProjectListCall {
	c.args["enabled"] = b
	return c
}
