package mediawiki

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMWApiLogin(t *testing.T) {
	const (
		testAccount  = "bot"
		testPassword = "1qaz2wsx"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/w/api.php" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		bodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		body, _ := url.ParseQuery(string(bodyByte))
		if body.Get("action") == "login" {
			if body.Get("lgname") == testAccount {
				pass := body.Get("lgpassword")
				token := body.Get("lgtoken")
				if token != "" {
					w.Write([]byte(`{
						"login": {
							"result": "Success",
							"lguserid": 30450961,
							"lgusername": "` + testAccount + `"
						}
					}`))
					return
				} else if pass == testPassword {
					w.Write([]byte(`{
						"login": {
							"result": "NeedToken",
							"token": "51dc9e83720ba363971b26e6bd0704ab675921b4+\\\\"
						}
					}`))
					return
				}

			}
		}
		w.Write([]byte(`{}`))
	}))

	defer t.Log("Close test serve")
	defer ts.Close()
	t.Logf("Create test serve %s", ts.URL)

	mw, err := New(ts.URL + "/w/api.php")
	if err != nil {
		t.Fatal(err)
	}
	defer t.Log("Logout from test serve")
	defer mw.Logout()

	t.Log("Login to test serve")

	if err := mw.Login(testAccount, testPassword); err != nil {
		t.Log("Login failed")
		t.Fatal(err)
	}
	t.Log("Login success")
}
