package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	pkgtopics "github.com/alvinantonius/reddigg-clone/topics"

	"github.com/julienschmidt/httprouter"
)

var topics pkgtopics.PkgTopics

type (
	// Response is a struct that used to return JSON object for all request
	Response struct {
		Error interface{} `json:"errors,omitempty"`
		Links Links       `json:"links,omitempty"`
		Data  interface{} `json:"data,omitempty"`
	}

	// Links is for JSON response link that provide link for easy pagination
	Links struct {
		Self string `json:"self"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}
)

// Init is for intialize this handler package
func Init() {
	topics = pkgtopics.New()
}

/*
NewTopic is handler function for creating new topic

Method = POST
DataType = x-www-form-urlencoded
	params : -title
*/
func NewTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	title := r.FormValue("title")

	// if title not defined
	if len(title) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	topic, err := topics.CreateTopic(title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// prepare result header
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// write result
	res := Response{Data: topic}
	res.Links.Self = r.Host + r.RequestURI
	jsonByte, _ := json.Marshal(res)
	w.Write(jsonByte)

	return
}

/*
ListTopic is handler function for get list of topic
will be used for homepage

Method = GET
Params = -page		(default 1)
		 -per_page 	(default 20)
*/
func ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	perPage, _ := strconv.Atoi(r.FormValue("per_page"))
	page, _ := strconv.Atoi(r.FormValue("page"))

	if page < 0 || perPage > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// set default page value to 1
	if page == 0 {
		page = 1
	}

	// set default perPage value to 20
	if perPage == 0 {
		perPage = 20
	}

	skip := (page - 1) * perPage
	data, err := topics.List(perPage, skip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// prepare result header
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// write result
	res := Response{Data: data}
	res.Links.Self = r.Host + r.RequestURI
	jsonByte, _ := json.Marshal(res)
	w.Write(jsonByte)

	return
}

// Upvote is the handler function for add upvote to certain topic
// Method = POST
// get topic_id from URI
func Upvote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	topicIDString := p.ByName("topic_id")
	topicID, _ := strconv.ParseInt(topicIDString, 10, 64)

	// if topic id is not number
	if res, _ := regexp.MatchString("^[0-9]+$", topicIDString); !res {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := topics.Upvote(topicID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	return

}
