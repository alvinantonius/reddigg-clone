package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func init() {
	Init()
}

func TestNewTopic(t *testing.T) {
	method := "POST"
	requrl := "http://www.example.com/api/v1/topic"

	testCase := []struct {
		InsertFormData func(*http.Request)
		ResStatus      int
	}{
		// good request
		{
			func(r *http.Request) {
				r.Form = url.Values{}
				r.Form.Add("title", "new topic")
			},
			201,
		},

		// bad request no topic defined
		{
			func(r *http.Request) {},
			400,
		},
	}

	for index, tcase := range testCase {
		// build request
		req := httptest.NewRequest(method, requrl, nil)
		w := httptest.NewRecorder()
		p := httprouter.Params{}

		// insert form data to request
		tcase.InsertFormData(req)

		// call request
		NewTopic(w, req, p)

		resp := w.Result()

		if resp.StatusCode != tcase.ResStatus {
			t.Errorf("tcase:%v res got %v | expect %v", index, resp.StatusCode, tcase.ResStatus)
		}
	}
}

func TestListTopic(t *testing.T) {
	method := "GET"

	testCase := []struct {
		URL       string
		ResStatus int
	}{
		// get homepage
		{
			"http://www.example.com/api/v1/topics",
			200,
		},

		// get with correct param
		{
			"http://www.example.com/api/v1/topics?page=2&per_page=10",
			200,
		},

		// get with incorrect param
		{
			"http://www.example.com/api/v1/topics?page=-1&per_page=-346",
			400,
		},

		// get with exceeded perpage limit
		{
			"http://www.example.com/api/v1/topics?page=1&per_page=100000",
			400,
		},
	}

	for index, test := range testCase {
		req := httptest.NewRequest(method, test.URL, nil)
		w := httptest.NewRecorder()
		p := httprouter.Params{}

		ListTopic(w, req, p)

		resp := w.Result()

		if resp.StatusCode != test.ResStatus {
			t.Errorf("case:%v res got %v | expect %v", index, resp.StatusCode, test.ResStatus)
		}
	}
}

func TestUpvote(t *testing.T) {
	method := "POST"

	testCase := []struct {
		URL       string
		Param     httprouter.Params
		ResStatus int
	}{
		// upvote prepopulated data (id=0 is always exist, populated on app start)
		{
			"http://www.example.com/api/v1/topic/0/upvote",
			httprouter.Params{
				httprouter.Param{"topic_id", "0"},
			},
			200,
		},

		// upvote unknown topic
		{
			"http://www.example.com/api/v1/topic/-10/upvote",
			httprouter.Params{
				httprouter.Param{"topic_id", "-10"},
			},
			400,
		},

		// upvote unknown topic
		{
			"http://www.example.com/api/v1/topic/asdasdasd/upvote",
			httprouter.Params{
				httprouter.Param{"topic_id", "asdasdasd"},
			},
			400,
		},
	}

	for index, test := range testCase {
		req := httptest.NewRequest(method, test.URL, nil)
		w := httptest.NewRecorder()
		p := test.Param

		Upvote(w, req, p)

		resp := w.Result()

		if resp.StatusCode != test.ResStatus {
			t.Errorf("case:%v res got %v | expect %v", index, resp.StatusCode, test.ResStatus)
		}
	}
}
