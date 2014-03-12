angular.module('BARD', [ 'ngRoute', 'BARD.controllers' ])
	.config(['$routeProvider, $locationProvider',
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

			$locationProvider.html5Mode(true);
		}
	]);

angular.module('BARD.controllers', [])
	.controller('HomeCtrl', [ '$scope', '$location',
		function($scope, $location) {

		}
	])
	.controller('CSSCtrl', [ '$scope', '$location',
		function($scope, $location) {

		}
	])