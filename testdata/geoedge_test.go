package testdata

import (
	"testing"
	"geoedge"
	"github.com/subosito/gotenv"
	"os"
)

var (
	NewProjectId string
)

func init() {
	gotenv.Load("../.env")
}

func TestAddProject(t *testing.T) {
	tk := os.Getenv("GEOEDGE_KEY")
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	newProjectParams := map[string]string{
		"name":      "이터널 라이트",
		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAoaFVuSgPrQW/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
		"auto_scan": "0",
		"scan_type": "2",
		"locations": "KR",
		"emulators": "4",
	}

	ar, err := g.AddProject(newProjectParams)
	if err != nil {
		t.Fatal("add projects method not performed correctly")
	} else {
		NewProjectId = ar
		t.Logf("project_id: %v", ar)
	}
}

//func TestAddMultipleProjects(t *testing.T) {
//	tk := os.Getenv("GEOEDGE_KEY")
//	g := geoedge.Geoedge{}
//	err := g.Init(tk)
//	if err != nil {
//		t.Fatal("token not set properly")
//	}
//
//	var projects []map[string]string
//	newProjectParams := map[string]string{
//		"name":      "이터널 라이트 v2",
//		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAoaFVuSgPrQW/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
//		"auto_scan": "0",
//		"scan_type": "2",
//		"locations": "KR",
//		"emulators": "4",
//	}
//	projects = append(projects, newProjectParams)
//	newProjectParams = map[string]string{
//		"name":      "Blades and Rings-ตำนานครูเสด",
//		"tag":       "http://trk.glispa.com/c/AAAAAAAAAAAAAAAAAAAAAlxbVqx1PvYs/CF?placement=1&m.idfa=%geoedge:idfa%&subid1=asdfasdf",
//		"auto_scan": "0",
//		"scan_type": "2",
//		"locations": "TH",
//		"emulators": "7",
//	}
//	projects = append(projects, newProjectParams)
//	ar, err := g.AddMultiProjects(projects)
//	if err != nil {
//		t.Fatal("add multi projects method not performed correctly")
//	} else {
//		t.Logf("projects: %v", ar)
//	}
//}

func TestListProjects(t *testing.T) {
	tk := os.Getenv("GEOEDGE_KEY")
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	ar, err := g.ListProjects()
	if err != nil {
		t.Fatal("list projects method could not be obtained")
	} else {
		t.Logf("list projects: %v", ar)
	}
}

func TestGetProject(t *testing.T) {
	tk := os.Getenv("GEOEDGE_KEY")
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	ar, err := g.GetProject(NewProjectId)
	if err != nil {
		t.Fatal("get project method could not be obtained")
	} else {
		t.Logf("project settings: %v", ar)
	}
}

func TestDeleteProject(t *testing.T) {
	tk := os.Getenv("GEOEDGE_KEY")
	g := geoedge.Geoedge{}
	err := g.Init(tk)
	if err != nil {
		t.Fatal("token not set properly")
	}

	ar, err := g.DeleteProject(NewProjectId)
	if err != nil {
		t.Fatal("delete project method could not be obtained")
	} else {
		t.Logf("delete project: %v", ar)
	}
}
