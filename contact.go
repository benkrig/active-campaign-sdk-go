package active_campaign

import (
	"net/http"
)

// ContactsService handles communication with contact related
// methods of the Active Campaign API.
//
// Active Campaign API docs: https://developers.activecampaign.com/reference#contact
type ContactsService service

type CreateContactRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type ContactResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Cdate     string `json:"cdate"`
	Udate     string `json:"udate"`
	Orgid     string `json:"orgid"`
	Orgname   string `json:"orgname"`
	Links     struct {
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
	ID           string `json:"id"`
	Organization string `json:"organization"`
}

func (s *ContactsService) Create(contact *CreateContactRequest) (*ContactResponse, *Response, error) {
	u := "contacts"
	req, err := s.client.NewRequest(http.MethodPost, u, contact)
	if err != nil {
		return nil, nil, err
	}

	c := &ContactResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}

type UpdateListStatusForContactRequest struct {
	List    string `json:"list"`
	Contact string `json:"contact"`
	Status  string `json:"status"`
}

type UpdateContactListStatusResponse struct {
	Contacts []struct {
		Cdate               string `json:"cdate"`
		Email               string `json:"email"`
		Phone               string `json:"phone"`
		FirstName           string `json:"firstName"`
		LastName            string `json:"lastName"`
		Orgid               string `json:"orgid"`
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
		DeletedAt           string `json:"deleted_at"`
		CreatedUtcTimestamp string `json:"created_utc_timestamp"`
		UpdatedUtcTimestamp string `json:"updated_utc_timestamp"`
		Links               struct {
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
		ID           string      `json:"id"`
		Organization interface{} `json:"organization"`
	} `json:"contacts"`
	ContactList struct {
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
