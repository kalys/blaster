package valve

type Player struct {
	Nickname string
	Kills    uint32
	Time     uint32
}

type PlayersInfo struct {
	NumPlayers byte
	Players    []Player
}

func (playersInfo *PlayersInfo) Diff(players []Player) (entered []Player, left []Player) {
	mb := map[string]bool{}
	for _, x := range players {
		mb[x.Nickname] = true
	}
	entered = []Player{}
	for _, x := range playersInfo.Players {
		if _, ok := mb[x.Nickname]; !ok {
			entered = append(entered, x)
		}
	}

	mb = map[string]bool{}
	for _, x := range playersInfo.Players {
		mb[x.Nickname] = true
	}
	left = []Player{}
	for _, x := range players {
		if _, ok := mb[x.Nickname]; !ok {
			left = append(left, x)
		}
	}

	return
}
