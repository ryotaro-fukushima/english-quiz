package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 単語データ
var words = []struct {
	Word         string   `json:"word"`
	Translations []string `json:"translations"`
}{
	{"apple", []string{"りんご", "アップル"}},
	{"run", []string{"走る", "駆ける"}},
	{"dog", []string{"犬"}},
	{"blue", []string{"青", "ブルー"}},
	{"school", []string{"学校"}},
	{"happy", []string{"幸せ", "嬉しい"}},
	{"friend", []string{"友達", "友人"}},
	{"big", []string{"大きい"}},
	{"eat", []string{"食べる"}},
	{"book", []string{"本"}},
}

// ランダムに10問取得
func getQuestions(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]struct {
		Word         string   `json:"word"`
		Translations []string `json:"translations"`
	}, len(words))
	copy(shuffled, words)
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

	// 10問取得（wordsが10未満ならそのまま）
	if len(shuffled) > 10 {
		shuffled = shuffled[:10]
	}

	c.JSON(http.StatusOK, shuffled)
}

// 解答をチェックするAPI
func checkAnswers(c *gin.Context) {
	var userAnswers map[string]string
	if err := c.ShouldBindJSON(&userAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 正解判定
	correctCount := 0
	results := make(map[string]bool)

	for _, word := range words {
		userAnswer, exists := userAnswers[word.Word]
		if !exists {
			continue
		}

		// 正解リストのいずれかに一致すれば正解
		for _, correctTranslation := range word.Translations {
			if userAnswer == correctTranslation {
				correctCount++
				results[word.Word] = true
				break
			}
		}
		if _, found := results[word.Word]; !found {
			results[word.Word] = false
		}
	}

	// スコアを計算
	score := (correctCount * 100) / len(words)
	message := "もっと頑張ろう！"
	if score >= 70 {
		message = "素晴らしい！よく頑張った！"
	}

	// 結果を返す
	c.JSON(http.StatusOK, gin.H{
		"score":   score,
		"results": results,
		"message": message,
	})
}

func main() {
	r := gin.Default()

	// CORSを設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.js のURL
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// エンドポイント
	r.GET("/api/questions", getQuestions)
	r.POST("/api/check", checkAnswers)

	// サーバー起動
	r.Run(":8080") // ポート8080で起動
}
