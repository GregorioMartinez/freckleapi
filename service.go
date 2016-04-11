package freckleapi

// https://api.letsfreckle.com/v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const basePath = "https://api.letsfreckle.com/"
const version = "v2/"

type Client struct {
	Client *http.Client
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error

	values := make(url.Values)
	for k, v := range params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

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
		req, err = http.NewRequest(method, basePath+version+path+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Header.Get("Link"))

	//@TODO do error handling
	/* if resp.StatusCode != 200 {
		var err ErrorResp
		json.Unmarshal(body, &err)
		return nil, err.Error
	} */

	return body, err
}
