package main

import (
	"math/rand"

	"log"

	"gopkg.in/resty.v0"
)

func main() {
	text := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	for i := 0; i < 1000000; i++ {
		body := text[rand.Intn(len(text))]
		doReq(body)
	}

}
func doReq(b string) {
	resp, err := resty.R().SetBody(b).Post("http://localhost:8088/upload/boltdb")
	if err != nil {
		log.Fatal("HTTP Request Error")
	}
	if resp.StatusCode() != 200 {
		log.Println("Request Failed")
		return
	}
	return
}
