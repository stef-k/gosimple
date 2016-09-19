# gosimple

A Go/Beego project used as a starting point for web applications. 
The main goal is to use a JavaScript framework to implement application's views,
with as few as possible server rendered templates from `views` directory.

## Features

* WebSocket integration with
    * basic models for websocket clients (client, room, pool of rooms)
    * JavaScript WebSocket client
* Configuration files for Nginx, SupervisorD and SystemD for *nix systems


## Main differences with vanilla Beego

* changed `controllers/default.go` to `basic.go`
* changed `routers/router.go`, now contains only package's init function setting up routes
* in models directory there are `client.go`, `message.go`, `room.go` and `pool.go` related to WebSocket integration
* `deployment` directory containing configuration files for web server, systemd, etc

## Todo

* Finish database initialization
* Integrate a JS front end framework(?), along with Gulp & NPM files(??)


## Tested on

```DISTRIB_ID=Ubuntu
 DISTRIB_RELEASE=16.04
 DISTRIB_CODENAME=xenial
 DISTRIB_DESCRIPTION="Ubuntu 16.04.1 LTS"
```