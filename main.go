package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	Game_date string `json:"game_date"`
	Status    string `json:"status"`
	Tries     int    `json:"tries"`
}

func NewResult(game_date time.Time, status bool, tries int) Result {
	if status {
		date_formatted := game_date.Format("2006-01-02 15:04:05")
		return Result{
			date_formatted,
			"win",
			tries,
		}
	} else {
		date_formatted := game_date.Format("2006-01-02 15:04:05")
		return Result{
			date_formatted,
			"lose",
			tries,
		}
	}
}

func main() {
	for i := true; i != false; {
		game_date := time.Now()
		rand.Seed(time.Now().UnixNano())
		rand_range, tries := ChooseDifficult()
		number := rand.Intn(rand_range)
		fmt.Println(number)

		for opt := 0; opt != 2; {
			fmt.Println("Pick an option:\n 1: Start game\n 2:Exit")
			winner := false
			fmt.Scan(&opt)
			switch opt {
			case 1:
				var n int
				tries_list := make([]int, 0)
				for g := 1; g <= tries; g++ {
					n = CheckInput(g)
					tries_list = append(tries_list, n)
					if CheckGuess(n, number) {
						winner = true
						g = tries + 1
					}
				}
				fmt.Println(tries_list)
				data := NewResult(game_date, winner, len(tries_list))
				SaveResult(data)
				if NewGame(winner) {
					opt = 2
				} else {
					i = false
					opt = 2
				}
			case 2:
				opt = 2
				return
			default:
				fmt.Println("Wrong input. Try again:)")
			}

		}
	}
}

func CheckInput(n int) int {
	var s string
	fmt.Println("Enter number:")
	var input int
	for {
		_, err := fmt.Scan(&s)
		input, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Not a valid number. Try again")
		} else {
			break
		}
	}
	out_guess := fmt.Sprintf("\033[33mYour #%d try: %d\033[0m", n, input)
	fmt.Println(out_guess)
	return input
}

func CheckGuess(guess int, ans int) bool {
	if guess != ans {
		if guess > ans {
			if guess-ans < 6 {
				fmt.Println("\033[35mMy number is smaller. but your's CLOSE \033[0m")
				return false
			} else if guess-ans < 15 {
				fmt.Println("\033[34mMy number is smaller. but your's close\033[0m")
				return false
			} else {
				fmt.Println("\033[36mMy number is smaller.\033[0m")
				return false
			}
		} else {
			if ans-guess < 6 {
				fmt.Println("\033[35mMy number is bigger. but your's CLOSE\033[0m")
				return false
			} else if ans-guess < 16 {
				fmt.Println("\033[34mMy number is bigger. but your's close\033[0m")
				return false
			} else {
				fmt.Println("\033[36mMy number is bigger.\033[0m")
				return false
			}
		}
	} else {
		return true
	}
}

func ChooseDifficult() (int, int) {
	var rand_range, tries int
	for i := true; i != false; {
		fmt.Println("Choose the difficult:\n \033[32m1. Easy: 1-50 15 tries\n \033[33m2. Medium: 1-100 10 tries\n \033[31m3. Hard: 1-200 5 tries\033[0m")
		var input int
		fmt.Scan(&input)
		switch input {
		case 1:
			rand_range = 50
			tries = 15
			i = false
		case 2:
			rand_range = 100
			tries = 10
			i = false
		case 3:
			rand_range = 200
			tries = 5
			i = false
		default:
			fmt.Println("Wrong input")
		}
	}
	return rand_range, tries
}

func NewGame(winner bool) bool {
	if winner {
		fmt.Println("\033[32mGood game. Wanna play again? y/n\033[0m")
		for {
			var input string
			fmt.Scan(&input)
			switch strings.ToLower(input) {
			case "y":
				return true
			case "n":
				return false
			default:
				fmt.Println("Wrong input.")
			}
		}
	} else {
		fmt.Println("\033[31mNice try. Wanna try again? y/n\033[0m")
		for {
			var input string
			fmt.Scan(&input)
			switch strings.ToLower(input) {
			case "y":
				return true
			case "n":
				return false
			default:
				fmt.Println("Wrong input.")
			}
		}
	}
}

func SaveResult(data Result) {
	f, err := os.OpenFile("results.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", " ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
