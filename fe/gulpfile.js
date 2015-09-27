var gulp = require('gulp'),
    gutil = require('gulp-util'),
    sass = require('gulp-sass'),
    minifyCss = require('gulp-minify-css'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify');

var paths = {
    sass: ['./scss/*/*.scss', './scss/*scss'],
    uglify: ['./js/*/*.js', './js/*.js'],
    html: ['./templates/*/*.html', './templates/*.html'],
    img: ['./img/*/*.*', './img/*.*']
};

gulp.task('default', ['sass', 'uglify', 'html', 'img']);

gulp.task('sass', function(done){
    gulp.src(paths.sass)
	.pipe(sass({
	    errLogToConsole:true
	}))
	.pipe(minifyCss({
	    keepSecialComments:0
	}))
	.pipe(rename({extname: '.min.css'}))
	.pipe(gulp.dest('../static/css'))
	.on('end', done);
});  // function finish;

gulp.task('uglify', function(){
    gulp.src(paths.uglify)
	.pipe(uglify())
	.pipe(rename({extname: '.min.js'}))
	.pipe(gulp.dest('../static/js'));
});  // function finish;

gulp.task('html', function(){
    gulp.src(paths.html)
	.pipe(gulp.dest('../views'));
});  // function finish;

gulp.task('img', function(){
    gulp.src(paths.img)
	.pipe(gulp.dest('../static/img'));
});  // function finish;

gulp.task('watch', function() {
  gulp.watch(paths.sass, ['sass']);
  gulp.watch(paths.uglify, ['uglify']);
  gulp.watch(paths.img, function(){
    gulp.run('default');
  });
  gulp.watch(paths.html, function(){
    gulp.run('default');
  });
});
