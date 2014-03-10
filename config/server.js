var express = require('express');
var routes = require('./routes');
var EXPRESS_PORT = 3000;

var server = express();

server.configure(function() {
	server.set('port',EXPRESS_PORT);
	server.set('view engine', 'jade');
	server.use(express.static(__dirname + '../views'));
});

server.listen(EXPRESS_PORT, function() {
	console.log('Server listening on port ' + EXPRESS_PORT);
});