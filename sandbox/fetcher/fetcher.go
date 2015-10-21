package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	SRC_HEAP = "/debug/pprof/heap"
	SRC_TRACE = "/debug/pprof/trace?seconds=5"
	SRC_PROFILE = "/debug/pprof/profile"
)

func main() {
	source := "http://devel:8080" + SRC_PROFILE

	var a []string

	strings.Join(a, "+")

	b, err := PostURL(source, strings.Join(a, "+"))
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(b)
	for {
		l, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(l)
	}
}

// PostURL issues a POST to a URL over HTTP.
func PostURL(source, post string) ([]byte, error) {
	resp, err := http.Post(source, "application/octet-stream", strings.NewReader(post))
	if err != nil {
		return nil, fmt.Errorf("http post %s: %v", source, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server response: %s", resp.Status)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
