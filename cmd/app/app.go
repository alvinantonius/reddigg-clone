package main

import (
	"log"
	"net/http"

	"github.com/alvinantonius/reddigg-clone/backend/cmd/app/handler"

	"github.com/julienschmidt/httprouter"
)

func init() {

}

func main() {
	// create router obj
	router := httprouter.New()

	// route for backend
	router.GET("api/v1/topics", handler.ListTopic)
	router.POST("api/v1/topic", handler.NewTopic)
	router.POST("api/v1/topic/:topic_id/upvote", handler.Upvote)

	// route for frontend
	router.ServeFiles("/app/*filepath", http.Dir("../../../frontend"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
