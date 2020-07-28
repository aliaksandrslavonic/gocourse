package main

import "fmt"

type Jukebox interface {
	PlayTrack() string
	AddTrack() string
	RemoveTrack() string
	SwitchTrack() string
	StopTrack() string
}

type Track struct {
	Name   string
	Artist string
}

func (t *Track) PlayTrack() string {
	return fmt.Sprintf("Play track: %q - %s.", t.Name, t.Artist)
}

// type Square struct {
// 	Side float64
// }

// func (s *Square) Area() float64 {
// 	return s.Side * s.Side
// }

// func SumAreas(shapes ...Shape) float64 {
// 	var sum float64
// 	for _, shape := range shapes {
// 		sum += shape.Area()
// 	}
// 	return sum
// }

func main() {
	a := Track{
		Name:   "Tool",
		Artist: "Pneuma",
	}
	fmt.Println(a.PlayTrack())
}

//Tool - Pneuma Cover
