package main 

import (
	"net/http"
	"log"
	"io"
)

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080",nil)
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	req, _ := http.NewRequest(r.Method, r.RequestURI, nil)
	req.Header.Set("Range", r.Header.Get("Range"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

}

