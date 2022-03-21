package jira

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

// ProjectService handles projects for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project
type IssuesService struct {
	client *Client
}

// ProjectList represent a list of Projects
type IssuesList []struct {
	Expand     string `json:"expand,omitempty" structs:"expand,omitempty"`
	StartAt    int32  `json:"startAt,omitempty" structs:"startAt,omitempty"`
	MaxResults int32  `json:"maxResults,omitempty" structs:"maxResults,omitempty"`
	Total      int32  `json:"total,omitempty" structs:"total,omitempty"`
	Issues     Issues `json:"issues,omitempty" structs:"issues,omitempty"`
}

// ProjectCatreqegory represents a single project category
type Issues struct {
	Expand string `json:"expand,omitempty" structs:"expand,omitempty"`
	ID     string `json:"id,omitempty" structs:"id,omitempty"`
	Self   string `json:"self,omitempty" structs:"self,omitempty"`
	Key    string `json:"key,omitempty" structs:"key,omitempty"`
	Fields Fields `json:"fields" structs:"fields"`
}

// Project represents a Jira Project.
type Fields struct {
	Project     Project  `json:"project,omitempty" structs:"project,omitempty"`
	Created     string   `json:"created,omitempty" structs:"created,omitempty"`
	Priority    Priority `json:"priority,omitempty" structs:"priority,omitempty"`
	Status      Status   `json:"status,omitempty" structs:"status,omitempty"`
	Description string   `json:"description,omitempty" structs:"description,omitempty"`
	Summury     string   `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator     Creator  `json:"creator,omitempty" structs:"creator,omitempty"`
}

type Creator struct {
	DisplayName string `json:"displayName,omitempty struct:"displayName,omitempty"`
}

type Priority struct {
	Name string `json:"name,omitempty" struct:"name,omitempty"`
}

type Status struct {
	Name string `json:"name,omitempty" struct:"name,omitempty"`
}

type Project struct {
	ID   string `json:"id,omitempty" structs:"id,omitempty"`
	Key  string `json:"key,omitempty" structrs:"key,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

func (s *IssuesService) GetListWithContext(ctx context.Context) (*IssuesList, *Response, error) {
	return s.ListWithOptionsWithContext(ctx, &GetQueryOptions{})
}

func (s *IssuesService) GetIssues() (*IssuesList, *Response, error) {
	return s.GetListWithContext(context.Background())
}

func (s *IssuesService) ListWithOptionsWithContext(ctx context.Context, options *GetQueryOptions) (*IssuesList, *Response, error) {
	apiEndpoint := "rest/api/3/search?jql=ORDER%20BY%20Created&maxResults=100"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	issuesList := new(IssuesList)
	resp, err := s.client.Do(req, issuesList)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}
	fmt.Println(resp)
	fmt.Println(err)

	return issuesList, resp, nil
}
