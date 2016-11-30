package blockrio

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

// Use the best user-agent for the user
func GetUserAgent() string {
	var user_agent string

	switch runtime.GOOS {
	case "linux":
		user_agent = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/53.0.2785.143 Chrome/53.0.2785.143 Safari/537.36`
	case "windows":
		user_agent = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36`
	case "mac":
		user_agent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36`
	}

	return user_agent
}

// Fetch the URL and return []byte
func FetchUrlByte(url string, user_agent string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("Error while fetching url\n", err)
	}

	request.Header.Set("User-Agent", user_agent)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error while trying to get response\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error status not OK!\n", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body\n", err)
	}

	return body
}
