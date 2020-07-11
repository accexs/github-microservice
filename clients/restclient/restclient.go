package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func GetMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enabledMocks = true
}

func StopMockups() {
	enabledMocks = false
}

func AddMockup(mock Mock) {
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enabledMocks {
		mock := mocks[GetMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found")
		}
		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}

func Get() {}
