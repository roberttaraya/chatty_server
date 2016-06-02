# chatty_server
Go server for [Chatty](https://github.com/roberttaraya/chatty)

## Setup and Run App Locally
+ install [rethinkdb](https://www.rethinkdb.com/docs/install/)
+ clone repo into the appropriate go/src/github.com directory

Set up database
===============
+ run `rethinkdb` to start rethinkdb
+ go to localhost:8080 to open admin panel
+ click on Data Explorer
+ to create chatty database, type

  `r.dbCreate('chatty')`

  and press Run

+ create user table `r.db('chatty').tableCreate('user')`

+ create channel table `r.db('chatty').tableCreate('channel')`

+ create message table `r.db('chatty').tableCreate('message')`

+ to list all the tables you just created `r.db('chatty').tableList()`

Start Server
==========
+ `go run *.go`
