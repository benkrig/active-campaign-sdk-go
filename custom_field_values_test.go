package active_campaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestContactsService_CreateCustomFieldValue(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &CreateCustomFieldValueRequest{
		&FieldValue{
			Contact: "1",
			Field:   "2",
			Value:   "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		},
	}

	mux.HandleFunc("/api/3/fieldValues", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateCustomFieldValueResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if v.FieldValue.Contact != input.FieldValue.Contact {
			t.Errorf("Request body 'contact' = %+v, want %+v", v.FieldValue.Contact, input.FieldValue.Contact)
		}
		if v.FieldValue.Field != input.FieldValue.Field {
			t.Errorf("Request body 'field' = %+v, want %+v", v.FieldValue.Field, input.FieldValue.Field)
		}
		if v.FieldValue.Value != input.FieldValue.Value {
			t.Errorf("Request body 'value' = %+v, want %+v", v.FieldValue.Value, input.FieldValue.Value)
		}

		_, _ = fmt.Fprint(w,
			`
			{
				"contacts": null,
				"fieldValue": {
					"contact": "1",
					"field": "2",
					"value": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
					"cdate": "2020-06-24T15:30:54-05:00",
					"udate": "2020-06-24T15:30:54-05:00",
					"links": null,
					"owner": "#",
					"id": "10"
				}
			}`)
	})
	fieldValue, _, err := c.Contacts.CreateCustomFieldValue(input)
	if err != nil {
		t.Errorf("Contacts.CreateCustomFieldValue returned error: %v", err)
	}

	want := &CreateCustomFieldValueResponse{
		Contacts: nil,
		FieldValue: &FieldValue{
			Contact: "1",
			Field:   "2",
			Value:   "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			Cdate:   "2020-06-24T15:30:54-05:00",
			Udate:   "2020-06-24T15:30:54-05:00",
			Links: struct {
				Owner string `json:"owner,omitempty"`
				Field string `json:"field,omitempty"`
			}{},
			ID:    "10",
			Owner: "#",
		},
	}
	if !reflect.DeepEqual(fieldValue, want) {
		t.Errorf("Contacts.CreateCustomFieldValue returned %+v, want %+v", fieldValue, want)
	}
}
