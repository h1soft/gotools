package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	rsp, _ := http.Get(url)
	defer rsp.Body.Close()
	content, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.New("read data error")
	}
	return string(content), nil
}
