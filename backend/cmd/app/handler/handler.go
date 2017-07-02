package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alvinantonius/reddigg-clone/backend/internal/topics"

	"github.com/julienschmidt/httprouter"
)

type (
	// Response is a struct that used to return JSON object for all request
	Response struct {
		Error interface{} `json:"errors,omitempty"`
		Links interface{} `json:"links,omitempty"`
		Data  interface{} `json:"data,omitempty"`
	}

	// Links is for JSON response link that provide link for easy pagination
	Links struct {
		Self string `json:"self"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}
)

/*
NewTopic is handler function for creating new topic

Method = POST
DataType = x-www-form-urlencoded
	params : -topic
*/
func NewTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

/*
ListTopic is handler function for get list of topic
will be used for homepage

Method = GET
Params = -page		(default 1)
		 -per_page 	(default 20)
*/
func ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// Upvote is the handler function for add upvote to certain topic
// Method = POST
// get topic_id from URI
func Upvote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
