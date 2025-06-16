package mediawiki_test

import (
	"fmt"
	"log"

	"github.com/shiyou0130011/go-mediawiki"
)

func Example_login_to_wiki() {
	// The bots' setting
	//
	// It can be found in wiki's Special:BotPasswords page.
	// More info. can see https://www.mediawiki.org/wiki/Manual:Bot_passwords
	const (
		user     = "my-sample-bot@mediawiki"
		password = "11f5c0050e1a2f05d60be79d671f38e1"
		url      = "https://zh.wikipedia.org/w/api.php"
	)

	mw, err := mediawiki.New(url)
	if err != nil {
		log.Fatal(err)
	}

	// login to wikipedia with bot's setting
	err = mw.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}

	mw.Logout()
}

func ExampleMWApi_Edit() {
	const (
		user     = "my-sample-bot@mediawiki"
		password = "11f5c0050e1a2f05d60be79d671f38e1"
		url      = "https://zh.wikipedia.org/w/api.php"
	)
	mw, err := mediawiki.New(url)
	if err != nil {
		log.Fatal(err)
	}

	err = mw.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}
	defer mw.Logout()

	editConfig := map[string]string{
		"title":   "SOME PAGE",
		"summary": "THIS IS WHAT SHOWS UP IN THE LOG",
		"text":    "THE ENTIRE TEXT OF THE PAGE",
	}
	err = mw.Edit(editConfig)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(`Update "SOME PAGE" successfully`)
}
func ExampleMWApi_PageCategoryList() {
	// When a page's wiki content is following:
	//
	// 	= Lorem Ipsum =
	//
	//	Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//
	//	Vestibulum et eleifend purus. Vivamus ultrices ex sed dolor ultricies tempor.
	//
	//	== Vivamus feugiat ==
	//
	//	Vivamus feugiat nulla tempor, lobortis sem eget. vulputate tortor. Sed tincidunt mi sem, nec fermentum dolor vestibulum ac.
	//
	//	[[category:foo]]
	//	[[category:bar]]
	//	[[category:example]]

	const (
		user     = "my-sample-bot@mediawiki"
		password = "11f5c0050e1a2f05d60be79d671f38e1"
		url      = "https://zh.wikipedia.org/w/api.php"

		pageTitle = "Lorem Ipsum" // The page title to query
	)
	mw, err := mediawiki.New(url)
	if err != nil {
		log.Fatal(err)
	}

	err = mw.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}
	defer mw.Logout()

	categories := mw.PageCategoryList(pageTitle)
	fmt.Println("Categories: ", categories) // will output ["foo", "bar", "example"]

}

func ExampleMWApi_SectionList() {
	// When a page's wiki content is following:
	//
	// 	= Lorem Ipsum =
	//
	//	Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//
	//	Vestibulum et eleifend purus. Vivamus ultrices ex sed dolor ultricies tempor.
	//
	//	== Vivamus feugiat ==
	//
	//	Vivamus feugiat nulla tempor, lobortis sem eget. vulputate tortor. Sed tincidunt mi sem, nec fermentum dolor vestibulum ac.
	//
	//	[[category:foo]]
	//	[[category:bar]]
	//	[[category:example]]

	const (
		user     = "my-sample-bot@mediawiki"
		password = "11f5c0050e1a2f05d60be79d671f38e1"
		url      = "https://zh.wikipedia.org/w/api.php"

		pageTitle = "Lorem Ipsum" // The page title to query
	)
	mw, err := mediawiki.New(url)
	if err != nil {
		log.Fatal(err)
	}

	err = mw.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}
	defer mw.Logout()

	sections, err := mw.SectionList(pageTitle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sections) // will output ["Lorem Ipsum", "Vivamus feugiat"]
}
