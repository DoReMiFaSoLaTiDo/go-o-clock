package main

import (
	"html/template"
	"os"
	"time"
)

type Match struct {
	Home          string
	Away          string
	EventDateTime time.Time
}

type Webpage struct {
	Title string
	Body  []byte
}

func cacheFile(title string) (*Webpage, error) {
	filename := title + ".txt"
	_, err := os.Open(filename)
	if err != nil {
		theFileName := title + "_raw.txt"
		theFile, err := os.Open(theFileName)
		if err != nil {
			panic(err)
		}
		data := make([]byte, 500)
		count, err := theFile.Read(data)
		theFile.Close()
		if err != nil {
			panic(err)
		}
		if count > 0 {
			filename := title + ".txt"
			err := os.WriteFile(filename, data, 0600)
			if err != nil {
				panic(err)
			}
		}
	}
	newData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &Webpage{Title: title, Body: newData}, nil
}

func loadPage(breadcrumb string) (*Webpage, error) {
	page, err := cacheFile(breadcrumb)
	if err != nil { // file may not exist
		panic(err)
	}
	return page, nil
}

func main() {
	loc, _ := time.LoadLocation("Europe/London")

	weekendMatches := []Match{
		{
			Home:          "Chelsea",
			Away:          "Brentford",
			EventDateTime: time.Date(2023, 10, 28, 12, 30, 45, 100, loc),
		},
		{
			Home:          "Arsenal",
			Away:          "Sheffield United",
			EventDateTime: time.Date(2023, 10, 28, 15, 00, 45, 100, loc),
		},
		{
			Home:          "AFC Bournemouth",
			Away:          "Burnley",
			EventDateTime: time.Date(2023, 10, 28, 15, 00, 45, 100, loc),
		},
		{
			Home:          "Wolverhampton Wanderers",
			Away:          "Newcastle United",
			EventDateTime: time.Date(2023, 10, 28, 17, 30, 45, 100, loc),
		},
	}

	funcMap := template.FuncMap{
		"timeConverter": func(eventTime time.Time) time.Time {
			return eventTime.Local()
		},
	}
	tmplFile := "schedule.html"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, weekendMatches)
	if err != nil {
		panic(err)
	}
}
