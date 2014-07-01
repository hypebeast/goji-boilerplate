var gulp = require('gulp');
var sass = require('gulp-sass');
var csso = require('gulp-csso');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var plumber = require('gulp-plumber');

gulp.task('sass', function() {
  gulp.src('public/stylesheets/styles.scss')
    .pipe(plumber())
    .pipe(sass())
    .pipe(csso())
    .pipe(gulp.dest('public/stylesheets'));
});

gulp.task('compress', function() {
  gulp.src([
    'bower_components/bootstrap-sass-official/javascripts/bootstrap.js',
    'bower_components/jquery/dist/jquery.min.js',
    'public/javascripts/*.js'
  ])
    .pipe(concat('app.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest('public/javascript'));
});

gulp.task('watch', function() {
  gulp.watch('public/stylesheets/*.scss', ['sass']);
  gulp.watch(['public/javascripts/*.js', '!public/app.min.js'], ['compress']);
});

gulp.task('default', ['sass', 'compress']);
