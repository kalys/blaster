package valve

import (
	"testing"
)

func TestDiff(t *testing.T) {
	info := PlayersInfo{
		NumPlayers: 2,
		Players:    []Player{Player{"kls", 2, 2}, Player{"rc", 2, 2}},
	}

	prev := []Player{Player{"kls", 2, 2}, Player{"boombox", 2, 2}}
	entered, left := info.Diff(prev)

	if len(entered) != 1 {
		t.Fatal("entered size not ok")
	}

	if len(left) != 1 {
		t.Fatal("left size not ok")
	}

	if entered[0].Nickname != "rc" {
		t.Fatal("entered content not ok")
	}

	if left[0].Nickname != "boombox" {
		t.Fatal("left content not ok")
	}
}
