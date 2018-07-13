package geoedge

import (
	"errors"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

const (
	URL = "https://api.geoedge.com/rest/analytics/v3/"
)

type Geoedge struct {
	Authorization string
}

type ApiResponse struct {
	Status   map[string]string      `json:"status"`
	Response map[string]interface{} `json:"response"`
}

func (g *Geoedge) Init(token string) error {
	if token == "" {
		return errors.New("auth token could not be empty")
	}
	g.Authorization = token
	return nil
}

func (g *Geoedge) Get(path string, method string, params interface{}, jsonBody bool) (ApiResponse, error) {
	var (
		r   ApiResponse
		err error
		req *http.Request
	)

	if method == "" {
		method = "GET"
	}

	client := &http.Client{}

	switch vl := params.(type) {
	case map[string]string:
		req, err = http.NewRequest(method, URL+path, nil)
		if len(vl) > 0 && !jsonBody {
			q := req.URL.Query()
			for v, k := range vl {
				q.Add(v, k)
			}
			req.URL.RawQuery = q.Encode()
		}
	case []map[string]string:
		if len(vl) > 0 && jsonBody {
			jsonValue, _ := json.Marshal(vl)
			req, err = http.NewRequest(method, URL+path, bytes.NewBuffer(jsonValue))
		} else {
			req, err = http.NewRequest(method, URL+path, nil)
		}
	default:
		req, err = http.NewRequest(method, URL+path, nil)
	}

	req.Header.Set("Authorization", g.Authorization)
	if err != nil {
		return r, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return r, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return r, err
	}

	return r, err
}

func (g *Geoedge) ParseResponse(m string, ar ApiResponse) interface{} {
	var (
		d interface{} // default
	)

	switch m {
	case "new project":
		if x, found := ar.Response["project_id"]; found {
			return x
		}
	case "add multi project":
		if x, found := ar.Response["projects"]; found {
			return x
		}
	case "get project":
		if x, found := ar.Response["project"]; found {
			return x
		}
	case "list projects":
		if x, found := ar.Response["projects"]; found {
			return x
		}
	case "delete projects":
		if x, found := ar.Status["code"]; found {
			return x
		}
	}

	return d
}

func (g *Geoedge) AddMultiProjects(projects []map[string]string) ([]string, error) {
	var (
		projectIds []string
		err        error
	)
	ar, err := g.Get("projects/bulk", "POST", projects, true)
	if err != nil {
		return projectIds, err
	}

	r := g.ParseResponse("add multi project", ar)
	fmt.Println(r)
	for _, k := range r.([]interface{}) {
		dict := k.(map[string]interface{})
		if val, ok := dict["project_id"]; ok {
			projectIds = append(projectIds, val.(string))
		}
	}

	return projectIds, err
}
