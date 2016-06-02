package main

import (
  "log"
  "net/http"
  r "github.com/dancannon/gorethink"
)

func main() {
  session, err := r.Connect(r.ConnectOpts{
    Address:  "localhost:28015",
    Database: "chatty",
  })

  if err != nil {
    log.Panic(err.Error())
  }

  router := NewRouter(session)

  router.Handle("channel add", addChannel)
  router.Handle("channel subscribe", subscribeChannel)
  router.Handle("channel unsubscribe", unsubscribeChannel)

  router.Handle("message add", addChannelMessage)
  router.Handle("message subscribe", subscribeChannelMessage)
  router.Handle("message unsubscribe", unsubscribeChannelMessage)

  router.Handle("user edit", editUser)
  router.Handle("user subscribe", subscribeUser)
  router.Handle("user unsubscribe", unsubscribeUser)

  http.Handle("/", router)
  http.ListenAndServe(":4000", nil)
}
