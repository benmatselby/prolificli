package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/benmatselby/prolificli/model"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// API represents what is allowed to be called on the Prolific client.
type API interface {
	CreateStudy(model.CreateStudy) (*model.Study, error)
	GetEligibilityRequirements() (*ListRequirementsResponse, error)
	GetMe() (*Me, error)
	GetStudies(status string) (*ListStudiesResponse, error)
	GetStudy(ID string) (*model.Study, error)
	GetSubmissions(ID string) (*ListSubmissionsResponse, error)
}

// Client is responsible for interacting with the Prolific API.
type Client struct {
	Client  *http.Client
	BaseURL string
	Token   string
	Debug   bool
}

// New will return a new Prolific client.
func New() Client {
	client := Client{
		Client:  http.DefaultClient,
		Token:   viper.GetString("PROLIFIC_TOKEN"),
		BaseURL: strings.TrimRight(viper.GetString("PROLIFIC_URL"), "/"),
		Debug:   viper.GetBool("PROLIFIC_DEBUG"),
	}

	return client
}

// Execute runs an HTTP request.
func (c *Client) Execute(method, url string, body interface{}, response interface{}) (*http.Response, error) {

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, c.BaseURL+url, buf)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "benmatselby/prolificli")
	request.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Token))

	httpResponse, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	responseBody, _ := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))

	if c.Debug {
		fmt.Println(string(responseBody))
	}

	if err := json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(responseBody))).Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding JSON response from %s failed: %v", request.URL, err)
	}

	return httpResponse, nil
}

// CreateStudy is responsible for hitting the Prolific API to create a study.
func (c *Client) CreateStudy(study model.CreateStudy) (*model.Study, error) {
	var response model.Study

	url := "/api/v1/studies/"
	httpResponse, err := c.Execute(http.MethodPost, url, study, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	if httpResponse.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(httpResponse.Body)
		return nil, fmt.Errorf("unable to create study: %v", string(body))
	}

	return &response, nil
}

// GetMe will return your user account details.
func (c *Client) GetMe() (*Me, error) {
	var response Me

	url := "/api/v1/users/me"
	_, err := c.Execute(http.MethodGet, url, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	return &response, nil
}

// GetStudies will return you a list of Study objects.
func (c *Client) GetStudies(status string) (*ListStudiesResponse, error) {
	var response ListStudiesResponse

	if !slices.Contains(model.StudyListStatus, status) {
		return nil, fmt.Errorf("%s is not a valid status: %s", status, strings.Join(model.StudyListStatus, ", "))
	}

	statusFragment := ""
	if status == model.StatusUnpublished {
		statusFragment = "published=0"
	} else {
		statusFragment = fmt.Sprintf("%s=1", status)
	}

	url := fmt.Sprintf("/api/v1/studies/?%s", statusFragment)

	_, err := c.Execute(http.MethodGet, url, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	return &response, nil
}

// GetStudy will return a single study
func (c *Client) GetStudy(ID string) (*model.Study, error) {
	var response model.Study

	url := fmt.Sprintf("/api/v1/studies/%s", ID)
	_, err := c.Execute(http.MethodGet, url, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	return &response, nil
}

// GetSubmissions will return submission data for a given study.
func (c *Client) GetSubmissions(ID string) (*ListSubmissionsResponse, error) {
	var response ListSubmissionsResponse

	url := fmt.Sprintf("/api/v1/studies/%s/submissions/?offset=0&limit=200", ID)
	_, err := c.Execute(http.MethodGet, url, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	return &response, nil
}

// GetEligibilityRequirements will return requirement data.
func (c *Client) GetEligibilityRequirements() (*ListRequirementsResponse, error) {
	var response ListRequirementsResponse

	url := "/api/v1/eligibility-requirements/"
	_, err := c.Execute(http.MethodGet, url, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("unable to fulfil request %s: %s", url, err)
	}

	return &response, nil
}
