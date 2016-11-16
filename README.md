# gosimple

A Go/Beego project used as a starting point for web applications.
The main goal is to use a JavaScript framework to implement application's views,
with as few as possible server rendered templates from `views` directory.
If you want to work with Beego's server side static templates,
just remove the `client` directory and modify the `index.html` template.

## Features

* WebSocket integration with
    * basic models for websocket clients (client, room, pool of rooms)
    * WebSocket React component
* Configuration files for Nginx, SupervisorD and SystemD for *nix systems
* Check `localhost:8080/api/v1/websocket` simple API setup returning number of connected websocket clients and rooms
* React ready
* Email sending with or without templates
* REST API ready
* JWT Token generation ready
* Session ready for classic login/logout handling
* User model to start with
* Login rate limiter for failed logins protection
* API rate limiting

### How to start working

1. clone this repository to a **different** directory, for example: `git clone https://github.com/stef-k/gosimple.git myproject`
2. run the `makemine.sh` script (it will remove the `.git` directory and rename all references of Go Simple to myproject)
3. on one terminal tab start the server `bee run`
4. if you want to use the React frontend starter cd into `client` dir and `npm install`
5. start webpack and gulp watchers with `npm run alldev`

Webpack watcher outputs the `bundle.js` file to `static/js/` directory.
Gulp watcher outpus `style.css` `style.min.css` and `style.css.min.map`
to `static/css` directory.

## Main differences with vanilla Beego

* changed `controllers/default.go` to `basic.go`
* changed `routers/router.go`, now contains only package's init function setting up routes
* in models directory there are `client.go`, `message.go`, `room.go` and `pool.go` related to WebSocket integration
* `deployment` directory containing configuration files for web server, systemd, etc

## Todo

* development - production configuration files for the frontend


## Tested on

```DISTRIB_ID=Ubuntu
 DISTRIB_RELEASE=16.04
 DISTRIB_CODENAME=xenial
 DISTRIB_DESCRIPTION="Ubuntu 16.04.1 LTS"
```

```
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.5.1

├── Beego     : 1.7.0
├── GoVersion : go1.7
```
