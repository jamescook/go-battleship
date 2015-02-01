package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

type GameBoard struct {
	Coordinates     [5][5]int
	ShipCoordinates Point
}

func (p *Player) PlaceShip(x, y int) Point {
	coords := Point{x, y}
	p.ShipCoordinates = coords
	return coords
}

type Player struct {
	Name string
	*GameBoard
	Dead  bool
	Human bool
}

func (p *Player) Attack(x, y int) int {
	if p.Dead {
		return 1
	}
	if x >= len(p.Coordinates) || y >= len(p.Coordinates) {
		fmt.Printf("You can't attack there\n")
		return 0
	}

	if p.Coordinates[x][y] == 1 {
		if p.Human {
			fmt.Printf("You already attacked at %d, %d\n", x, y)
		}
	} else {
		p.Coordinates[x][y] = 1
	}

	if p.ShipCoordinates.X == x && p.ShipCoordinates.Y == y {
		fmt.Printf("You sunk %s at %d,%d !\n", p.Name, x, y)
		p.Dead = true
	}

	return 1
}

func GuessLocation() (int, int) {
	var (
		row, col int
	)
	fmt.Printf("Attack Position: ")
	fmt.Scanf("%d,%d", &row, &col)
	//if err.Error() != "" {
	//fmt.Println("Error", err.Error())
	//}
	return row, col
}

func Reseed() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	return true
}

func Play(player *Player, computer *Player) bool {
	row, col := GuessLocation()
	computer.Attack(row, col)
	if computer.Dead == false {
		fmt.Println("Missed!")
		fmt.Println(computer.Name, computer.Coordinates)
		Reseed()
		randomX := rand.Intn(len(player.Coordinates))
		randomY := rand.Intn(len(player.Coordinates))
		player.Attack(randomX, randomY)
		if player.Dead == false {
			fmt.Println(player.Name, player.Coordinates)
			Play(player, computer)
		}
	}

	return true
}

func AskPlayerName(playerName *string) string {
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", playerName)
	return *playerName
}

func AskShipLocation(playerRow *int, playerCol *int) bool {
	fmt.Printf("Where is your ship on a 5x5 board? ")
	fmt.Scanf("%d,%d", playerRow, playerCol)

	return true
}

func main() {
	var (
		playerRow  int
		playerCol  int
		playerName string
	)

	fmt.Println("Let's play Battleship!\n")
	Reseed()

	AskPlayerName(&playerName)
	p1 := Player{playerName, new(GameBoard), false, true}
	p2 := Player{"Computer", new(GameBoard), false, false}

	AskShipLocation(&playerRow, &playerCol)
	p1.PlaceShip(playerRow, playerCol)

	randomX := rand.Intn(len(p1.Coordinates))
	randomY := rand.Intn(len(p1.Coordinates))
	p2.PlaceShip(randomX, randomY)

	fmt.Println(p1.Name, p1.Coordinates)
	fmt.Println(p2.Name, p2.Coordinates)

	Play(&p1, &p2)
}
