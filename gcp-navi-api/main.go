package main

import (
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type question struct {
	Id         string `json:"id"`
	Question   string `json:"question"`
	Ans1       string `json:"ans1"`
	Ans2       string `json:"ans2"`
	Ans3       string `json:"ans3"`
	Ans4       string `json:"ans4"`
	CorrectAns string `json:"correct_ans"`
}

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
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func main() {
	r := gin.Default()
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4200"} //許可するオリジンを指定
    r.Use(cors.New(config))

	r.GET("/questions", getQuestions)

	r.Run()
}
