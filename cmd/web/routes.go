package main

import (
  "net/http"
 "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
  router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static"))
  router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

  // And then create the routes using the appropriate methods, patterns and 
  // handlers.
  router.HandlerFunc(http.MethodGet, "/", app.home)
  router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
  router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
  router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

  // Create the middleware chain as normal.
  standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

  // Wrap the router with the middleware and return it as normal.
  return standard.Then(router)
}
