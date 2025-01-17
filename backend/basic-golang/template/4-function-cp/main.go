package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
<<<<<<< HEAD
	//beginanswer
	var total int
	for _, user := range leaderboard.Users {
		total += user.Score
	}
	return total
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
<<<<<<< HEAD
	//beginanswer
	funcMap := template.FuncMap{
		"calculateScore": CalculateScore,
	}
	textTemplate = `{{ range .Users }}{{ .Name }}: {{ .Score }}{{ end }}Total Score: {{ calculateScore . }}`

	tmpl, err := template.New("leaderboard").Funcs(funcMap).Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, leaderboard)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
