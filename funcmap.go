package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

func getData(a string) string {
	fp, err := AssetFS.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func html(v string) template.HTML {
	return template.HTML(v)
}

func css(v string) template.CSS {
	return template.CSS(v)
}

func js(v string) template.JS {
	return template.JS(v)
}

func snake(v string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9%]+")
	resURI := url.QueryEscape(v)
	if err != nil {

		log.Fatal(err)
	}
	result := reg.ReplaceAllString(resURI, "")
	result = strings.Replace(result, "%", "_", -1)
	return result
}

func trimQueryParams(v string) string {
	if strings.Contains(v, "?") {
		return strings.Split(v, "?")[0]
	}
	return v
}

func addOne(v int) string {
	return strconv.Itoa(v + 1)
}

func trim(v string) string {
	return strings.TrimSpace(v)
}

func lower(v string) string {
	return strings.ToLower(v)
}

func upper(v string) string {
	return strings.ToUpper(v)
}

func githubLink(v string) string {
	v = strings.ToLower(v)

	v = strings.Replace(v, " ", "-", -1)
	v = strings.Replace(v, ".", "", -1)
	v = strings.Replace(v, "/", "", -1)
	return v
}

func githubLinkIncrementer(v string) string {
	k, ok := githubLinkInc[v]
	if ok {
		githubLinkInc[v]++
		return v + "-" + strconv.Itoa((k + 1))
	}
	githubLinkInc[v] = 0
	return v
}

func merge(v1 int, v2 string) string {
	return strconv.Itoa(v1+1) + ". " + v2
}

func markdown(v string) string {
	return string(blackfriday.Run([]byte(v)))
}

func dateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func color(v string) string {
	switch v {
	case "GET":
		return "info"
	case "POST":
		return "success"
	case "PATCH":
		return "warning"
	case "PUT":
		return "warning"
	case "DELETE":
		return "danger"
	default:
		return "info"

	}
}
