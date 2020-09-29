package main

import (
	"fmt"
	"github.com/michael-j-taylor/go-a2s"
	"time"
)


func main() {

	addr := "nobps.rustoria.us:28015"

	c := make(chan *a2s.PlayerInfo)

	go getPlayers(addr, 10, c)

	oldPlayerInfo := <-c

	for {
		newPlayerInfo := <-c

		updatePlayers(oldPlayerInfo, newPlayerInfo)

		oldPlayerInfo = newPlayerInfo
	}
}


func getPlayers(addr string, delay int, c chan<- *a2s.PlayerInfo) {

	//establish client
	client, err := a2s.NewClient(addr)

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("client established")
	}

	//Loop forever, sending updated player info back via channel
	//every (delay) seconds. No need to close channel, as when main()
	//exits, this goroutine should be killed from what I understand
	for {
		playerInfo, err := client.QueryPlayer()

		if err != nil {
			fmt.Println("could not retrieve playerInfo\n", err)
		}

		c <- playerInfo
		time.Sleep(time.Duration(delay) * time.Second)
	}
}


func updatePlayers(oldPlayerInfo, newPlayerInfo *a2s.PlayerInfo) {

	left, joined := []string{}, []string{}
	left, joined = comparePlayers(oldPlayerInfo, newPlayerInfo)

	for _, v := range left {
		fmt.Println(v, " has left the server")
	}

	for _, v := range joined {
		fmt.Println(v, " has joined the server")
	}

}

func comparePlayers(playerInfo1, playerInfo2 *a2s.PlayerInfo) ([]string, []string) {

	map1 := make(map[string]bool)
	map2 := make(map[string]bool)

	//generate maps from array of Players
	for _, v := range playerInfo1.Players {
		map1[v.Name] = true
	}

	for _, v := range playerInfo2.Players {
		map2[v.Name] = true
	}

	//find players who joined and left since last update
	left, joined := symmetricDifference(map1, map2)
	return left, joined
}

//return two slices of strings:
//everything in a not in b and everything in b not in a
//or all players who left and joined server since last a2s query
func symmetricDifference(a, b map[string]bool) ([]string, []string) {
	return difference(a, b), difference(b, a)
}

//everything in a not in b
func difference(a, b map[string]bool) []string {
	rslt := []string{}

	for k, _ := range a {

		if !b[k] {
			rslt = append(rslt, k)
		}
	}
	return rslt
}
