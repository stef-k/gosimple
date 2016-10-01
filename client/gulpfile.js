/**
 * Gulp build and automation tasks and watchers
 */
// Import the plugins
// the build tool
var gulp = require('gulp');
// compile scss to css
var sass = require('gulp-sass');
// concatenate files
var concat = require('gulp-concat');
// autoprefix css for various vendors based on Can I Use as recommended by Google
var autoprefixer = require('gulp-autoprefixer');
// javascript uglifier
var uglify = require('gulp-uglify');
// css minifier
var cleanCss = require('gulp-clean-css');
// sourcemaps generation
var sourcemaps = require('gulp-sourcemaps');
// rename files
var rename = require('gulp-rename');
// delete directories
var del = require('del');
// path module
var path = require('path');
/*----------------------------------------------------------------------------*/
// SCSS - SASS - CSS Tasks
/*----------------------------------------------------------------------------*/
/**
 * Cleans the output file style.css
 * This is step 1
 */
gulp.task('cleancss', function() {
  return del.sync(['../static/**/style.css',
    '../static/**/style.min.css',
    './src/**/style.css',
    './src/**/style.min.css'
  ], {force: true});
});
/**
 * Compiles SASS to CSS from:
 *    /static/scss
 * The output file(s) will be saved at each scss directory.
 * The compilation will output CSS with prefixes supporting
 * the last 2 versions of browsers.
 * Depends and will call cleancss task
 * This is step 2
 */
gulp.task('sass', ['cleancss'], function() {
  // compile *.scss  to css
  return gulp.src([
    './src/scss/**/*.+(scss|sass)'
  ],{base: 'src'})
  // set output to expanded code style
    .pipe(sass({
      outputStyle: 'expanded'
    }))
    // auto prefix and keep last two browser versions
    .pipe(autoprefixer('last 2 version'))
    // save inplace
    .pipe(gulp.dest(
      function(file){
        return path.join(file.base);
      }
    ))
    .on('end', function() {
      // process.exit();
    });
});
/**
 * Merge all CSS files from /static/css/ to style.css
 * Depends and will call sass task
 * This is step 3
 */
gulp.task('mergecss', ['sass'], function() {
  return gulp.src([
    './src/**/*.css'
  ])
  // concatenate and output file with name..
    .pipe(concat('style.css'))
    // save to destination...
    .pipe(gulp.dest('../static/css'))
    // report end of task and exit
    .on('end', function() {
      // process.exit();
    });
});

/**
 * Minify the produced style.css file
 * Depends and will call mergecss task
 * This is step 4
 */
gulp.task('minifycss', ['mergecss'], function() {
  // select source
  return gulp.src('../static/css/style.css')
  // minify
    .pipe(cleanCss({
      debug: true
    }))
    // rename
    .pipe(rename('style.min.css'))
    // save to...
    .pipe(gulp.dest('../static/css/'))
    // report end of task and exit
    .on('end', function() {
      // process.exit();
    });
});
/**
 * Generate source maps for CSS
 * This is step 5
 */
gulp.task('generateCssMaps', ['minifycss'], function(){
  return gulp.src('../static/css/style.min.css')
    .pipe(sourcemaps.init())
    .pipe(sourcemaps.write('.'))
    .pipe(gulp.dest('../static/css'));
});
/**
 * Cleans temp css files
 * This is step 6 (final step)
 */
gulp.task('cleantmps', ['generateCssMaps'], function() {
  return del.sync(['./src/**/style.css','./src/**/main.css'], {force: true});
});
/*----------------------------------------------------------------------------*/
// Watchers
/*----------------------------------------------------------------------------*/
/**
 * A watcher task, watching all directories under /static/scss/
 * and calls the mergecss task.
 */
gulp.task('watch', function() {
  // SCSS - SASS - CSS
  gulp.watch([
    './src/scss/**/*.+(scss|sass)'
  ], function() {
    gulp.start('cleantmps');
  });
});
