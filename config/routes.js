module.exports = function(server) {

	// Serve index page
	server.get('/', function(req, res, next) {
		res.render('index');
	});

	// Serve notes
	server.get('/notes/:lang/:topic', function(req, res, next) {
		res.render('/notes/'+req.params.lang+'/'+req.params.topic);
	});

	// Redirect to index page for all other requests
	server.get('*', function(req, res, next) {
		res.render('index');
	});

};