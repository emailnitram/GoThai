package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
)

type Question struct {
	Id            int
	Name          string
	Answers       []string
	correctAnswer int
}

var currentQuestionNum int = 0

func main() {

	// err = c.Insert(&Question{0, "What is the capital of Thailand?", []string{"Phuket", "Bangkok", "Chiang Mai", "Krabi"}, 1},
	// 	&Question{1, "What is currency used in Thailand?", []string{"Dollar", "Dong", "Baht", "Yen"}, 2})
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/question", getQuestion)
	log.Fatal(http.ListenAndServe(":4747", nil))
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout("localhost:27017", maxWait)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("thaiQuiz").C("questions")
	result := Question{}
	err = c.Find(bson.M{"id": currentQuestionNum}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Write(b)
	// currentQuestionNum++
}
