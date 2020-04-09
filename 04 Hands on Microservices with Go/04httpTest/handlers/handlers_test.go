package handlers

import(
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T){
	handler := http.HandlerFunc(MyHandler)

	req := httptest.NewRequest("GET", "/example", nil)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	//you can get the full response

	resp := recorder.Result()

	//now we can check all the elements in the response

	//status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Test failed, expected Status OK (200), got - %d", resp.StatusCode)
	}

	expectedBody := "Request successful"

	if recorder.Body.String() != expectedBody {
		t.Errorf("Test failed, expected - %s, got - %s", expectedBody, recorder.Body.String())
	}

	expectedContentType := "application/text"

	if recorder.HeaderMap.Get("Content-Type") != expectedContentType {
		t.Errorf("Test failed. Expected - %s, got - %s", expectedContentType, recorder.HeaderMap.Get("Content-Type"))
	}
}