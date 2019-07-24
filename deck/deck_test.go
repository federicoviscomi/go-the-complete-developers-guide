package main

import (
	"reflect"
	"testing"
)

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		{
			name: "create a new deck",
			want: deck([]string{"Ace of ♣ clubs", "2 of ♣ clubs", "3 of ♣ clubs", "4 of ♣ clubs", "5 of ♣ clubs", "6 of ♣ clubs", "7 of ♣ clubs", "8 of ♣ clubs", "9 of ♣ clubs", "10 of ♣ clubs", "Jack of ♣ clubs", "Queen of ♣ clubs", "King of ♣ clubs", "Ace of ♦ diamonds", "2 of ♦ diamonds", "3 of ♦ diamonds", "4 of ♦ diamonds", "5 of ♦ diamonds", "6 of ♦ diamonds", "7 of ♦ diamonds", "8 of ♦ diamonds", "9 of ♦ diamonds", "10 of ♦ diamonds", "Jack of ♦ diamonds", "Queen of ♦ diamonds", "King of ♦ diamonds", "Ace of ♥ hearts", "2 of ♥ hearts", "3 of ♥ hearts", "4 of ♥ hearts", "5 of ♥ hearts", "6 of ♥ hearts", "7 of ♥ hearts", "8 of ♥ hearts", "9 of ♥ hearts", "10 of ♥ hearts", "Jack of ♥ hearts", "Queen of ♥ hearts", "King of ♥ hearts", "Ace of ♠ spades", "2 of ♠ spades", "3 of ♠ spades", "4 of ♠ spades", "5 of ♠ spades", "6 of ♠ spades", "7 of ♠ spades", "8 of ♠ spades", "9 of ♠ spades", "10 of ♠ spades", "Jack of ♠ spades", "Queen of ♠ spades", "King of ♠ spades"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
			got := newDeck()
			if !equals(got, tt.want) {
				t.Errorf("\nnewDeck() = %v\nwant %v", got.toString(), tt.want.toString())
			}
		})
	}
}

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("deck.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_shuffleLinear(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		{"shuffle test", newDeck()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffleLinear()
			tt.d.print()
		})
	}
}

func Test_randomInteger(t *testing.T) {
	type args struct {
		startRangeIncluded int
		endRangeExcluded   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomInteger(tt.args.startRangeIncluded, tt.args.endRangeExcluded); got != tt.want {
				t.Errorf("randomInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_shuffle_nSquare(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffle_nSquare()
		})
	}
}

func Test_randomDecision(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomDecision(); got != tt.want {
				t.Errorf("randomDecision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deal(t *testing.T) {
	type args struct {
		d        deck
		handSize int
	}
	tests := []struct {
		name  string
		args  args
		want  deck
		want1 deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deal(tt.args.d, tt.args.handSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deck_print(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.print()
		})
	}
}

func Test_deck_toByteArray(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deck.toByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_saveDeckToFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveDeckToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("deck.saveDeckToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deck_length(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.length(); got != tt.want {
				t.Errorf("deck.length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadDeckFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    deck
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadDeckFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadDeckFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
