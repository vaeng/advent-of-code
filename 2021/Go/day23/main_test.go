package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getDistance(t *testing.T) {
	type args struct {
		from position
		to   position
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"L-AF", args{L, AF}, 2},
		{"LL-AF", args{LL, AF}, 3},
		{"L-AB", args{L, AB}, 3},
		{"LL-AB", args{LL, AB}, 4},
		{"L-BF", args{L, BF}, 4},
		{"LL-BF", args{LL, BF}, 5},
		{"L-BB", args{L, BB}, 5},
		{"LL-BB", args{LL, BB}, 6},
		{"AtB-CB", args{AtB, CB}, 5},
		{"DF-R", args{DF, R}, 2},
		{"R-CF", args{R, CF}, 4},
		{"R-AF", args{R, AF}, 8},
		{"CtD-DB", args{CtD, DB}, 3},
		{"DB-CtD", args{DB, CtD}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("getDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamestate_isOccupied(t *testing.T) {
	type args struct {
		pos position
	}
	gsCorridorEmpty := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	gsCorridorFull := gamestate{LL, L, AtB, BtC, CtD, R, RR, DB}
	tests := []struct {
		name string
		gs   *gamestate
		args args
		want bool
	}{
		{"AF empty", &gsCorridorFull, args{AF}, false},
		{"BB full", &gsCorridorEmpty, args{BB}, true},
		{"CB empty", &gsCorridorFull, args{CB}, false},
		{"DF full", &gsCorridorEmpty, args{DF}, true},
		{"L empty", &gsCorridorEmpty, args{L}, false},
		{"RR full", &gsCorridorFull, args{RR}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.isOccupied(tt.args.pos); got != tt.want {
				t.Errorf("gamestate.isOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleDestinations(t *testing.T) {
	type args struct {
		pos  position
		amph amphipod
	}
	coridor := []position{LL, L, AtB, BtC, CtD, R, RR}
	tests := []struct {
		name string
		args args
		want []position
	}{
		{"Apod From L", args{L, A}, []position{AF, AB}},
		{"Apod From LL", args{LL, A}, []position{AF, AB}},
		{"Apod From AtB", args{AtB, A}, []position{AF, AB}},
		{"Dpod From AtB", args{AtB, D}, []position{DF, DB}},
		{"Apod From CtD", args{CtD, A}, []position{AF, AB}},
		{"Apod From R", args{R, A}, []position{AF, AB}},
		{"Bpod From RR", args{RR, B}, []position{BF, BB}},
		{"Cpod From BtC", args{BtC, C}, []position{CF, CB}},
		{"Apod From AB", args{AB, A}, []position{}},
		{"Apod From AF", args{AF, A}, coridor},
		{"Bpod From BF", args{BF, B}, coridor},
		{"Cpod From DF", args{DF, C}, coridor},
		{"Dpod From CB", args{CB, D}, coridor},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleDestinations(tt.args.pos, tt.args.amph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("possibleDestinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWayInPosVisited(t *testing.T) {
	type args struct {
		from position
		to   position
	}
	tests := []struct {
		name string
		args args
		want []position
	}{
		{"From LL, AF", args{LL, AF}, []position{L, AF}},
		{"From L, AB", args{L, AB}, []position{AF, AB}},
		{"From AtB, AB", args{AtB, AB}, []position{AF, AB}},
		{"From AtB, BB", args{AtB, BB}, []position{BF, BB}},
		{"From LL, CB", args{LL, CB}, []position{L, AtB, BtC, CF, CB}},
		{"From L, CF", args{L, CF}, []position{AtB, BtC, CF}},
		{"From CF, L", args{CF, L}, []position{AtB, BtC, L}},
		{"From DB, AtB", args{DB, AtB}, []position{DF, CtD, BtC, AtB}},
		{"From BF, RR", args{BF, RR}, []position{BtC, CtD, R, RR}},
		{"From RR, BF", args{RR, BF}, []position{BtC, CtD, R, BF}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWayInPosVisited(tt.args.from, tt.args.to)
			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			want := tt.want
			sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })
			if !reflect.DeepEqual(got, want) {
				t.Errorf("getWayInPosVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamestate_isWayPossible(t *testing.T) {
	type args struct {
		from position
		to   position
	}
	gsCorridorEmpty := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	gsCorridorBtCBlock := gamestate{AF, AB, BF, BB, CF, CB, DF, BtC}
	gsCorridorFull := gamestate{LL, L, AtB, BtC, CtD, R, RR, DB}
	tests := []struct {
		name string
		gs   *gamestate
		args args
		want bool
	}{
		{"Empty: AF - RR", &gsCorridorEmpty, args{AF, RR}, true},
		{"Empty: AF - AtB", &gsCorridorEmpty, args{AF, AtB}, true},
		{"Full: AF - AtB", &gsCorridorFull, args{AF, AtB}, false},
		{"Empty: AB - AtB", &gsCorridorEmpty, args{AB, AtB}, false},
		{"Full: AtB - AB", &gsCorridorFull, args{AtB, AB}, true},
		{"Empty: CF - R", &gsCorridorEmpty, args{CF, R}, true},
		{"Empty: CF - CtD", &gsCorridorEmpty, args{CF, CtD}, true},
		{"Empty: CF - BtC", &gsCorridorEmpty, args{CF, BtC}, true},
		{"Full: CF - R", &gsCorridorFull, args{CF, R}, false},
		{"MidBlock: AF - R", &gsCorridorBtCBlock, args{AF, R}, false},
		{"MidBlock: AF - BtC", &gsCorridorBtCBlock, args{AF, BtC}, false},
		{"MidBlock: AF - L", &gsCorridorBtCBlock, args{AF, L}, true},
		{"MidBlock: CF - BtC", &gsCorridorBtCBlock, args{CF, BtC}, false},
		{"MidBlock: CF - L", &gsCorridorBtCBlock, args{CF, L}, false},
		{"MidBlock: CF - R", &gsCorridorBtCBlock, args{CF, R}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.isWayPossible(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("gamestate.isWayPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamestate_nextPossibleGamestates(t *testing.T) {
	gsCorridorEmpty := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	gsCorridorBtCBlock := gamestate{AF, AB, BF, BB, CF, CB, DF, BtC}
	tests := []struct {
		name string
		gs   *gamestate
		want []gamestate
	}{
		{"Start", &gsCorridorEmpty, []gamestate{}},
		{"Blocked", &gsCorridorBtCBlock, []gamestate{
			{AF, AB, BF, BB, CF, CB, CtD, BtC},
			{AF, AB, BF, BB, CF, CB, R, BtC},
			{AF, AB, BF, BB, CF, CB, RR, BtC},
		}},
		{"Only BF missing", &gamestate{AB, AF, BtC, BB, CF, CB, DF, DB}, []gamestate{{AB, AF, BF, BB, CF, CB, DF, DB}}},
		{"Only BF missing 2", &gamestate{AB, AF, BB, BtC, CF, CB, DF, DB}, []gamestate{{AB, AF, BB, BF, CF, CB, DF, DB}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.gs.nextPossibleGamestates()
			want := tt.want
			// test if all want are in got
		gs:
			for _, g := range got {
				gn := normalize(g)
				for _, w := range want {
					wn := normalize(w)
					if wn == gn {
						continue gs
					}
				}
				// fmt.Printf("%v not in want\n", g)
				t.Errorf("gamestate.nextPossibleGamestates() = %v, want %v", got, tt.want)
			}
			// test if all got are in want
		ws:
			for _, w := range want {
				wn := normalize(w)
				for _, g := range got {
					gn := normalize(g)
					if wn == gn {
						continue ws
					}
				}
				// fmt.Printf("%v not in got\n", w)
				t.Errorf("gamestate.nextPossibleGamestates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveCost(t *testing.T) {
	type args struct {
		current   gamestate
		nextState gamestate
	}
	gsCorridorEmpty := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"Normalizing Test", args{gamestate{AB, L, BB, BF, CB, CF, DB, DF}, gamestate{L, AB, BF, BB, CF, CB, DF, DB}}, 0},
		{"Move A from Start to L", args{gsCorridorEmpty, gamestate{L, AB, BF, BB, CF, CB, DF, DB}}, 2},
		{"Move A from Start to LL", args{gsCorridorEmpty, gamestate{LL, AB, BF, BB, CF, CB, DF, DB}}, 3},
		{"Move B from Start to LL", args{gsCorridorEmpty, gamestate{AF, AB, LL, BB, CF, CB, DF, DB}}, 50},
		{"Move C from Start to LL", args{gsCorridorEmpty, gamestate{AF, AB, BF, BB, LL, CB, DF, DB}}, 700},
		{"Move C from AtB to CF", args{gamestate{AF, AB, BF, BB, AtB, CB, DF, DB}, gamestate{AF, AB, BF, BB, CF, CB, DF, DB}}, 400},
		{"Move D from Start to LL", args{gsCorridorEmpty, gamestate{AF, AB, BF, BB, CF, CB, LL, DB}}, 9000},
		{"Move D from Start to R", args{gsCorridorEmpty, gamestate{AF, AB, BF, BB, CF, CB, R, DB}}, 2000},
		{"Move D from Start to RR", args{gsCorridorEmpty, gamestate{AF, AB, BF, BB, CF, CB, RR, DB}}, 3000},
		{"Move B from BtC to Start", args{gsCorridorEmpty, gamestate{AF, AB, BtC, BB, CF, CB, DF, DB}}, 20},
		{"Only BF missing", args{gamestate{AB, AF, BtC, BB, CF, CB, DF, DB}, gamestate{AB, AF, BF, BB, CF, CB, DF, DB}}, 20},
		{"Only BF missing 2", args{gamestate{AB, AF, BB, BtC, CF, CB, DF, DB}, gamestate{AB, AF, BB, BF, CF, CB, DF, DB}}, 20},
		{"Only BF missing 2, normalized", args{gamestate{AB, AF, BB, BtC, CF, CB, DF, DB}, gamestate{AB, AF, BF, BB, CF, CB, DF, DB}}, 20},
		{"Only AF missing", args{gamestate{L, AB, BF, BB, CF, CB, DF, DB}, gamestate{AB, AF, BF, BB, CF, CB, DF, DB}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveCost(tt.args.current, tt.args.nextState); got != tt.want {
				t.Errorf("moveCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aStar2(t *testing.T) {
	type args struct {
		start gamestate
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"AoC example", args{gamestate{AB, DB, AF, CF, BF, CB, BB, DF}}, 12521},
		{"Only AF missing", args{gamestate{L, AB, BF, BB, CF, CB, DF, DB}}, 2},
		{"Only AF missing 2", args{gamestate{AB, RR, BF, BB, CF, CB, DF, DB}}, 9},
		{"Only BF missing", args{gamestate{AB, AF, BtC, BB, CF, CB, DF, DB}}, 20},
		{"Only BF missing 2", args{gamestate{AB, AF, BB, BtC, CF, CB, DF, DB}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aStar2(tt.args.start); got != tt.want {
				t.Errorf("aStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamestate_isComplete(t *testing.T) {
	tests := []struct {
		name string
		gs   *gamestate
		want bool
	}{
		{"Everything in Order", &gamestate{AF, AB, BF, BB, CF, CB, DF, DB}, true},
		{"Everything mixed", &gamestate{AB, AF, BB, BF, CB, CF, DB, DF}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.isComplete(); got != tt.want {
				t.Errorf("gamestate.isComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gamestate_amphipodsDone(t *testing.T) {
	type args struct {
		pod amphipod
	}
	orderedFinish := gamestate{AF, AB, BF, BB, CF, CB, DF, DB}
	inverseFinish := gamestate{AB, AF, BB, BF, CB, CF, DB, DF}
	oneOutside := gamestate{AB, AtB, LL, BF, L, CF, DB, R} //
	tests := []struct {
		name string
		gs   *gamestate
		args args
		want bool
	}{
		{"all done, check A ordered", &orderedFinish, args{A}, true},
		{"all done, check A reverse", &inverseFinish, args{A}, true},
		{"one outside, check A reverse", &oneOutside, args{A}, false},
		{"all done, check B ordered", &orderedFinish, args{B}, true},
		{"all done, check B reverse", &inverseFinish, args{B}, true},
		{"one outside, check B reverse", &oneOutside, args{B}, false},
		{"all done, check C ordered", &orderedFinish, args{C}, true},
		{"all done, check C reverse", &inverseFinish, args{C}, true},
		{"one outside, check C reverse", &oneOutside, args{C}, false},
		{"all done, check D ordered", &orderedFinish, args{D}, true},
		{"all done, check D reverse", &inverseFinish, args{D}, true},
		{"one outside, check D reverse", &oneOutside, args{D}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.amphipodsDone(tt.args.pod); got != tt.want {
				t.Errorf("gamestate.amphipodsDone() = %v, want %v", got, tt.want)
			}
		})
	}
}
