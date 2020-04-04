package testing

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestDownload(t *testing.T){
	url := "https://www.goinggo.net/post/index.xml"

	statusCode := 200

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\t test 0 = \t When checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call - %v", failed, err)
			}

			t.Logf("\t %s \t Should be able to make the Get call", succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t %s \t Should receive a %d status code", succeed, statusCode)
			} else {
				t.Errorf("\t %s \t Should receive a %d status code - %d", failed, statusCode, resp.StatusCode)
			}
		}
	}
}

/*

1. $ go test
2. $ go test -v
3. $ go test -run Down
4. $ go test -cover (shows how much code is covered by tests)
5. $ go test -coverprofile c.out (this generates "c.out" file, & it contains profiling information
for the cover report)
6. to view c.out file in browser, execute following command

$ go tool cover -html c.out

*/