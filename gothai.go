package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// "os"
)

type Question struct {
	Id            int
	Name          string
	Answers       []string
	correctAnswer int
}

type AnswerAttempt struct {
	QuestionId string `json:"QuestionId"`
	AnswerId   string `json:"AnswerId"`
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	// fmt.Println("method", r.Method)
	// if r.Method == "OPTIONS" {
	// 	w.Write("200 OK")
	// }
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
	fmt.Println(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}
	defer r.Body.Close()
	fmt.Println("BODY", string(body))
	var answer AnswerAttempt
	if body == nil {
		fmt.Println("BODY is nil")
		return
	}
	err = json.Unmarshal(body, &answer)
	_ = err
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("%s", answer)

	// respond with next question
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	// currentQuestionNum++
}
