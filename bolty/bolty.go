package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/boltdb/bolt"
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

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		fmt.Println("Inserting")
		for i := 0; i < 100000000; i++ {
			randText := text[rand.Intn(len(text))]
			randKey := text[rand.Intn(len(text))] + string(rand.Intn(9999999))
			err := b.Put([]byte(randKey), []byte(randText))
			if err != nil {
				return err
			}
			if i%25 == 0 {
				fmt.Printf("\n%v", i)
			}
		}

		return err
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("test"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

	return
}
