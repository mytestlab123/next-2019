package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/game", game)
	http.HandleFunc("/title", title)

	log.Println("Listening on port 8888")
	http.ListenAndServe(":8888", nil)
}

func game(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	req, err := http.NewRequest("GET", "http://game:8888/game", nil)
	if err != nil {
		http.Error(w, "Unable to create request", 500)
		return
	}

	copyHeader(req.Header, r.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "unable to connect to upstream", 500)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read hexagons", 500)
		return
	}

	copyHeader(w.Header(), resp.Header)
	w.Write(buf)
}

func title(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	req, err := http.NewRequest("GET", "http://title:8888/title", nil)
	if err != nil {
		http.Error(w, "Unable to create request", 500)
		return
	}

	copyHeader(req.Header, r.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "unable to connect to upstream", 500)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read hexagons", 500)
		return
	}

	copyHeader(w.Header(), resp.Header)
	w.Write(buf)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
