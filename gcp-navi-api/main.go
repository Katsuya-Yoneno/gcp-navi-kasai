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

<<<<<<< HEAD
func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
    db, err := sql.Open(dbDriver, dsn)
    if err != nil {
        return nil, err
    }
    return db, err
=======
var questions = []question{
	{
		Id:         "1",
		Question:   "GCPでハードディスクを暗号化するための推奨される方法はどれですか？",
		Ans1:       "Key Management Service (KMS)を使用する",
		Ans2:       "Google Cloud Storageの暗号化を使用する",
		Ans3:       "サードパーティのディスク暗号化製品を使用する",
		Ans4:       "Identity and Access Management（IAM）を使用する",
		CorrectAns: "ans1",
	},
	{
		Id:         "2",
		Question:   "Google Cloud SQLでサポートされているプログラミング言語はどれですか？",
		Ans1:       "Java",
		Ans2:       "Python",
		Ans3:       "Go",
		Ans4:       "上記全て",
		CorrectAns: "ans4",
	},
	{
		Id:         "3",
		Question:   "Google Kubernetes Engine (GKE)でアプリケーションをデプロイするために必要なものは何ですか？",
		Ans1:       "Kubernetesの機能を理解する",
		Ans2:       "Dockerコンテナイメージを作成する",
		Ans3:       "kubenetes CLIをインストールする",
		Ans4:       "上記全て",
		CorrectAns: "ans4",
	},
	{
		Id:         "4",
		Question:   "Google Cloud Pub/Subの用途は何ですか？",
		Ans1:       "大規模なデータセットの処理",
		Ans2:       "メッセージングとイベント配信",
		Ans3:       "ユーザー認証とアクセス管理",
		Ans4:       "上記全て",
		CorrectAns: "ans2",
	},
	{
		Id:         "5",
		Question:   "Google Cloud Compute Engineにおいて、Googleによって管理されるサービスを利用するにはどうすればいいですか？",
		Ans1:       "Google Cloud Storageにアクセスする",
		Ans2:       "Google Cloud SQLを使用する",
		Ans3:       "APIキーを生成する",
		Ans4:       "Google Cloud Marketplaceからソフトウェアをインストールする",
		CorrectAns: "ans4",
	},
	{
		Id:         "6",
		Question:   "Google Cloud Compute Engineにおいて、Googleによって管理されるサービスを利用するにはどうすればいいですか？",
		Ans1:       "Google Cloud Storageにアクセスする",
		Ans2:       "Google Cloud SQLを使用する",
		Ans3:       "APIキーを生成する",
		Ans4:       "Google Cloud Marketplaceからソフトウェアをインストールする",
		CorrectAns: "ans4",
	},
>>>>>>> develop/yoneno
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
