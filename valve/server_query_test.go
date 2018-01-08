package valve

import (
	"fmt"
	"testing"
	"time"
)

func TestQueryPlayers(t *testing.T) {
	querier, err := NewServerQuerier("46.42.18.18:27018", time.Second*10)

	players, err := querier.QueryPlayers()
	fmt.Println(players)
	if err != nil {
		t.Fatalf("Something went wrong", err)
	}
}
