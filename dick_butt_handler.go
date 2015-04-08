package main

import (
	"github.com/gorilla/mux"

	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var dickTemplate = `
{{define "page"}}
<html>
	<head>
		<link rel="shortcut icon" href="/assets/favicon.ico" type="image/x-icon">
		<link rel="icon" href="/assets/favicon.ico" type="image/x-icon">
		<style>
			img {
				position: absolute;
			}
			.permlink {
				margin-left: -7px;
				padding: 5px;
				background-color: white;
				border: 2px solid #333;
				border-radius: 5px;
		</style>
	</head>
	<a href='{{.Place}}?top={{.Top}}&bottom={{.Bottom}}&image={{.ImgurSource}}' class="permlink">Permlink</a>
	<body style="background-image: url('{{.ImgurSource}}'); background-size: cover; background-position: center;">
	<a href='{{.Place}}'><img style="top:{{.Top}}%; left:{{.Bottom}}%" src="/assets/dickbutt.png"/></a>
	<script>
		(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
		m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

		ga('create', 'UA-61847115-1', 'auto');
		ga('send', 'pageview');

	</script>
	</body>
</html>
{{end}}
`

type Page struct {
	ImgurSource string
	Top         int
	Bottom      int
	Place       string
}

func DickButtHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	p := Page{
		ImgurSource: ImgurSearcher(vars["place"]),
		Top:         rand.Intn(80),
		Bottom:      rand.Intn(80),
		Place:       vars["place"],
	}

	top, errT := strconv.Atoi(req.URL.Query().Get("top"))
	bottom, errB := strconv.Atoi(req.URL.Query().Get("bottom"))
	image := req.URL.Query().Get("image")
	if errT == nil && errB == nil && len(image) > 0 {
		p.Top = top
		p.Bottom = bottom
		p.ImgurSource = image
	}

	log.Println(p)
	if !stringInSlice(p.Place, strings.Split(os.Getenv("IGNORE_LIST"), ",")) {
		SaveSearch(p.Place)
	}

	templ, err := template.New("page").Parse(dickTemplate)

	if err != nil {
		panic(err)
	}
	err = templ.ExecuteTemplate(res, "page", p)

	if err != nil {
		panic(err)
	}
}
