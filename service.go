package freckleapi

// https://api.letsfreckle.com/v2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tomnomnom/linkheader"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ErrNoMorePages = errors.New("Last of 'em")

const basePath = "https://api.letsfreckle.com/"
const version = "v2/"

type Client struct {
	Client *http.Client
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request

	if method == "POST" {
		j, err := json.Marshal(params)
		if err != nil {
			panic(err)
		}

		r := bytes.NewBuffer(j)
		req, err = http.NewRequest("POST", basePath+version+path, r)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

	} else {
		values := make(url.Values)
		for k, v := range params {
			values.Set(k, fmt.Sprintf("%v", v))
		}
		req, err = http.NewRequest(method, basePath+version+path+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var err Errors
		json.Unmarshal(body, &err)
		return nil, err
	}

	header := resp.Header.Get("Link")
	links := linkheader.Parse(header)

	l := links.FilterByRel("next")
	if len(l) == 0 {
		return body, ErrNoMorePages
	}

	return body, err
}
