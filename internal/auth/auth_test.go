package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", "ApiKey blah")
	res, err := GetAPIKey(req.Header)
	if res == "" || err != nil {
		t.Errorf("getapikey returned: %q \nERROR: %v", res, err)
	}
}

func TestGetAPIKey2(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", "test")
	res, err := GetAPIKey(req.Header)
	if err != nil || res != "" {
		t.Errorf("expected error, but did not get one")
	}
}
