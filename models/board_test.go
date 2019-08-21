package models

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		dimensions int
		startX     int
		startY     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - create new board ",
			args:    args{startX: 0, startY: 9, dimensions: 10},
			wantErr: false,
		},
		{
			name:    "fail - start coordinates greater than dimension",
			args:    args{startX: 0, startY: 11, dimensions: 10},
			wantErr: true,
		},
		{
			name:    "fail - negative start coordinates",
			args:    args{startX: -1, startY: 0, dimensions: 10},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBoard(tt.args.dimensions, tt.args.startX, tt.args.startY)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_board_Move(t *testing.T) {
	brd, _ := NewBoard(10, 0, 0)
	type args struct {
		x       int
		y       int
		visited int
	}
	tests := []struct {
		name    string
		b       board
		args    args
		wantErr bool
	}{
		{
			name:    "success - move point",
			args:    args{x: 1, y: 0, visited: 0},
			wantErr: false,
			b:       brd,
		},
		{
			name:    "fail - move out of bounds",
			args:    args{x: -10, y: 0, visited: 0},
			wantErr: true,
			b:       brd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Move(tt.args.x, tt.args.y, tt.args.visited); (err != nil) != tt.wantErr {
				t.Errorf("NewBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_board_Visited(t *testing.T) {
	brd, _ := NewBoard(10, 0, 0)
	tests := []struct {
		name string
		b    board
		want int
	}{
		{
			name: "success - count visits",
			want: 1,
			b:    brd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Visited(); got != tt.want {
				t.Errorf("board.Visited() = %v, want %v", got, tt.want)
			}
		})
	}
}
