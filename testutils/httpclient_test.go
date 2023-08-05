package testutils

import (
	"io"
	"net/http"
	"testing"
)

func Test_HttpClient(t *testing.T) {
	urls := []struct {
		ConfPattern  string
		URL          string
		ExpectedFail bool
	}{
		{ConfPattern: "/url/1/test", URL: "/url/1/test?someArg=1&someArg=2#fragment"},
		{ConfPattern: "/url/1/test", URL: "/url/1/tes", ExpectedFail: true},
	}
	for _, url := range urls {
		client := NewTestHttpClient()
		client.AddConfig(TestHTTPClientConf{
			URL: url.ConfPattern,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("OKTest"))
			},
		})
		request, err := http.NewRequest("GET", url.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := client.Do(request)
		if url.ExpectedFail {
			if nil == err {
				t.Fatal("Expected fail")
			}
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusAccepted {
			t.Fatalf("Expected status %d, got %d", http.StatusAccepted, resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "OKTest" {
			t.Fatalf("Expected body %s, got %s", "OKTest", string(body))
		}

	}
}
