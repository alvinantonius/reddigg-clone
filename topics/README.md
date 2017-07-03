## How to build and run
on this package i save the data on map of struct with int as the key
then i also create one slice of Topic struct pointers to order topics based on their upvotes

everytime a post got new upvotes, i compare it to topic above it, if the upvotes is higher, i swap the position.
using this approach, the sorting mechanism is become very efficient and easy to scale as well because the data is sorted as new upvotes submitted
