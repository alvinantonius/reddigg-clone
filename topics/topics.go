package topics

import (
	"errors"
)

type (
	// PkgTopics is the interfaces abstraction for any package that import it
	// to interact with this package
	PkgTopics interface {
		CreateTopic(string) (Topic, error)
		List(int, int) ([]Topic, error)
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

		sortIndex int64
	}
)

// tSequence is for uniqueID, will auto increment as new topic is created
// start from 0
var tSequence int64

// data is our in memory data storage for storing all topics data in the app
// the index is from tSequence value
var data map[int64]*Topic

// descSorted is our slice of pointer that contains list of sorted topic based on the upvotes
// sorted in descending
var descSorted []*Topic

// init for intialize our empty map for data storage
func init() {
	tSequence = 0
	data = make(map[int64]*Topic)
	descSorted = []*Topic{}

}

// New is for creating our new package obj
func New() PkgTopics {
	pkg := pkgTopics{}

	// populate data with 1 topic
	pkg.CreateTopic("first topic")

	return &pkg
}

// CreateTopic is for creating new topics
// will return created Topic object
func (pkg *pkgTopics) CreateTopic(title string) (Topic, error) {

	// validate max 255 char title length
	if len(title) > 255 {
		return Topic{}, errors.New("title length exceed 255 char limit")
	}

	// get latest position based on upvotes
	sortindex := int64(len(descSorted))

	// create Topic obj
	nTopic := Topic{
		ID:        tSequence,
		Title:     title,
		Upvotes:   0,
		sortIndex: sortindex,
	}

	// add to data collection
	data[tSequence] = &nTopic

	// add to sorted list
	descSorted = append(descSorted, data[tSequence])

	// increment sequence
	tSequence++

	return nTopic, nil
}

// List will get the list of the topic
// sorted by upvotes
func (pkg *pkgTopics) List(take, skip int) ([]Topic, error) {
	// check param
	if take <= 0 || skip < 0 {
		return []Topic{}, errors.New("invalid parameter")
	}

	var result []Topic

	// make sure if param doesn't exceed available data
	if skip < len(data) {
		// if take is exceed data count, overwrite it
		if take+skip > len(data) {
			take = len(data) - skip
		}

		indexStart := skip
		indexEnd := indexStart + take - 1

		for i := indexStart; i <= indexEnd; i++ {
			tData := descSorted[i]
			result = append(result, *tData)
		}
	}

	return result, nil
}

// Upvote is for add 1 upvote to certain topic
func (pkg *pkgTopics) Upvote(tID int64) error {

	if topic, ok := data[tID]; ok {
		// add upvote
		topic.Upvotes++

		// if not on top position
		if topic.sortIndex > 0 {
			// compare current uppvotes to upper sort index
			// if larger, swap position
			if topic.Upvotes > descSorted[topic.sortIndex-1].Upvotes {
				// save current position to temp variable
				temp := topic.sortIndex

				// swap to one topic above it
				topic.sortIndex = temp - 1
				descSorted[temp-1].sortIndex = temp

				// swap position in descSorted slice
				tempTopic := descSorted[temp-1]
				descSorted[temp-1] = topic
				descSorted[temp] = tempTopic
			}
		}

	} else {
		return errors.New("topic id not found")
	}

	return nil
}
