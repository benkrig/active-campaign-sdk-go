package active_campaign

import "net/http"

// Custom Field Values are part of the Contacts Service.

// CreateCustomFieldValueRequest is the request body used for updating a custom field value on a contact.
type CreateCustomFieldValueRequest struct {
	FieldValue *FieldValue `json:"fieldValue"`
}

// FieldValue stores a custom field value and the contact information it is attached to.
type FieldValue struct {
	Contact string      `json:"contact"`
	Field   interface{} `json:"field"`
	Value   interface{} `json:"value"`
	Cdate   string      `json:"cdate,omitempty"`
	Udate   string      `json:"udate,omitempty"`
	Links   struct {
		Owner string `json:"owner,omitempty"`
		Field string `json:"field,omitempty"`
	} `json:"links,omitempty"`
	ID    string      `json:"id,omitempty"`
	Owner interface{} `json:"owner,omitempty"`
}

// CreateCustomFieldValueResponse is the response body from updating a custom field value on a contact.
type CreateCustomFieldValueResponse struct {
	Contacts []struct {
		Cdate               string        `json:"cdate"`
		Email               string        `json:"email"`
		Phone               string        `json:"phone"`
		FirstName           string        `json:"firstName"`
		LastName            string        `json:"lastName"`
		Orgid               string        `json:"orgid"`
		Orgname             string        `json:"orgname"`
		SegmentioID         string        `json:"segmentio_id"`
		BouncedHard         string        `json:"bounced_hard"`
		BouncedSoft         string        `json:"bounced_soft"`
		BouncedDate         string        `json:"bounced_date"`
		IP                  string        `json:"ip"`
		Ua                  string        `json:"ua"`
		Hash                string        `json:"hash"`
		SocialdataLastcheck string        `json:"socialdata_lastcheck"`
		EmailLocal          string        `json:"email_local"`
		EmailDomain         string        `json:"email_domain"`
		Sentcnt             string        `json:"sentcnt"`
		RatingTstamp        string        `json:"rating_tstamp"`
		Gravatar            string        `json:"gravatar"`
		Deleted             string        `json:"deleted"`
		Anonymized          string        `json:"anonymized"`
		Adate               string        `json:"adate"`
		Udate               string        `json:"udate"`
		Edate               interface{}   `json:"edate"`
		DeletedAt           string        `json:"deleted_at"`
		CreatedUtcTimestamp string        `json:"created_utc_timestamp"`
		UpdatedUtcTimestamp string        `json:"updated_utc_timestamp"`
		CreatedTimestamp    string        `json:"created_timestamp"`
		UpdatedTimestamp    string        `json:"updated_timestamp"`
		CreatedBy           string        `json:"created_by"`
		UpdatedBy           string        `json:"updated_by"`
		EmailEmpty          bool          `json:"email_empty"`
		AccountContacts     []interface{} `json:"accountContacts"`
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
			AutomationEntryCounts string `json:"automationEntryCounts"`
		} `json:"links"`
		ID           string      `json:"id"`
		Organization interface{} `json:"organization"`
	} `json:"contacts"`
	FieldValue *FieldValue `json:"fieldValue"`
}

// CreateCustomFieldValue adds a custom field to a contact.
func (s *ContactsService) CreateCustomFieldValue(fieldValue *CreateCustomFieldValueRequest) (*CreateCustomFieldValueResponse, *Response, error) {
	u := "fieldValues"
	req, err := s.client.NewRequest(http.MethodPost, u, fieldValue)
	if err != nil {
		return nil, nil, err
	}

	c := &CreateCustomFieldValueResponse{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}
	defer func() { _ = resp.Body.Close() }()

	return c, resp, nil
}
