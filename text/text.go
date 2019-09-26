package text

import "math/rand"

func GetTweetText() string {
	stations := [13]string{
		"新高岡", "黒部宇奈月温泉", "糸魚川", "上越妙高", "飯山", "上田", "佐久平", "軽井沢", "安中榛名", "高崎", "本庄早稲田", "熊谷", "上野",
	}

	text := "ただ今 " + stations[rand.Intn(12)] + " 付近を通過中"
	return text
}
