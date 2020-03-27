package active_campaign

import "net/http"

// TagsService handles communication with tag related
// methods of the Active Campaign API.
//
// Active Campaign API docs: https://developers.activecampaign.com/reference#tags
type TagsService service

// Tags are labels that you can apply to contacts to help you organize them.
// The API enables you to add, view, update, and delete tags.
type Tag struct {
	Tag         string `json:"tag,omitempty"`
	TagType     string `json:"tagType,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateTagRequest is the request body used for creating a tag.
type CreateTagRequest struct {
	Tag *Tag `json:"tag"`
}

// Links is embedded in the CreatedTag struct.
type Links struct {
	ContactGoalTags string `json:"contactGoalTags"`
}

// CreatedTag is a struct embedded in the response for creating or retrieving a tag.
type CreatedTag struct {
	Tag             string `json:"tag"`
	Description     string `json:"description"`
	TagType         string `json:"tagType"`
	SubscriberCount string `json:"subscriber_count"`
	Cdate           string `json:"cdate"`
	Links           *Links `json:"links"`
	ID              string `json:"id"`
}

// TagResponse is the response body returned from creating or retrieving a tag.
type TagResponse struct {
	Tag *CreatedTag `json:"tag"`
}

// Meta is embedded in the ListAllResponse struct.
type Meta struct {
	Total string `json:"total"`
}

// ListAllResponse is the response body returned from listing all tags.
type ListAllResponse struct {
	Tags []*CreatedTag `json:"tags"`
	Meta *Meta         `json:"meta"`
}

// Create a tag.
func (s *TagsService) Create(tag *CreateTagRequest) (*TagResponse, *Response, error) {
	u := "tags"
	req, err := s.client.NewRequest(http.MethodPost, u, tag)
	if err != nil {
		return nil, nil, err
	}

	c := &TagResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// Retrieve a tag.
func (s *TagsService) Retrieve(id string) (*TagResponse, *Response, error) {
	u := "tags/" + id
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &TagResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

// Lists all tags.
func (s *TagsService) ListAll() (*ListAllResponse, *Response, error) {
	u := "tags"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &ListAllResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
