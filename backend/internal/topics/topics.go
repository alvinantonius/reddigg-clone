package topics

import ()

type (
	// PkgTopics is the interfaces abstraction for any package that import it
	// to interact with this package
	PkgTopics interface {
		CreateTopic(string) (Topic, error)
		List(int64, int64) ([]Topic, error)
		Upvote(int64) error
	}

	// pkgTopics is the
	pkgTopics struct {
	}

	// Topic it the main data object
	// represent post in our app
	Topic struct {
		ID      int64  `json:"topic_id"`
		Title   string `json:"topic_title"`
		Upvotes int64  `json:"topic_upvotes"`
	}
)

// tSequence is for uniqueID, will auto increment as new topic is created
// start from 0
var tSequence int64

// data is our in memory data storage for storing all topics data in the app
// the index is from tSequence value
var data map[int64]Topic

// descSorted is our slice of pointer that contains list of sorted topic based on the upvotes
// sorted in descending
var descSorted []*Topic

// init for intialize our empty map for data storage
func init() {
	tSequence = 0
	data = make(map[int64]Topic)
	descSorted = []*Topic{}
}

// New is for creating our new package obj
func New() PkgTopics {
	return &pkgTopics{}
}

// CreateTopic is for creating new topics
// will return created Topic object
func (pkg *pkgTopics) CreateTopic(title string) (Topic, error) {
	return Topic{}, nil
}

// List will get the list of the topic
// sorted by upvotes
func (pkg *pkgTopics) List(take, skip int64) ([]Topic, error) {
	return []Topic{}, nil
}

// Upvote is for add 1 upvote to certain topic
func (pkg *pkgTopics) Upvote(tID int64) error {
	return nil
}
