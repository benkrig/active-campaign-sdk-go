package active_campaign

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

		_, _ = fmt.Fprint(w,
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
		&ContactList{
			List:    "l",
			Contact: "c",
			Status:  "1",
		},
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
				Udate                 interface{} `json:"udate"`
				Status                string      `json:"status"`
				Responder             string      `json:"responder"`
				Sync                  string      `json:"sync"`
				Unsubreason           string      `json:"unsubreason"`
				Campaign              interface{} `json:"campaign"`
				Message               interface{} `json:"message"`
				FirstName             string      `json:"first_name"`
				LastName              string      `json:"last_name"`
				IP4Sub                string      `json:"ip4Sub"`
				Sourceid              interface{} `json:"sourceid,string"`
				AutosyncLog           interface{} `json:"autosyncLog"`
				IP4Last               string      `json:"ip4_last"`
				IP4Unsub              string      `json:"ip4Unsub"`
				CreatedTimestamp      string      `json:"created_timestamp"`
				UpdatedTimestamp      string      `json:"updated_timestamp"`
				CreatedBy             interface{} `json:"created_by"`
				UpdatedBy             interface{} `json:"updated_by"`
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
			}{
				Contact: "c",
				List:    "l",
				Status:  "1",
			},
		}

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, response) {
			t.Errorf("Request body = %+v, want %+v", response, input)
		}

		_, _ = fmt.Fprint(w, `{}`)
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
			Udate                 interface{} `json:"udate"`
			Status                string      `json:"status"`
			Responder             string      `json:"responder"`
			Sync                  string      `json:"sync"`
			Unsubreason           string      `json:"unsubreason"`
			Campaign              interface{} `json:"campaign"`
			Message               interface{} `json:"message"`
			FirstName             string      `json:"first_name"`
			LastName              string      `json:"last_name"`
			IP4Sub                string      `json:"ip4Sub"`
			Sourceid              interface{} `json:"sourceid,string"`
			AutosyncLog           interface{} `json:"autosyncLog"`
			IP4Last               string      `json:"ip4_last"`
			IP4Unsub              string      `json:"ip4Unsub"`
			CreatedTimestamp      string      `json:"created_timestamp"`
			UpdatedTimestamp      string      `json:"updated_timestamp"`
			CreatedBy             interface{} `json:"created_by"`
			UpdatedBy             interface{} `json:"updated_by"`
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

func TestContactService_AddTagToContact(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &AddTagToContactRequest{
		ContactTag: &ContactTag{
			Contact: "1",
			Tag:     "2",
		},
	}

	mux.HandleFunc("/api/3/contactTags", func(w http.ResponseWriter, r *http.Request) {
		v := new(AddTagToContactResponse)
		_ = json.NewDecoder(r.Body).Decode(v)

		response := &AddTagToContactResponse{
			ContactTag: &ContactTag{
				CDate:   "",
				Contact: "1",
				ID:      "",
				Links:   nil,
				Tag:     "2",
			},
		}

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, response) {
			t.Errorf("Request body = %+v, want %+v", response, input)
		}

		_, _ = fmt.Fprint(w,
			`
			{
				"contactTag": {
					"contact": "1",
        			"tag": "2",
        			"cdate": "2020-06-08T19:49:42-05:00",
        			"links": {
            			"tag": "https://your_base_url.api-us1.com/api/3/contactTags/3/tag",
            			"contact": "https://your_base_url.api-us1.com/api/3/contactTags/3/contact"
        			},
        			"id": "3"
				}
			}`)
	})
	contact, _, err := c.Contacts.AddTagToContact(input)
	if err != nil {
		t.Errorf("Contacts.AddTagToContact returned error: %v", err)
	}

	want := &AddTagToContactResponse{
		ContactTag: &ContactTag{
			CDate:   "2020-06-08T19:49:42-05:00",
			Contact: "1",
			ID:      "3",
			Links: &struct {
				Contact string `json:"contact,omitempty"`
				Tag     string `json:"tag,omitempty"`
			}{
				Contact: "https://your_base_url.api-us1.com/api/3/contactTags/3/contact",
				Tag:     "https://your_base_url.api-us1.com/api/3/contactTags/3/tag",
			},
			Tag: "2",
		},
	}
	if !reflect.DeepEqual(contact, want) {
		t.Errorf("Contacts.AddTagToContact returned %+v, want %+v", contact, want)
	}
}

func TestContactService_AddTagToContact_contactNotFound(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	input := &AddTagToContactRequest{
		ContactTag: &ContactTag{
			Contact: "9999999",
			Tag:     "1",
		},
	}

	mux.HandleFunc("/api/3/contactTags", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, `{"message": "Contact not found"}`)
	})
	contactTag, resp, err := c.Contacts.AddTagToContact(input)
	if err == nil {
		t.Error("Contacts.AddTagToContact returned nil err, want not nil")
	}
	if contactTag != nil {
		t.Errorf("Contacts.AddTagToContact returned %+v, want nil", contactTag)
	}
	if resp == nil {
		t.Error("Contacts.AddTagToContact returned nil resp, want not nil")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Contacts.AddTagToContact returned status code %+v, want %+v", resp.StatusCode, http.StatusNotFound)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	want := `{"message": "Contact not found"}`
	if !reflect.DeepEqual(bodyString, want) {
		t.Errorf("Contacts.AddTagToContact resp.Body returned %+v, want %+v", bodyString, want)
	}
}
