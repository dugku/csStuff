package main

import (
	"encoding/json"
	"fmt"
	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Position struct {
	Name   string   `json:"name"`
	Points []vector `json:"points"`
}

type PositionData struct {
	Positions []Position `json:"positions"`
}

type singleVect struct {
	x, y, z float64
}

func main() {

	//var positiondata PositionData
	//JSONfile := ""
	//posData := jsonLoader(JSONfile, positiondata)

	var callouts []string
	calloutsptr := &callouts

	f, err := os.Open("C:\\Users\\Mike\\Desktop\\csStuff\\Heat\\nouns-fe-vs-lumen-fe-m1-anubis.dem")
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	killcount := 0
	killcountptr := &killcount
	// Register handler on kill events

	p.RegisterEventHandler(func(e events.MatchStartedChanged) {
		mapName := p.Header().MapName

		jsonStuff := "C:\\Users\\Mike\\Desktop\\csStuff\\Heat\\mapsCoords"

		filepath.WalkDir(jsonStuff, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				if mapName == path {
					fmt.Println("here")
				}
			}

			return nil
		})
	})

	p.RegisterEventHandler(func(e events.Kill) {
		//calloutsLocal := []string{}
		*killcountptr = *killcountptr + 1
		victPoint := singleVect{
			x: e.Victim.Position().X,
			y: e.Victim.Position().Y,
		}

		for _, pos := range posData.Positions {
			inter, where := raycast(victPoint.x, victPoint.y, pos.Points, pos.Name)
			if inter == true {
				*calloutsptr = append(*calloutsptr, where)
			}
		}

		//fmt.Println(calloutsLocal)
	})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

	//off by one kill wtf this sucks gotta rewrite it all smh
	fmt.Println(*killcountptr)
	fmt.Println(len(*calloutsptr))
	//fmt.Println(*calloutsptr)

	countsEntries(calloutsptr, "a short")

}

func raycast(victx, victy float64, edges []vector, name string) (bool, string) {
	count := 0
	where := ""
	tolerance := 1e-7

	for i := 0; i < len(edges); i++ {
		curr := edges[i]
		next := edges[(i+1)%len(edges)]

		if curr.Y == next.Y {
			continue
		}

		if victy < math.Min(curr.Y, next.Y) || victy > math.Max(curr.Y, next.Y) {
			continue
		}

		xIntercept := (victy-curr.Y)*(next.X-curr.X)/(next.Y-curr.Y) + curr.X

		if victx < xIntercept+tolerance {
			if where == "" {
				where = name
			}
			count++
		}
	}

	return count%2 == 1, where
}

func countsEntries(calls *[]string, e string) {
	var n int
	for _, i := range *calls {
		temp := strings.ToLower(i)
		if temp == strings.ToLower(e) {
			n++
		}
	}
	fmt.Println(e, n)
}

func jsonLoader(file string, data PositionData) PositionData {
	jsonFile, err := os.ReadFile(file)
	check(err)

	e := json.Unmarshal(jsonFile, &data)

	check(e)

	return data
}

/*
(
func pointOnEdge(victx, victy, x1, y1, x2, y2, tolerance float64) bool {

		if math.Abs(y2-y1) < tolerance {
			return math.Abs(ypoint-y1) < tolerance && xpoint >= math.Min(x1, x2) && xpoint <= math.Max(x1, x2)
		}

		if x1 == x2 {
			return math.Abs(xpoint-x1) < tolerance && ypoint >= math.Min(y1, y2) && ypoint <= math.Max(y1, y2)
		}

		m := (y2 - y1) / (x2 - x1)
		b := y1 - m*1

		return math.Abs(ypoint-(m*xpoint+b)) < tolerance && xpoint >= math.Min(x1, x2) && xpoint <= math.Max(x1, x2)
	}
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
