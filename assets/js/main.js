var app = angular.module('BARD', [ 'ngRoute' ]);

app.config(['$routeProvider, $locationProvider',
	function($routeProvider, $locationProvider) {
		function noteTemplate(lang, topic) {
			return '/notes/'+lang+'/'+topic;
		}
	}
]);