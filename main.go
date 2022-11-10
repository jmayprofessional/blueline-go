package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	intro()
	teamInfo()
	teamRoster()
	fmt.Println("Individual player stats coming up next")
	getPlayerStats("8471817")
}

func teamInfo() {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams/3")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func teamRoster() {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams/3/roster")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func teamStats() {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams/3/?expand=team.stats")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func getPlayerStats(playerId string) {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/people/" + playerId)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

// individual player stats

// player vs player stats

// previous game stats

// + - stats for previous game

// standard input output below for prompting users

func intro() {
	fmt.Println("Greetings, Welcome to Blueline Stat app. What stats would you like to see first?")
	fmt.Println("Type any of the following")
	fmt.Println("* Team Info")
	fmt.Println("* Team Stats")
	fmt.Println("* Team Roster")
	fmt.Println("* Individual Player Stats")

}
