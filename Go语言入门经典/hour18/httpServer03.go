package main

//响应不同类型的请求
import "net/http"

func helloWorld03(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		w.Write([]byte("Received a GET request\n"))
	case "POST":
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc("/", helloWorld03)
	http.ListenAndServe(":8000", nil)
}
