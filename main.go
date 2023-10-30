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
