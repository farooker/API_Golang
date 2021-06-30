package fizzbuzz

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Fizzbuzz : fizzbuzz function
func Fizzbuzz(number int) string { // Fizzbuzz : fizzbuzz function

	if number%15 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func Connected(url string) bool {
    client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}