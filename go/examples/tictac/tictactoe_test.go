package tictactoe

import (
	"os"
	"testing"
)

func TestBoard(t *testing.T) {
	t.Run("board size", func(t *testing.T) {
		b := NewBoard()
		b.SetSize(10)

		if len(b.GetBoard()) != 10 {
			t.Errorf("board size = %d; want 10", len(b.GetBoard()))
		}
	})

	t.Run("put symbol", func(t *testing.T) {
		b := NewBoard()
		b.SetSize(3)

		if err := b.PutSymbol(1, 2, "X"); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if err := b.PutSymbol(2, 3, "O"); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if b.GetBoard()[1][2] != "O" {
			t.Errorf("symbol at (2,3) = %s; want O", b.GetBoard()[1][2])
		}
	})

	t.Run("horizontal winner", func(t *testing.T) {
		b := NewBoard()
		b.SetSize(3)
		b.PutSymbol(1, 1, "X")
		b.PutSymbol(1, 2, "O")
		b.PutSymbol(2, 1, "X")
		b.PutSymbol(2, 3, "O")
		b.PutSymbol(3, 1, "X")

		winner, err := b.GetWinner()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if winner != "X" {
			t.Errorf("winner = %s; want X", winner)
		}
	})
}

func TestLoadFile(t *testing.T) {
	// Create temporary test files
	boardContent := []byte("4")
	symbolsContent := []byte("X,O,A")

	err := os.WriteFile("board.txt", boardContent, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("board.txt")

	err = os.WriteFile("symbols.txt", symbolsContent, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("symbols.txt")

	tests := []struct {
		name    string
		file    string
		want    string
		wantErr bool
	}{
		{"existing file", "board.txt", "4", false},
		{"non-existent file", "board2.txt", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := LoadFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if content != tt.want {
				t.Errorf("LoadFile() = %v, want %v", content, tt.want)
			}
		})
	}
}

func TestGetBoardSize(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    int
		wantErr bool
	}{
		{"valid size", "4", 4, false},
		{"invalid size large", "11", 0, true},
		{"empty content", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBoardSize(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBoardSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSymbols(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []string
		wantErr bool
	}{
		{"valid symbols", "X,O,A", []string{"X", "O", "A"}, false},
		{"invalid symbols", "A,B,C,D", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSymbols(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbols() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !stringSliceEqual(got, tt.want) {
				t.Errorf("GetSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInputFromConsole(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantX   int
		wantY   int
		wantErr bool
	}{
		{"valid input", "2,3", 2, 3, false},
		{"empty input", "", 0, 0, true},
		{"invalid format", "10", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInputFromConsole(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInputFromConsole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.X != tt.wantX || got.Y != tt.wantY {
					t.Errorf("GetInputFromConsole() = {%v,%v}, want {%v,%v}",
						got.X, got.Y, tt.wantX, tt.wantY)
				}
			}
		})
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
