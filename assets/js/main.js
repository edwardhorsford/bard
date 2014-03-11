var app = angular.module('BARD', [ 'ngRoute' ]);

app.config(['$routeProvider, $locationProvider',
	function($routeProvider, $locationProvider) {
		
		function getPartial(type, page) {
			return '/pages/'+type+'/'+page;
		}
		function getNotes(lang, topic) {
			return '/notes/'+lang+'/'+topic;
		}

		$routeProvider
			.when('/', {
				templateUrl: getPartial('static','home'),
				controller: HomeCtrl
			})
			.when('/css/css', {
				templateUrl: getNotes('css', 'css'),
				controller: CSSCtrl
			})
	}
]);

app.controller('HomeCtrl', [ '$scope', '$location',
	function($scope, $location) {

	}
]);

app.controller('CSSCtrl', [ '$scope', '$location',
	function($scope, $location) {

	}
]);