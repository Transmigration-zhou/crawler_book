package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "crawler_book")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
