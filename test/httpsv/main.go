package httpsv

import (
	"fmt"
	"net/http"
)

// 1. 基本实现
func Server1() {
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	http.ListenAndServe(":8080", nil)
}

// 2. 实现 ServeHTTP
type HelloHandlerStruct struct {
	content string
}

func (handler *HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}
func Server2() {
	http.Handle("/", &HelloHandlerStruct{content: "Hello World"})
	http.ListenAndServe(":8000", nil)
}


// 3. 自定义ServeMux
type WelcomeHandlerStruct struct {

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func (*WelcomeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}

func Server3 () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HelloHandler)
    mux.Handle("/welcome", &WelcomeHandlerStruct{})
    http.ListenAndServe(":8080", mux)
}
