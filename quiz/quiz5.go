package quiz

type Gamer struct {
	Name  string
	Games []string
}

// gunakan * untuk pointer karena kita merubah isi memori nya
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}
