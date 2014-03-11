var gulp = require('gulp');
var nodemon = require('gulp-nodemon');
var stylus = require('gulp-stylus');
var jade = require('gulp-jade');
var markdown = require('gulp-markdown');
var jshint = require('gulp-jshint');

var files = {
	js: [ './assets/js/**/*.js' ],
	styl: ['./assets/stylus/**/*.styl' ],
	jade: ['./views/**/*.jade' ],
	notes: [ './notes/**/*.md' ]
};

gulp.task('watch', function() {
	gulp.watch(files.js, [ 'client-js' ]);
	gulp.watch(files.notes, [ 'notes' ]);
	gulp.watch(files.jade, [ 'jade' ]);
	gulp.watch(files.styl, [ 'stylus' ]);
});

gulp.task('client-js', function() {
	gulp.src(files.js)
		.pipe(jshint())
		.pipe(gulp.dest('public/scripts'))
});

gulp.task('stylus', function() {
	gulp.src(files.styl)
		.pipe(stylus())
		.pipe(gulp.dest('public/styles'))
});

gulp.task('notes', function() {
	gulp.src(files.notes)
        .pipe(markdown())
        .pipe(gulp.dest('views/notes'));
});

gulp.task('jade', function() {
	console.log('Put code for converting jade to html files here');
});

gulp.task('server', function() {
	nodemon({script: './config/server.js'})
		.on('restart', [ 'jade' ]);
});

// Default task
gulp.task('default', [ 'server', 'watch' ]);