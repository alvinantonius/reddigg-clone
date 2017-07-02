package topics

import (
	"reflect"
	"testing"
)

var pkg PkgTopics

func init() {
	pkg = New()
}

func TestCreate(t *testing.T) {
	testCase := []struct {
		Title      string
		WillError  bool
		ExpectedID int64
	}{
		{
			"first topic",
			false,
			0,
		},

		// topic with 256 char long title
		{
			"another topic but is invalid because the title is too looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
			true,
			1,
		},

		// topic with 255 char ling title
		{
			"another topic with 255 char title loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
			false,
			1,
		},

		// third normal topic
		{
			"third topic",
			false,
			2,
		},

		// fourth normal topic
		{
			"fourth topic",
			false,
			3,
		},
	}

	for index, test := range testCase {
		topic, err := pkg.CreateTopic(test.Title)

		// check id sequence
		if topic.ID != test.ExpectedID {
			t.Errorf("invalid TopicID sequence for test case no:%v", index)
		}

		// check error
		if (err == nil && test.WillError) || (!test.WillError && err != nil) {
			t.Errorf("invalid error for test case no:%v got error->%v", index, err)
		}
	}
}

func TestUpvote(t *testing.T) {
	/*
		on this test case, we will give upvotes to previously created topics
		since now we has 4 topic
		after upvotes so if will be sorted like this

		topic_id	upvote
		1			5
		2			3
		3			1
		0			0
	*/

	testCase := []struct {
		ID        int64
		WillError bool
	}{
		// give 5 upvote to topic with ID=1
		{1, false},
		{1, false},
		{1, false},
		{1, false},
		{1, false},

		// give 3 upvote to topic with ID=2
		{2, false},
		{2, false},
		{2, false},

		// give 1 upvote to topic with ID=3
		{3, false},

		// upvote invalid topic_id
		{10, true},
		{-1, true},
	}

	for index, test := range testCase {
		err := pkg.Upvote(test.ID)

		// check error
		if (err == nil && test.WillError) || (!test.WillError && err != nil) {
			t.Errorf("invalid error for test case no:%v got error->%v", index, err)
		}
	}
}

func TestList(t *testing.T) {
	testCase := []struct {
		Take       int64
		Skip       int64
		ExpectData []Topic
		WillError  bool
	}{
		// expect data is sorted because the upvotes from prev test case
		{
			20,
			0,
			[]Topic{
				Topic{
					1,
					"another topic with 255 char title loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
					5,
				},
				Topic{
					2,
					"third topic",
					3,
				},
				Topic{
					3,
					"fourth topic",
					1,
				},
				Topic{
					0,
					"first topic",
					0,
				},
			},
			false,
		},

		// invalid calls
		{
			-234,
			0,
			[]Topic{},
			true,
		},

		// invalid calls again
		{
			34,
			-23,
			[]Topic{},
			true,
		},
	}

	for index, test := range testCase {
		list, err := pkg.List(test.Take, test.Skip)

		// check if data match as expected
		if !reflect.DeepEqual(list, test.ExpectData) {
			t.Errorf("invalid result for test case no:%v \ngot %v\n\nexpecting %v", index, list, test.ExpectData)
		}

		// check error
		if (err == nil && test.WillError) || (!test.WillError && err != nil) {
			t.Errorf("invalid error for test case no:%v got error->%v", index, err)
		}
	}
}
