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

// CreatedTag is a struct embedded in the response for creating a tag.
type CreatedTag struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
	TagType     string `json:"tagType"`
	Cdate       string `json:"cdate"`
	Links       *Links `json:"links"`
	ID          string `json:"id"`
}

type Links struct {
	ContactGoalTags string `json:"contactGoalTags"`
}

// CreateTagResponse is the response body returned from creating a tag.
type CreateTagResponse struct {
	Tag *CreatedTag `json:"tag"`
}

// Create a tag.
func (s *TagsService) Create(tag *CreateTagRequest) (*CreateTagResponse, *Response, error) {
	u := "tags"
	req, err := s.client.NewRequest(http.MethodPost, u, tag)
	if err != nil {
		return nil, nil, err
	}

	c := &CreateTagResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
