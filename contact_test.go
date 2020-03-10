package active_campaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestContactService_Create(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &CreateContactRequest{
		&Contact{
			Email: "e",
		},
	}

	mux.HandleFunc("/api/3/contacts", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateContactResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		response := &CreateContactResponse{&CreatedContact{
			Email: "e",
		}}

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, response) {
			t.Errorf("Request body = %+v, want %+v", response, input)
		}

		fmt.Fprint(w,
			`
			{
				"contact": {
					"email": "e"
				}
			}`)
	})
	contact, _, err := c.Contacts.Create(input)
	if err != nil {
		t.Errorf("Contacts.Create returned error: %v", err)
	}

	want := &CreateContactResponse{
		&CreatedContact{
			Email: "e",
			Cdate: "",
			Udate: "",
			Orgid: "",
			Links: struct {
				BounceLogs         string `json:"bounceLogs"`
				ContactAutomations string `json:"contactAutomations"`
				ContactData        string `json:"contactData"`
				ContactGoals       string `json:"contactGoals"`
				ContactLists       string `json:"contactLists"`
				ContactLogs        string `json:"contactLogs"`
				ContactTags        string `json:"contactTags"`
				ContactDeals       string `json:"contactDeals"`
				Deals              string `json:"deals"`
				FieldValues        string `json:"fieldValues"`
				GeoIps             string `json:"geoIps"`
				Notes              string `json:"notes"`
				Organization       string `json:"organization"`
				PlusAppend         string `json:"plusAppend"`
				TrackingLogs       string `json:"trackingLogs"`
				ScoreValues        string `json:"scoreValues"`
			}{},
			ID:           "",
			Organization: "",
		}}
	if !reflect.DeepEqual(contact, want) {
		t.Errorf("Contacts.Create returned %+v, want %+v", contact, want)
	}
}

func TestContactService_UpdateListStatusForContact(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &UpdateListStatusForContactRequest{
		List:    "l",
		Contact: "c",
		Status:  "s",
	}

	mux.HandleFunc("/api/3/contactLists", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateContactListStatusResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		response := &UpdateContactListStatusResponse{
			Contacts: nil,
			ContactList: struct {
				Contact               string      `json:"contact"`
				List                  string      `json:"list"`
				Form                  interface{} `json:"form"`
				Seriesid              string      `json:"seriesid"`
				Sdate                 string      `json:"sdate"`
				Status                int         `json:"status"`
				Responder             string      `json:"responder"`
				Sync                  string      `json:"sync"`
				Unsubreason           string      `json:"unsubreason"`
				Campaign              interface{} `json:"campaign"`
				Message               interface{} `json:"message"`
				FirstName             string      `json:"first_name"`
				LastName              string      `json:"last_name"`
				IP4Sub                string      `json:"ip4Sub"`
				Sourceid              string      `json:"sourceid"`
				AutosyncLog           interface{} `json:"autosyncLog"`
				IP4Last               string      `json:"ip4_last"`
				IP4Unsub              string      `json:"ip4Unsub"`
				UnsubscribeAutomation interface{} `json:"unsubscribeAutomation"`
				Links                 struct {
					Automation            string `json:"automation"`
					List                  string `json:"list"`
					Contact               string `json:"contact"`
					Form                  string `json:"form"`
					AutosyncLog           string `json:"autosyncLog"`
					Campaign              string `json:"campaign"`
					UnsubscribeAutomation string `json:"unsubscribeAutomation"`
					Message               string `json:"message"`
				} `json:"links"`
				ID         string      `json:"id"`
				Automation interface{} `json:"automation"`
			}{},
		}

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, response) {
			t.Errorf("Request body = %+v, want %+v", response, input)
		}

		fmt.Fprint(w, `{}`)
	})
	contact, _, err := c.Contacts.UpdateListStatusForContact(input)
	if err != nil {
		t.Errorf("Contacts.UpdateListStatusForContact returned error: %v", err)
	}

	want := &UpdateContactListStatusResponse{
		Contacts: nil,
		ContactList: struct {
			Contact               string      `json:"contact"`
			List                  string      `json:"list"`
			Form                  interface{} `json:"form"`
			Seriesid              string      `json:"seriesid"`
			Sdate                 string      `json:"sdate"`
			Status                int         `json:"status"`
			Responder             string      `json:"responder"`
			Sync                  string      `json:"sync"`
			Unsubreason           string      `json:"unsubreason"`
			Campaign              interface{} `json:"campaign"`
			Message               interface{} `json:"message"`
			FirstName             string      `json:"first_name"`
			LastName              string      `json:"last_name"`
			IP4Sub                string      `json:"ip4Sub"`
			Sourceid              string      `json:"sourceid"`
			AutosyncLog           interface{} `json:"autosyncLog"`
			IP4Last               string      `json:"ip4_last"`
			IP4Unsub              string      `json:"ip4Unsub"`
			UnsubscribeAutomation interface{} `json:"unsubscribeAutomation"`
			Links                 struct {
				Automation            string `json:"automation"`
				List                  string `json:"list"`
				Contact               string `json:"contact"`
				Form                  string `json:"form"`
				AutosyncLog           string `json:"autosyncLog"`
				Campaign              string `json:"campaign"`
				UnsubscribeAutomation string `json:"unsubscribeAutomation"`
				Message               string `json:"message"`
			} `json:"links"`
			ID         string      `json:"id"`
			Automation interface{} `json:"automation"`
		}{},
	}
	if !reflect.DeepEqual(contact, want) {
		t.Errorf("Contacts.UpdateListStatusForContact returned %+v, want %+v", contact, want)
	}
}
