var gulp = require('gulp');
var sass = require('gulp-sass');

gulp.task('sass', function() {
  return gulp.src("scss/styles.scss")
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest("public/css"));
});

gulp.task('default', gulp.series('sass'));

