package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	results := TopSearches()
	ga := "<script>(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m);})(window,document,'script','//www.google-analytics.com/analytics.js','ga');ga('create', 'UA-61847115-1', 'auto');</script>"
	fmt.Fprintf(res, "<html><body><strong>Please goto http://www.dickbutt.in/[whatever you want goes here] for some dick butt fun</br>I use Imgur api to grab images, credits to come. </br></strong>%s%s</body></html>\n", results, ga)
}
