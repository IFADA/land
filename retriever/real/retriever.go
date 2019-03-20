package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct { //结构也是一中类型 相当于Java中的Class
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
