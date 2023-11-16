package main

import (
	"encoding/json"
	"fmt"
	"os"

	"edgaru089.ink/go/imcal/internal/sync"
	"edgaru089.ink/go/imcal/internal/util"
	_ "github.com/motemen/go-loghttp/global"
)

type Config struct {
	Username, Password string
	URL                string
}

func main() {
	/*
		http.DefaultTransport = &loghttp.Transport{
			Transport: http.DefaultTransport,
			LogResponse: func(resp *http.Response) {
				log.Printf(" <-- %3d %s (%s)", resp.StatusCode, resp.Request.URL, resp.Header.Get("Content-Type"))
				mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
				if strings.EqualFold(mediaType, "text/html") {
					io.Copy(os.Stderr, resp.Body)
				}
			},
			LogRequest: func(req *http.Request) {
				log.Printf("--> %s %s\n %v", req.Method, req.URL, req.Header)
			},
		}
	*/

	fmt.Println("Hello")

	var conf Config
	json.Unmarshal(util.Unwrap(os.ReadFile("config.json")), &conf)

	c, err := sync.NewClient(conf.Username, conf.Password, conf.URL)
	if err != nil {
		panic(err)
	}

	_, err = c.PullCalendar("/remote.php/dav/calendars/edgar/1/")
	if err != nil {
		panic(err)
	}

}
