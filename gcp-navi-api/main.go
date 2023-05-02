package main

import (
	"fmt"
	"net/http"
	//"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"database/sql"
    "log"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)

type question struct {
	QuestionId		string  `json:"id"`
	Question    	string `json:"question"`
	Category      	string `json:"category"`
	CorrectAns 		int `json:"correct_ans`
	Ans1     		string `json:"ans1"`
	Ans2     		string `json:"ans2"`
	Ans3     		string `json:"ans3"`
	Ans4     		string `json:"ans4"`
	
}
var questions []question 

func getQuestionsFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT q.question_id, q.question, c.category, a.correct_answer, a.answer1, a.answer2, a.answer3, a.answer4" +
    						" FROM questions as q" +
							" FULL OUTER JOIN answers as a ON q.question_id = a.question_id" +
							" FULL OUTER JOIN category as c ON q.category_id = c.category_id")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 取得したデータをstructにマッピング
	for rows.Next() {
		var q question
		if err := rows.Scan(&q.QuestionId, &q.Question, &q.Category, &q.CorrectAns, &q.Ans1, &q.Ans2, &q.Ans3, &q.Ans4); err != nil {
		   log.Fatal(err)
	   }
		questions = append(questions, q)
	}

	fmt.Println(questions)
}

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
    db, err := sql.Open(dbDriver, dsn)
    if err != nil {
        return nil, err
    }
    return db, err
}

	 func getQuestions(c *gin.Context) {
		c.JSON(http.StatusOK, questions)
	}

func main() {
	// PostgreSQLに接続
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 接続確認
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("PostgreSQLに接続しました！")

	getQuestionsFromDB(db)

	r := gin.Default()
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4200"} //許可するオリジンを指定
    r.Use(cors.New(config))

	r.GET("/questions", getQuestions)

	r.Run()
}
