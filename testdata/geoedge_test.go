package testdata

import (
	"testing"
	"geoedge"
)

var (
	NewProjectId string
)

//func TestAddProject(t *testing.T) {
//	tk := "388e1fde50eb5883ad7b020fdb42b250"
//	g := geoedge.Geoedge{}
//	err := g.Init(tk)
//	if err != nil {
//		t.Fatal("token not set properly")
//	}
//
//	newProjectParams := map[string]string{
//		"name":      "이터널 라이트",
//		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAoaFVuSgPrQW/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
//		"auto_scan": "0",
//		"scan_type": "2",
//		"locations": "KR",
//		"emulators": "4",
//	}
//	ar, err := g.Get("projects", "POST", newProjectParams, false)
//	r := g.ParseResponse("new project", ar)
//	if r == nil {
//		t.Fatal("add projects method not performed correctly")
//	} else {
//		NewProjectId = r.(string)
//		t.Logf("project_id: %v", r)
//	}
//}

func TestAddMultipleProjects(t *testing.T) {
	tk := "388e1fde50eb5883ad7b020fdb42b250"
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	var projects []map[string]string
	newProjectParams := map[string]string{
		"name":      "이터널 라이트 v2",
		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAoaFVuSgPrQW/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
		"auto_scan": "0",
		"scan_type": "2",
		"locations": "KR",
		"emulators": "4",
	}
	projects = append(projects, newProjectParams)
	newProjectParams = map[string]string{
		"name":      "Blades and Rings-ตำนานครูเสด",
		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAlxbVqx1PvYs/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
		"auto_scan": "0",
		"scan_type": "2",
		"locations": "TH",
		"emulators": "7",
	}
	projects = append(projects, newProjectParams)
	ar, err := g.AddMultiProjects(projects)
	if err != nil {
		t.Fatal("add multi projects method not performed correctly")
	} else {
		t.Logf("projects: %v", ar)
	}
}

func TestListProjects(t *testing.T) {
	tk := "388e1fde50eb5883ad7b020fdb42b250"
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	ar, err := g.Get("projects", "GET", nil, false)
	r := g.ParseResponse("list projects", ar)
	if err != nil {
		t.Fatal("list projects method could not be obtained")
	} else {
		t.Logf("list projects: %v", r)
	}
}

//func TestGetProject(t *testing.T) {
//	tk := "388e1fde50eb5883ad7b020fdb42b250"
//	g := geoedge.Geoedge{}
//	err := g.Init(tk)
//	if err != nil {
//		t.Fatal("token not set properly")
//	}
//
//	ar, err := g.Get("projects/"+NewProjectId, "GET", nil, false)
//	r := g.ParseResponse("get project", ar)
//	if err != nil {
//		t.Fatal("get project method could not be obtained")
//	} else {
//		t.Logf("project settings: %v", r)
//	}
//}
//
//func TestDeleteProject(t *testing.T) {
//	tk := "388e1fde50eb5883ad7b020fdb42b250"
//	g := geoedge.Geoedge{}
//	err := g.Init(tk)
//	if err != nil {
//		t.Fatal("token not set properly")
//	}
//
//	ar, err := g.Get("projects/"+NewProjectId, "DELETE", nil, false)
//	r := g.ParseResponse("delete projects", ar)
//	if err != nil {
//		t.Fatal("delete project method could not be obtained")
//	} else {
//		t.Logf("delete project: %v", r)
//	}
//}
