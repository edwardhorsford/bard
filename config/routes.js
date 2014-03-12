module.exports = function(server) {

	var path = require('path');
	var title = { title: 'BARD' };

	// Serve notes
	server.get('/notes/:lang/:topic', function(req, res, next) {
		res.render('notes/'+req.params.lang+'/'+req.params.topic);
	});

	// Serve partials
	server.get('/pages/:type/:page', function(req, res, next) {
		res.render(req.params.type+'/'+req.params.page);
	});

	// Redirect to index page for all other requests
	server.get('*', function(req, res, next) {
		res.render('index', title);
	});

};