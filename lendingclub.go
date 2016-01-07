package lendingclub

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiVersion  = "v1"
	userAgent   = "Lending Club Go " + version
	version     = "0.1.0"
	contentType = "application/json"
)

const (
	lendingClubAPI = "https://api.lendingclub.com/api/investor/" + apiVersion
)

type Client struct {
	*http.Client
	authToken string
}

type ErrorResponse struct {
	Errors []APIError `json:"errors"`
}

type APIError struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewClient(authToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{Client: client, authToken: authToken}
}

func (c *Client) newRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", c.authToken)
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func debug(bdy io.Reader) {
	bs, err := ioutil.ReadAll(bdy)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(string(bs))
}

func (c *Client) processResponse(res *http.Response, body interface{}) error {
	switch res.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return err
		}
	case http.StatusBadRequest:
		log.Println("BAD REQUEST")
		debug(res.Body)
	case http.StatusForbidden:
		log.Println("FORBIDDEN")
		debug(res.Body)
	case http.StatusUnauthorized:
		return errors.New("unauthorized")
	case http.StatusNotFound:
		log.Println("Not Found")
		debug(res.Body)
	case http.StatusInternalServerError:
		return errors.New(res.Status)
	default:
		return errors.New("unknown status code " + res.Status)
	}

	return nil
}
