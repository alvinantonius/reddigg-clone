var app = angular.module("topicsApp", []); 
app.controller("topicsController", function($scope, $http) {
	
	$scope.topicList = [];
	$scope.currPage = 1;
	$scope.newTitle = "";

	$scope.GetList = function(page) {
		$http.get("/api/v1/topics?page="+page+"&per_page=20").then(function(response){
			$scope.topicList = response.data.data;
			$scope.currPage = page;
		})
	}

	$scope.NewTopic = function() {

		// ignore empty title
		if($scope.newTitle == ""){
			return 
		}

		$http({
		    method: 'POST',
		    url: '/api/v1/topic',
		    data: "title=" + $scope.newTitle,
		    headers: {'Content-Type': 'application/x-www-form-urlencoded'}
		}).then(function(response){
			$scope.newTitle = "";
			$scope.GetList(1);
		});
	}

	$scope.Upvote = function(listIndex, topicID) {
		$http({
		    method: 'POST',
		    url: '/api/v1/topic/' + topicID + '/upvote',
		    data: "",
		    headers: {'Content-Type': 'application/x-www-form-urlencoded'}
		});

		currUpvotes = $scope.topicList[listIndex].topic_upvotes
		$scope.topicList[listIndex].topic_upvotes = currUpvotes+1;
	}

	$scope.GetList(1);

});