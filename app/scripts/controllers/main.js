'use strict';

var Main = function($scope, $http) {
  $scope.selectedAnswer = null;
  function getQuestion() {
    $http.get('http://localhost:4747/question').success(function(data){
      console.log('getQuestion data: ', data);
      if(data.success === false) {
        $scope.score = data.score;
        // alert('hi');
      }
      $scope.question = data;
    });
  }
  function postAnswer(params) {
    $http.post('http://localhost:4747/question', params).success(function(data){
      console.log('postAnswer data:', data);
      if(data.success === true) {
        getQuestion();
      }
    });
  }
  $scope.submitAnswer = function() {
    if ($scope.selectedAnswer === null) {
      $scope.error = true;
      return;
    } else {
      $scope.error = false;
    }

    var params = { QuestionId: $scope.question.Id, AnswerId: $scope.selectedAnswer };
    $scope.selectedAnswer = null;
    postAnswer(params);
  };
  // on initial load get 1st question
  getQuestion();
};

module.exports = Main;
