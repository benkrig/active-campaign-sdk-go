package active_campaign

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	// myToken is a non-empty string to use during tests.
	myToken = "my-token"
)

// setup sets up a test HTTP server along with a active_campaign.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client, _ = NewClient(&ClientOpts{
		HttpClient: nil,
		BaseUrl:    server.URL,
		Token:      myToken,
	})

	return client, mux, server.URL, server.Close
}

func TestNewClient_addsTrailingSlashToURLs(t *testing.T) {
	baseURL := "https://custom-url/api/3"
	formattedBaseURL := baseURL + "/"

	c, err := NewClient(
		&ClientOpts{
			nil,
			baseURL,
			"",
		},
	)
	if err != nil {
		t.Fatalf("NewClient returned unexpected error: %v", err)
	}

	if got, want := c.baseURL.String(), formattedBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestNewClient_addsApi3ToURLs(t *testing.T) {
	baseURL := "https://custom-url/"
	formattedBaseURL := baseURL + "api/3/"

	c, err := NewClient(
		&ClientOpts{
			nil,
			baseURL,
			"",
		},
	)
	if err != nil {
		t.Fatalf("NewClient returned unexpected error: %v", err)
	}

	if got, want := c.baseURL.String(), formattedBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	baseURL := "https://custom-url/api/3/"

	c, err := NewClient(
		&ClientOpts{
			nil,
			baseURL,
			"my-token",
		},
	)
	if err != nil {
		t.Fatalf("NewClient returned unexpected error: %v", err)
	}

	inURL, outURL := "foo", baseURL+"foo"
	inBody, outBody := &CreateContactRequest{
		Email:     "test@email.com",
		FirstName: "testf",
		LastName:  "testl",
		Phone:     "",
	}, `{"email":"test@email.com","firstName":"testf","lastName":"testl"}`+"\n"
	req, _ := c.NewRequest(http.MethodPost, inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that token is attached to the request
	if got, want := req.Header.Get("Api-Token"), c.token; got != want {
		t.Errorf("NewRequest() Api-Token is %v, want %v", got, want)
	}
}

func TestCheckResponse(t *testing.T) {
	codes := []int{
		http.StatusOK, http.StatusPartialContent, 299,
	}

	for _, c := range codes {
		r := &http.Response{
			StatusCode: c,
		}
		if err := CheckResponse(r); err != nil {
			t.Errorf("CheckResponse throws an error: %s", err)
		}
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestClient_NewRequest_BadURL(t *testing.T) {
	c, err := NewClient(&ClientOpts{
		HttpClient: nil,
		BaseUrl:    "",
		Token:      "",
	})
	if err != nil {
		t.Errorf("An error occurred. Expected nil. Got %+v.", err)
	}
	_, err = c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

// If a nil body is passed to jira.NewRequest, make sure that nil is also passed to http.NewRequest.
// In most cases, passing an io.Reader that returns no content is fine,
// since there is no difference between an HTTP request body that is an empty string versus one that is not set at all.
// However in certain cases, intermediate systems may treat these differently resulting in subtle errors.
func TestClient_NewRequest_EmptyBody(t *testing.T) {
	c, err := NewClient(&ClientOpts{
		HttpClient: nil,
		BaseUrl:    "",
		Token:      "",
	})
	if err != nil {
		t.Errorf("An error occurred. Expected nil. Got %+v.", err)
	}
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestClient_Do(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/api/3/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		_, _ = fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := c.NewRequest("GET", "/api/3/", nil)
	body := new(foo)
	_, _ = c.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestClient_Do_HTTPResponse(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		_, _ = fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := c.NewRequest("GET", "/", nil)
	res, _ := c.Do(req, nil)
	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error on parsing HTTP Response = %v", err.Error())
	} else if res.StatusCode != 200 {
		t.Errorf("Response code = %v, want %v", res.StatusCode, 200)
	}
}

func TestClient_Do_HTTPError(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/3/contacts", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := c.NewRequest("GET", "/api/3/contacts", nil)
	_, err := c.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

func TestTransport_HeaderContainsApiToken(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/3/contacts", func(w http.ResponseWriter, r *http.Request) {
		// look for Api-Token in the header
		val := r.Header.Get("Api-Token")
		if !strings.Contains(val, "my-token") {
			t.Errorf("request does not contain Api-Token header")
		}
	})
	c.Contacts.Create(nil)
}
