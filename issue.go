package jira

type GetQueryOptions struct {
	// Fields is the list of fields to return for the issue. By default, all fields are returned.
	Fields string `url:"fields,omitempty"`
	Expand string `url:"expand,omitempty"`
	// Properties is the list of properties to return for the issue. By default no properties are returned.
	Properties string `url:"properties,omitempty"`
	// FieldsByKeys if true then fields in issues will be referenced by keys instead of ids
	FieldsByKeys  bool   `url:"fieldsByKeys,omitempty"`
	UpdateHistory bool   `url:"updateHistory,omitempty"`
	ProjectKeys   string `url:"projectKeys,omitempty"`
}
