package main

import "encoding/json"

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	//beginanswer
	Name  string `json:"name"`
	Email string `json:"-"`
	Rank  int    `json:"rank"`
	//endanswer
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	//beginanswer
	var leaderboard Leaderboard
	err := json.Unmarshal(jsonData, &leaderboard)
	if err != nil {
		return Leaderboard{}, err
	}

	return leaderboard, nil
	//endanswer
}
