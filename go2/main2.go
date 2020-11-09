package main

import (
	// "fmt"
	// "net/http"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/danilopolani/gocialite.v1"
)

// Define our gocialite instance
var gocial = gocialite.NewDispatcher()

func main() {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/auth/:provider", redirectHandler)
	router.GET("/auth/:provider/callback", callbackHandler)

	router.Run("127.0.0.1:9000")
}

// Show the homepage
func indexHandler(c *gin.Context) {
	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"<a href='/auth/linkedin'><button>Login with LinkedIn</button></a><br>" +
		"<a href='/auth/facebook'><button>Login with Facebook</button></a><br>" +
		"<a href='/auth/google'><button>Login with Google</button></a><br>" +
		"<a href='/auth/bitbucket'><button>Login with Bitbucket</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Amazon</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Slack</button></a><br>" +
		"</body></html>"))
}

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     "7fda4d247b5f83c1cc52",
			"clientSecret": "06f7fe0773dc41753b4f6269e3556b6d0db9916b",
			"redirectURL":  "http://localhost:9000/auth/github/callback",
		},
		"linkedin": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/linkedin/callback",
		},
		"facebook": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/facebook/callback",
		},
		"google": {
			"clientID":     "298623530953-q39343kik2tpbf8qhl933o5ken6sfru3.apps.googleusercontent.com",
			"clientSecret": "KxT1ZfRgyPPbErRnNFf2DVIN",
			"redirectURL":  "http://localhost:9000/auth/google/callback",
		},
		"bitbucket": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/bitbucket/callback",
		},
		"amazon": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/amazon/callback",
		},
		"slack": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/slack/callback",
		},
		"asana": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9000/auth/asana/callback",
		},
	}

	providerScopes := map[string][]string{
		"github":    []string{"public_repo"},
		"linkedin":  []string{},
		"facebook":  []string{},
		"google":    []string{},
		"bitbucket": []string{},
		"amazon":    []string{},
		"slack":     []string{},
		"asana":     []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func callbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")
	fmt.Printf("\n\nprovider:%#v", provider)

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("\n\ntoken:%#v", token)
	fmt.Printf("\n\nuser:%#v\n\n", user)

	// If no errors, show provider name
	c.Writer.Write([]byte("Hi, " + user.FullName))
}
