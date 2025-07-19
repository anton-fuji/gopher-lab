package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ハンドラ
type helloHandler struct{}

func (hh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hello!, %q\n", name)
}

// ハンドラ
type byeHandler struct{}

func (bh *byeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Bye!, %q\n", name)
}

// ハンドラ
type customGreetHandler struct {
	Greet string `json:"greet"`
	Name  string `json:"name"`
}

func (cgh *customGreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&cgh); err != nil {
		fmt.Fprintf(w, "decode err: %v", err)
	}

	fmt.Fprintf(w, "%v!, %q\n", cgh.Greet, cgh.Name)
}

// ハンドラ
func fetchUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fuji")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	type userData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var user userData
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, "decode err: %v", err)
	}
	fmt.Fprintf(w, "user was added, id : %d, name: %q", user.ID, user.Name)
}

// ロガーミドルウェア
func loggingMidlleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		log.Printf("%v => %v", method, path)
		h.ServeHTTP(w, r)
	})
}

// ルーティングを初期化する
func initRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	{
		greetMux := http.NewServeMux()
		greetMux.Handle("GET /greet/hello/{name...}", loggingMidlleware(&helloHandler{}))
		greetMux.Handle("GET /greet/bye/{name...}", loggingMidlleware(&byeHandler{}))
		greetMux.Handle("POST /greet/custom", loggingMidlleware(&customGreetHandler{}))

		mux.Handle("/greet/", greetMux)
	}

	{
		userMux := http.NewServeMux()
		userMux.Handle("GET /user/{id}", loggingMidlleware(http.HandlerFunc(fetchUser)))
		userMux.Handle("POST /user/{id}", loggingMidlleware(http.HandlerFunc(addUser)))

		mux.Handle("/user/", userMux)
	}

	return mux
}

func main() {
	mux := initRoutes()

	fmt.Println("Running Server at 8000")
	http.ListenAndServe(":8000", mux)
}
