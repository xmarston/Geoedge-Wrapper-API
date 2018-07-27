package geoedge

import (
	"errors"
	"net/http"
	"encoding/json"
	"bytes"
	"strings"
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

type Project struct {
	Id           string
	Name         string
	AutoScan     int
	CreationTime int64
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

func (g *Geoedge) AddProject(project map[string]string) (string, error) {
	var (
		projectId string
		err       error
		projects  = make([]map[string]string, 0)
	)

	projects = append(projects, project)
	projectIds, err := g.AddMultiProjects(projects)
	if len(projectIds) != 1 {
		return projectId, err
	}

	projectId = projectIds[0]

	return projectId, err
}

func (g *Geoedge) AddMultiProjects(projects []map[string]string) ([]string, error) {
	var (
		projectIds []string
		err        error
	)

	ar, err := g.Get("projects/bulk", "POST", projects, true)
	r := g.ParseResponse("add multi project", ar)

	switch v := r.(type) {
	case map[string]string:
		if x, found := v["message"]; found {
			return projectIds, errors.New(x)
		}
	case []interface{}:
		for _, k := range r.([]interface{}) {
			dict := k.(map[string]interface{})
			if val, ok := dict["project_id"]; ok {
				projectIds = append(projectIds, val.(string))
			}
		}
	}

	return projectIds, err
}

func (g *Geoedge) ListProjects() ([]Project, error) {
	var (
		projects []Project
		err      error
	)

	ar, err := g.Get("projects", "GET", nil, false)
	r := g.ParseResponse("list projects", ar)

	switch v := r.(type) {
	case map[string]string:
		if x, found := v["message"]; found {
			return projects, errors.New(x)
		}
	case []interface{}:
		for _, k := range r.([]interface{}) {
			dict := k.(map[string]interface{})
			projects = append(projects, Project{Id: dict["id"].(string), Name: dict["name"].(string), AutoScan: int(dict["auto_scan"].(float64)), CreationTime: int64(dict["creation_time"].(float64))})
		}
	}

	return projects, err
}

func (g *Geoedge) GetProject(projectId string) (Project, error) {
	var (
		err     error
		project Project
	)

	ar, err := g.Get("projects/"+projectId, "GET", nil, false)
	r := g.ParseResponse("get project", ar)

	switch v := r.(type) {
	case map[string]string:
		if x, found := v["message"]; found {
			return project, errors.New(x)
		}
	case interface{}:
		pr := v.(map[string]interface{})
		project.Id = pr["id"].(string)
		project.Name = pr["name"].(string)
		project.AutoScan = int(pr["auto_scan"].(float64))
		project.CreationTime = int64(pr["creation_time"].(float64))
	}

	return project, err
}

func (g *Geoedge) DeleteProject(projectId string) (bool, error) {
	var (
		success bool
		err     error
	)

	success, err = g.DeleteMultiProjects(projectId)

	return success, err
}

func (g *Geoedge) DeleteMultiProjects(projects interface{}) (bool, error) {
	var (
		success    bool
		err        error
		projectIds string
	)

	switch v := projects.(type) {
	case string:
		projectIds = v
	case []string:
		projectIds = strings.Join(v[:], ",")
	}

	ar, err := g.Get("projects/"+projectIds, "DELETE", nil, false)
	r := g.ParseResponse("delete projects", ar)

	switch v := r.(type) {
	case map[string]string:
		if x, found := v["message"]; found {
			return success, errors.New(x)
		}
	case string:
		return true, err
	}

	return success, err
}

func (g *Geoedge) ParseResponse(m string, ar ApiResponse) interface{} {
	var (
		d interface{} // default
	)

	if x, found := ar.Status["message"]; found {
		if x != "Success" {
			return ar.Status
		}
	}

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
