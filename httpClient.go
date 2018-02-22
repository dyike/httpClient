package httpClient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Request request struct.
type Requedt struct {
	Method string
	URL    string
	Body   []byte
	Header http.Header
}

func DoRequst(args Request) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(args.Method, args.URL, bytes.NewReader(args.Body))
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.Header = args.Header

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
