package main

import (
	"context"
	"log"
	"net/http"

	"hands-on-microservices/05ContextExample2/repo"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", myHandlerFunc)

	log.Fatal(http.ListenAndServe(":8080", AddSessionData(mux)))
}

type key int

const (
	//Session is
	Session key = iota

	//Authorized is
	Authorized

	//SessionData is
	SessionData
)

//AddSessionData is
func AddSessionData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Header.Get("Session")
		if session != "" {
			sessionData := repo.GetSessionData(session)
			//Logic to validate that session was valid...
			//.....

			//Add data to context
			ctx := context.WithValue(r.Context(), Session, session)
			ctx = context.WithValue(ctx, Authorized, true)
			ctx = context.WithValue(ctx, SessionData, sessionData)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			ctx := context.WithValue(r.Context(), Authorized, false)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	if r.Context().Value(Authorized).(bool) {
		//Write Status Code
		w.WriteHeader(http.StatusOK)

		sessionData := r.Context().Value(SessionData).(map[string]string)
		w.Write([]byte("Hello: " + sessionData["Name"] + " " + sessionData["LastName"] + "\n"))

	} else {

		//Write Unauthorized Status Code
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized222"))
	}
}
