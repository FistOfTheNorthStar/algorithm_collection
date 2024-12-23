package tictactoe

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	empty = " "
	draw  = "DRAW"
)

// Board represents a tic-tac-toe board
type Board struct {
	board             [][]string
	numOfPlays        int
	numOfPlaysAllowed int
}

// NewBoard creates a new board instance
func NewBoard() *Board {
	return &Board{}
}

// SetSize initializes the board with given size
func (b *Board) SetSize(size int) {
	b.numOfPlaysAllowed = size * size
	b.board = make([][]string, size)
	for x := range b.board {
		b.board[x] = make([]string, size)
		for y := range b.board[x] {
			b.board[x][y] = empty
		}
	}
}

// GetBoard returns the current board state
func (b *Board) GetBoard() [][]string {
	return b.board
}

// PutSymbol places a symbol on the board
func (b *Board) PutSymbol(x, y int, symbol string) error {
	if b.board == nil {
		return fmt.Errorf("board is not initialized")
	}

	if x < 1 || x > len(b.board) {
		return fmt.Errorf("X coordinate invalid, must be between 1 and %d", len(b.board))
	}

	if y < 1 || y > len(b.board) {
		return fmt.Errorf("Y coordinate invalid, must be between 1 and %d", len(b.board))
	}

	if b.board[x-1][y-1] != empty {
		return fmt.Errorf("coordinates are busy")
	}

	b.board[x-1][y-1] = symbol
	b.numOfPlays++
	return nil
}

// Print displays the current board state
func (b *Board) Print() error {
	if b.board == nil {
		return fmt.Errorf("board is not initialized")
	}

	size := len(b.board)
	for y := size - 1; y >= 0; y-- {
		row := make([]string, size)
		for x := 0; x < size; x++ {
			row[x] = b.board[x][y]
		}
		fmt.Println(strings.Join(row, "|"))
	}
	return nil
}

// GetWinner returns the winner symbol or draw
func (b *Board) GetWinner() (string, error) {
	if b.board == nil {
		return "", fmt.Errorf("board is not initialized")
	}

	// Check horizontal lines
	if winner := b.winnerInHorizontalLine(); winner != "" {
		return winner, nil
	}

	// Check vertical lines
	if winner := b.winnerInVerticalLine(); winner != "" {
		return winner, nil
	}

	// Check diagonals
	if winner := b.winnerInDiagonalBottomLeft(); winner != "" {
		return winner, nil
	}

	if winner := b.winnerInDiagonalTopLeft(); winner != "" {
		return winner, nil
	}

	// Check for draw
	if b.numOfPlays == b.numOfPlaysAllowed {
		return draw, nil
	}

	return "", nil
}

func (b *Board) winnerInHorizontalLine() string {
	for y := 0; y < len(b.board); y++ {
		symbol := b.board[0][y]
		if symbol == empty {
			continue
		}

		counter := 1
		for x := 1; x < len(b.board); x++ {
			if symbol == b.board[x][y] {
				counter++
			}
		}

		if counter == len(b.board) {
			return symbol
		}
	}
	return ""
}

func (b *Board) winnerInVerticalLine() string {
	for x := 0; x < len(b.board); x++ {
		symbol := b.board[x][0]
		if symbol == empty {
			continue
		}

		counter := 1
		for y := 1; y < len(b.board); y++ {
			if symbol == b.board[x][y] {
				counter++
			}
		}

		if counter == len(b.board) {
			return symbol
		}
	}
	return ""
}

func (b *Board) winnerInDiagonalBottomLeft() string {
	symbol := b.board[0][0]
	if symbol == empty {
		return ""
	}

	counter := 1
	for idx := 1; idx < len(b.board); idx++ {
		if symbol == b.board[idx][idx] {
			counter++
		}
	}

	if counter == len(b.board) {
		return symbol
	}
	return ""
}

func (b *Board) winnerInDiagonalTopLeft() string {
	symbol := b.board[0][len(b.board)-1]
	if symbol == empty {
		return ""
	}

	counter := 1
	for idx := 1; idx < len(b.board); idx++ {
		if symbol == b.board[idx][len(b.board)-1-idx] {
			counter++
		}
	}

	if counter == len(b.board) {
		return symbol
	}
	return ""
}

// Input represents console input coordinates
type Input struct {
	X int
	Y int
}

// LoadFile reads content from a file
func LoadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("file %s not found", fileName)
	}
	return string(content), nil
}

// GetBoardSize validates and returns board size from content
func GetBoardSize(content string) (int, error) {
	if content == "" {
		return 0, fmt.Errorf("invalid setting for the board")
	}

	size, err := strconv.Atoi(strings.TrimSpace(content))
	if err != nil {
		return 0, fmt.Errorf("invalid board size format")
	}

	if size < 3 || size > 9 {
		return 0, fmt.Errorf("invalid setting for the board")
	}

	return size, nil
}

// GetSymbols splits and validates symbol content
func GetSymbols(content string) ([]string, error) {
	symbols := strings.Split(content, ",")
	if len(symbols) != 3 {
		return nil, fmt.Errorf("invalid settings for symbols")
	}
	return symbols, nil
}

// GetInputFromConsole parses console input
func GetInputFromConsole(inputFromConsole string) (*Input, error) {
	if inputFromConsole == "" {
		return nil, fmt.Errorf("invalid input from console")
	}

	pattern := regexp.MustCompile(`[0-9]*\.?,[0-9]+`)
	if !pattern.MatchString(inputFromConsole) {
		return nil, fmt.Errorf("invalid input format")
	}

	parts := strings.Split(inputFromConsole, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid input format")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid X coordinate")
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid Y coordinate")
	}

	return &Input{X: x, Y: y}, nil
}
