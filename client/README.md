# React

This is a front end starter made with React.

## How it works

### React related development

From client directory run `npm run webpackw` and the webpack watcher will
watch all files and build the `bundle.js` file which will be saved in 
`static/js` directory. You can run `npm run webpackp` for minified 
`bundle.js` for production builds.

### Stylesheet development

From the client directory, run `gulp watch` and the watcher will check
all styles and compile them into `style.css` which will be saved in
`static/css` directory.

### Together

Both `webpack` and `gulp` can run concurrently from `npm` to save you
from many open terminals. To do so from the `client` directory run 
`npm run alldev`
