package active_campaign

import (
	"net/http"
)

// ContactsService handles communication with contact related
// methods of the Active Campaign API.
//
// Active Campaign API docs: https://developers.activecampaign.com/reference#contact
type ContactsService service

type Contact struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type CreateContactRequest struct {
	Contact *Contact `json:"contact"`
}

type CreatedContact struct {
	Email string `json:"email"`
	Cdate string `json:"cdate"`
	Udate string `json:"udate"`
	Orgid string `json:"orgid"`
	Links struct {
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
	} `json:"links"`
	ID           string `json:"id"`
	Organization string `json:"organization"`
}

type CreateContactResponse struct {
	Contact *CreatedContact `json:"contact"`
}

func (s *ContactsService) Create(contact *CreateContactRequest) (*CreateContactResponse, *Response, error) {
	u := "contacts"
	req, err := s.client.NewRequest(http.MethodPost, u, contact)
	if err != nil {
		return nil, nil, err
	}

	c := &CreateContactResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

type ContactList struct {
	List    string `json:"list"`
	Contact string `json:"contact"`
	Status  string `json:"status"`
}

type UpdateListStatusForContactRequest struct {
	ContactList *ContactList `json:"contactList"`
}

type UpdateContactListStatusResponse struct {
	Contacts []struct {
		Cdate               string `json:"cdate"`
		Email               string `json:"email"`
		Phone               string `json:"phone"`
		FirstName           string `json:"firstName"`
		LastName            string `json:"lastName"`
		Orgid               string `json:"orgid"`
		Orgname             string `json:"orgname"`
		SegmentioID         string `json:"segmentio_id"`
		BouncedHard         string `json:"bounced_hard"`
		BouncedSoft         string `json:"bounced_soft"`
		BouncedDate         string `json:"bounced_date"`
		IP                  string `json:"ip"`
		Ua                  string `json:"ua"`
		Hash                string `json:"hash"`
		SocialdataLastcheck string `json:"socialdata_lastcheck"`
		EmailLocal          string `json:"email_local"`
		EmailDomain         string `json:"email_domain"`
		Sentcnt             string `json:"sentcnt"`
		RatingTstamp        string `json:"rating_tstamp"`
		Gravatar            string `json:"gravatar"`
		Deleted             string `json:"deleted"`
		Anonymized          string `json:"anonymized"`
		Adate               string `json:"adate"`
		Udate               string `json:"udate"`
		Edate               string `json:"edate"`
		DeletedAt           string `json:"deleted_at"`
		CreatedUtcTimestamp string `json:"created_utc_timestamp"`
		UpdatedUtcTimestamp string `json:"updated_utc_timestamp"`
		CreatedTimestamp    string `json:"created_timestamp"`
		UpdatedTimestamp    string `json:"updated_timestamp"`
		CreatedBy           string `json:"created_by"`
		UpdatedBy           string `json:"updated_by"`
		Links               struct {
			BounceLogs            string `json:"bounceLogs"`
			ContactAutomations    string `json:"contactAutomations"`
			ContactData           string `json:"contactData"`
			ContactGoals          string `json:"contactGoals"`
			ContactLists          string `json:"contactLists"`
			ContactLogs           string `json:"contactLogs"`
			ContactTags           string `json:"contactTags"`
			ContactDeals          string `json:"contactDeals"`
			Deals                 string `json:"deals"`
			FieldValues           string `json:"fieldValues"`
			GeoIps                string `json:"geoIps"`
			Notes                 string `json:"notes"`
			Organization          string `json:"organization"`
			PlusAppend            string `json:"plusAppend"`
			TrackingLogs          string `json:"trackingLogs"`
			ScoreValues           string `json:"scoreValues"`
			AccountContacts       string `json:"accountContacts"`
			AutomationEntryCounts string `json:"automationEntryCounts"`
		} `json:"links"`
		ID           string      `json:"id"`
		Organization interface{} `json:"organization"`
	} `json:"contacts"`
	ContactList struct {
		Contact     string      `json:"contact"`
		List        string      `json:"list"`
		Form        interface{} `json:"form"`
		Seriesid    string      `json:"seriesid"`
		Sdate       string      `json:"sdate"`
		Udate       interface{} `json:"udate"`
		Status      string      `json:"status"`
		Responder   string      `json:"responder"`
		Sync        string      `json:"sync"`
		Unsubreason string      `json:"unsubreason"`
		Campaign    interface{} `json:"campaign"`
		Message     interface{} `json:"message"`
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		IP4Sub      string      `json:"ip4Sub"`
		// Update list status for a contact does not return a uniform type for Sourceid.
		// If a contact is not a member of the list, it will return a number. Otherwise, a string is returned.
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
	} `json:"contactList"`
}

func (s *ContactsService) UpdateListStatusForContact(contact *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *Response, error) {
	u := "contactLists"
	req, err := s.client.NewRequest(http.MethodPost, u, contact)
	if err != nil {
		return nil, nil, err
	}

	c := &UpdateContactListStatusResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
