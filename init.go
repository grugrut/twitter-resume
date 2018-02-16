package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	var v url.Values
	u, err := api.GetSelf(v)
	fmt.Printf("%v, %v, %v\n", v, u, err)
}
