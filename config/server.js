var express = require('express');
var path = require('path');
var engines = require('consolidate');
var routes = require('./routes');
var EXPRESS_PORT = 3000;

var server = express();

server.configure(function() {
	server.set('port',EXPRESS_PORT)
		.set('view engine', 'jade')
		.set('views', path.join(__dirname, '../views'))
		.engine('jade', engines.jade)
		// .engine('.html', engines.html)
		.use(express.static(path.join(__dirname,'../public')))
		.use(server.router);
});

routes(server);

server.listen(EXPRESS_PORT, function() {
	console.log('Server listening on port ' + EXPRESS_PORT);
});