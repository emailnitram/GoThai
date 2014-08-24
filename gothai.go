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
	CorrectAnswer int `json:"-"`
}

type AnswerAttempt struct {
	QuestionId int `json:"QuestionId"`
	AnswerId   int `json:"AnswerId"`
}

type Response struct {
	Success bool `json:"success"`
	Score   int  `json:"score"`
}

var currentQuestionNum int = 0
var score int = 0

func main() {

	http.HandleFunc("/question", getQuestion)
	log.Fatal(http.ListenAndServe(":4747", nil))
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
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

	// err = c.Insert(&Question{0, "What is the word for eat?", []string{"วิ่ง", "นอน", "ดื่ม", "กิน"}, 3},
	// 	&Question{1, "What is the word for hungry?", []string{"หิว", "เหนื่อย", "สบาย", "บาท"}, 0},
	// 	&Question{2, "What is the word for happy?", []string{"ไกล", "สบาย", "ใจ", "โกรธ"}, 1})
	// if err != nil {
	// 	panic(err)
	// }

	result := Question{}
	err = c.Find(bson.M{"id": currentQuestionNum}).One(&result)
	if err != nil {
		res := Response{Success: false, Score: score}
		b, err := json.Marshal(res)
		_ = err
		w.Write(b)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}
	defer r.Body.Close()
	var answer AnswerAttempt

	fmt.Println("Body", string(body))
	err = json.Unmarshal(body, &answer)
	_ = err
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Result", result)
	if r.Method == "POST" {
		if answer.QuestionId == result.Id && answer.AnswerId == result.CorrectAnswer {
			score++
		}
		res := Response{Success: true, Score: 0}
		b, err := json.Marshal(res)
		_ = err
		w.Write(b)
		currentQuestionNum++
		return
	}
	fmt.Println("Score: ", score)

	// respond with next question
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}
