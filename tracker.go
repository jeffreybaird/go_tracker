package main
import (
    "log"
    "net/http"
	"github.com/mrjones/oauth"
)

func main() {
    log.Print("Starting tracker...")

    consumerKey :=  "putkeyhere"
    consumerSecret := "putsecrethere"

	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	c.Debug(true)

	requestToken, url, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("(1) Go to: " + url)
	//fmt.Println("(2) Grant access, you should get back a verification code.")
	//fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode):

	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	response, err := c.Get(
		"http://api.twitter.com/1/statuses/home_timeline.json",
		map[string]string{"count": "1"},
		accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Println("The newest item in your home timeline is: " + string(bits))

    /*
	if *postUpdate {
		status := fmt.Sprintf("Test post via the API using Go (http://golang.org/) at %s", time.Now().String())

		response, err = c.Post(
			"http://api.twitter.com/1/statuses/update.json",
			map[string]string{
				"status": status,
			},
			accessToken)

		if err != nil {
			log.Fatal(err)
		}
	}
    */
}
