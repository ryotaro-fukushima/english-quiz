package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 単語データ
var words = []struct {
	Word         string   `json:"word"`
	Translations []string `json:"translations"`
}{
	// 📘 名詞
	{"student", []string{"生徒"}},
	{"teacher", []string{"先生"}},
	{"school", []string{"学校"}},
	{"class", []string{"授業", "クラス"}},
	{"desk", []string{"机"}},
	{"chair", []string{"いす"}},
	{"pencil", []string{"鉛筆"}},
	{"pen", []string{"ペン"}},
	{"bag", []string{"バッグ", "かばん"}},
	{"book", []string{"本"}},
	{"cat", []string{"猫"}},
	{"dog", []string{"犬"}},
	{"fish", []string{"魚"}},
	{"friend", []string{"友達"}},
	{"name", []string{"名前"}},

	// 🔧 動詞
	{"have", []string{"持っている", "〜がある"}},
	{"like", []string{"好き"}},
	{"play", []string{"遊ぶ", "する"}},
	{"run", []string{"走る"}},
	{"eat", []string{"食べる"}},
	{"go", []string{"行く"}},
	{"come", []string{"来る"}},
	{"speak", []string{"話す"}},
	{"write", []string{"書く"}},
	{"read", []string{"読む"}},
	{"study", []string{"勉強する"}},
	{"see", []string{"見る"}},

	// 🎨 形容詞
	{"big", []string{"大きい"}},
	{"small", []string{"小さい"}},
	{"good", []string{"よい", "すぐれた", "行儀のよい", "じょうずな", "うまい"}},
	{"bad", []string{"悪い"}},
	{"happy", []string{"幸せ", "うれしい", "楽しい"}},
	{"sad", []string{"悲しい"}},
	{"new", []string{"新しい"}},
	{"old", []string{"古い"}},

	// 🚀 副詞
	{"very", []string{"とても", "非常に", "大変"}},
	{"well", []string{"よく", "上手に"}},
	{"fast", []string{"速く"}},
	{"slowly", []string{"ゆっくりと"}},
	{"here", []string{"ここに"}},
	{"there", []string{"そこに"}},
	{"always", []string{"いつも"}},
	{"sometimes", []string{"ときどき"}},

	// ⚙️ 助動詞
	// {"can", []string{"〜することができる", "できる"}},
	// {"will", []string{"〜するつもりである", "〜しようと思う", "〜するでしょう", "だろう"}},

	// 📅 曜日
	{"Sunday", []string{"日曜日"}},
	{"Monday", []string{"月曜日"}},
	{"Tuesday", []string{"火曜日"}},
	{"Wednesday", []string{"水曜日"}},
	{"Thursday", []string{"木曜日"}},
	{"Friday", []string{"金曜日"}},
	{"Saturday", []string{"土曜日"}},

	// 🗓 月
	{"January", []string{"1月"}},
	{"February", []string{"2月"}},
	{"March", []string{"3月"}},
	{"April", []string{"4月"}},
	{"May", []string{"5月"}},
	{"June", []string{"6月"}},
	{"July", []string{"7月"}},
	{"August", []string{"8月"}},
	{"September", []string{"9月"}},
	{"October", []string{"10月"}},
	{"November", []string{"11月"}},
	{"December", []string{"12月"}},

	// ⏰ その他の時間関連語
	{"morning", []string{"朝", "おはよう"}},
	{"afternoon", []string{"午後", "昼"}},
	{"evening", []string{"夕方", "晩"}},
	{"night", []string{"夜"}},
	{"today", []string{"今日"}},
	{"tomorrow", []string{"明日"}},
	{"yesterday", []string{"昨日"}},
	{"now", []string{"今"}},
	{"soon", []string{"すぐに"}},
	{"later", []string{"あとで"}},

	// 🔢 数字
	{"one", []string{"1"}},
	{"two", []string{"2"}},
	{"three", []string{"3"}},
	{"four", []string{"4"}},
	{"five", []string{"5"}},
	{"six", []string{"6"}},
	{"seven", []string{"7"}},
	{"eight", []string{"8"}},
	{"nine", []string{"9"}},
	{"ten", []string{"10"}},
	{"eleven", []string{"11"}},
	{"twelve", []string{"12"}},
	{"thirteen", []string{"13"}},
	{"fourteen", []string{"14"}},
	{"fifteen", []string{"15"}},
	{"sixteen", []string{"16"}},
	{"seventeen", []string{"17"}},
	{"eighteen", []string{"18"}},
	{"nineteen", []string{"19"}},
	{"twenty", []string{"20"}},
	{"thirty", []string{"30"}},
	{"forty", []string{"40"}},
	{"fifty", []string{"50"}},
	{"sixty", []string{"60"}},
	{"seventy", []string{"70"}},
	{"eighty", []string{"80"}},
	{"ninety", []string{"90"}},
	{"one hundred", []string{"100"}},

	// 🌸 季節
	{"spring", []string{"春"}},
	{"summer", []string{"夏"}},
	{"autumn", []string{"秋"}},
	{"fall", []string{"秋"}}, // fall も秋！
	{"winter", []string{"冬"}},

	// 🌦 天気
	{"sunny", []string{"晴れ"}},
	{"cloudy", []string{"くもり"}},
	{"rainy", []string{"雨"}},
	{"snowy", []string{"雪"}},
	{"hot", []string{"暑い"}},
	{"cold", []string{"寒い"}},
	{"warm", []string{"暖かい"}},
	{"cool", []string{"涼しい"}},
	{"wind", []string{"風"}},
	{"windy", []string{"風が強い"}},
	{"weather", []string{"天気"}},
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// ランダムに10問取得
func getQuestions(c *gin.Context) {
	shuffled := make([]struct {
		Word         string   `json:"word"`
		Translations []string `json:"translations"`
	}, len(words))
	copy(shuffled, words)

	// ✅ `rand.NewSource` を使ったランダム生成
	rng.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

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

	fmt.Println("受け取ったデータ:", userAnswers) // ログで確認

	correctCount := 0
	results := make(map[string]bool)

	for _, word := range words {
		// 「英語 → 日本語」の場合（en-to-jp）
		userAnswer, exists := userAnswers[word.Word]
		if exists {
			for _, correctTranslation := range word.Translations {
				if strings.ToLower(strings.TrimSpace(userAnswer)) == strings.ToLower(strings.TrimSpace(correctTranslation)) {
					correctCount++
					results[word.Word] = true
					break
				}
			}
			if _, found := results[word.Word]; !found {
				results[word.Word] = false
			}
		}

		// 「日本語 → 英語」の場合（jp-to-en）
		for _, translation := range word.Translations {
			userAnswer, exists := userAnswers[translation] // ← 日本語訳をキーにする！
			if exists {
				// `strings.ToLower()` と `strings.TrimSpace()` を使って正解と比較
				if strings.ToLower(strings.TrimSpace(userAnswer)) == strings.ToLower(strings.TrimSpace(word.Word)) {
					correctCount++
					results[translation] = true
				} else {
					results[translation] = false
				}
			}
		}
	}

	// スコア計算
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
