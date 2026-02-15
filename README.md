# Project Name

## About
This is CLI mini-game on Golang to guess the pseudo-random number.

## Files
- `go.mod` - Keep dependencies
- `main.go` - Entry point of the application
- `results.json` - Saved results in json
- `README.md` - Project documentation

## How to Launch

### Prerequisites
- Go 1.19+

### Installation
```bash
git clone <https://github.com/tiblocko2/guess_number>
cd guess_number
go mod download
```

### Running the Program
```bash
go run main.go
```

Or build and run:
```bash
go build -o program
./program
```

## Usage
```bash
$ go run main.go
Choose the difficult:
 1. Easy: 1-50 15 tries
 2. Medium: 1-100 10 tries
 3. Hard: 1-200 5 tries
1
48
Pick an option:
 1: Start game
 2:Exit
1
Enter number:
50
Your #1 try: 50
My number is smaller. but your's CLOSE
Enter number:
48
Your #2 try: 48
[50 48]
Good game. Wanna play again? y/n
n
```