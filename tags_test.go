package active_campaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagService_Create(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &CreateTagRequest{
		&Tag{
			Tag:         "My Tag",
			TagType:     "Contact",
			Description: "Description",
		},
	}

	mux.HandleFunc("/api/3/tags", func(w http.ResponseWriter, r *http.Request) {
		v := new(TagResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if v.Tag.Tag != input.Tag.Tag {
			t.Errorf("Request body 'tag' = %+v, want %+v", v.Tag.Tag, input.Tag.Tag)
		}
		if v.Tag.TagType != input.Tag.TagType {
			t.Errorf("Request body 'tagType' = %+v, want %+v", v.Tag.TagType, input.Tag.TagType)
		}
		if v.Tag.Description != input.Tag.Description {
			t.Errorf("Request body 'description' = %+v, want %+v", v.Tag.Description, input.Tag.Description)
		}

		_, _ = fmt.Fprint(w,
			`
			{
				"tag": {
					"tag": "My Tag",
					"tagType": "contact",
					"description": "Description",
					"cdate": "2020-03-27T13:09:10-05:00",
					"links": {
						"contactGoalTags": "https://:account.api-us1.com/api/:version/tags/1/contactGoalTags"
					},
					"id": "1"
				}
			}`)
	})
	tag, _, err := c.Tags.Create(input)
	if err != nil {
		t.Errorf("Tags.Create returned error: %v", err)
	}

	want := &TagResponse{
		&CreatedTag{
			Tag:         "My Tag",
			TagType:     "contact",
			Description: "Description",
			Cdate:       "2020-03-27T13:09:10-05:00",
			Links:       &Links{ContactGoalTags: "https://:account.api-us1.com/api/:version/tags/1/contactGoalTags"},
			ID:          "1",
		}}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Create returned %+v, want %+v", tag, want)
	}
}

func TestTagService_Create_EmptyTag(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	// Create Tag can be called with an empty request body.
	// Note: there is a unique constraint on TagType, so this can only happen once.
	// Subsequent empty calls would return "Duplicate entry '' for key 'typetag'"
	input := &CreateTagRequest{
		&Tag{},
	}

	mux.HandleFunc("/api/3/tags", func(w http.ResponseWriter, r *http.Request) {
		v := new(TagResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if v.Tag.Tag != input.Tag.Tag {
			t.Errorf("Request body 'tag' = %+v, want %+v", v.Tag.Tag, input.Tag.Tag)
		}
		if v.Tag.TagType != input.Tag.TagType {
			t.Errorf("Request body 'tagType' = %+v, want %+v", v.Tag.TagType, input.Tag.TagType)
		}
		if v.Tag.Description != input.Tag.Description {
			t.Errorf("Request body 'description' = %+v, want %+v", v.Tag.Description, input.Tag.Description)
		}

		_, _ = fmt.Fprint(w,
			`
			{
				"tag": {
					"cdate": "2020-03-27T13:09:10-05:00",
					"links": {
						"contactGoalTags": "https://:account.api-us1.com/api/:version/tags/1/contactGoalTags"
					},
					"id": "1"
				}
			}`)
	})
	tag, _, err := c.Tags.Create(input)
	if err != nil {
		t.Errorf("Tags.Create returned error: %v", err)
	}

	want := &TagResponse{
		&CreatedTag{
			Tag:         "",
			Description: "",
			TagType:     "",
			Cdate:       "2020-03-27T13:09:10-05:00",
			Links:       &Links{ContactGoalTags: "https://:account.api-us1.com/api/:version/tags/1/contactGoalTags"},
			ID:          "1",
		}}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Create returned %+v, want %+v", tag, want)
	}
}

func TestTagService_Retrieve(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/3/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/api/3/tags/1")
		_, _ = fmt.Fprint(w,
			`
			{
				"tag": {
					"tag": "My Tag",
					"tagType": "contact",
					"description": "Description",
					"cdate": "2020-03-27T13:09:10-05:00",
					"links": {
						"contactGoalTags": "https://:account.api-us1.com/api/:version/tags/1/contactGoalTags"
					},
					"id": "1"
				}
			}`)
	})
	tag, _, err := c.Tags.Retrieve("1")
	if err != nil {
		t.Errorf("Tags.Retrieve returned error: %v", err)
	}
	if tag == nil {
		t.Errorf("Expected tag. Tag is nil")
	}
	if tag.Tag.ID != "1" {
		t.Errorf("Expected tag.Tag.ID = 1. Got tag.Tag.ID = %s", tag.Tag.ID)
	}
}

func TestTagService_Retrieve_NotFound(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/3/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/api/3/tags/1")
		w.WriteHeader(http.StatusNotFound)
	})

	_, resp, err := c.Tags.Retrieve("1")
	if err == nil {
		t.Errorf("Expected error. Error is nil")
	}
	if resp == nil {
		t.Errorf("Expected response. Response is nil")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d. Got %d", http.StatusNotFound, resp.StatusCode)
	}
}
