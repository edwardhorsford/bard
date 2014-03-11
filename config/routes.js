module.exports = function(server) {

	var path = require('path');

	// Serve index page
	server.get('/', function(req, res, next) {
		res.render('index');
	});

	// Serve notes
	server.get('/notes/:lang/:topic', function(req, res, next) {
		res.render('notes/'+req.params.lang+'/'+req.params.topic+'.html');
	});

	// Serve partials
	server.get('/pages/:type/:page', function(req, res, next) {
		res.render(req.params.type+'/'+req.params.page);
	});

	// Redirect to index page for all other requests
	server.get('*', function(req, res, next) {
		res.render('index');
	});

};