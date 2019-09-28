package text

import (
	"fmt"
	"math/rand"
	"time"
)

func GetTweetText() string {
	stations := [13]string{
		"新高岡", "黒部宇奈月温泉", "糸魚川", "上越妙高", "飯山", "上田", "佐久平", "軽井沢", "安中榛名", "高崎", "本庄早稲田", "熊谷", "上野",
	}

	rand.Seed(time.Now().UnixNano())
	passingStation := stations[rand.Intn(len(stations))]
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	tweetText := fmt.Sprintf("ただ今　%s　付近を通過中　\n[%s]", passingStation, currentTime)
	return tweetText
}
