package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type result struct {
	name   string
	played int
	won    int
	lost   int
	draw   int
	points int
}

//Tally calculates the tournament results and copies it to writer the output is nil if no error encountred
func Tally(reader io.Reader, writer io.Writer) error {
	ma := make(map[string]*result, 4)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		s := strings.Split(line, ";")

		if len(s) != 3 {
			return fmt.Errorf("error wrong format")
		}
		team1 := s[0]
		team2 := s[1]

		if _, ok := ma[team1]; !ok {
			ma[team1] = &result{
				name: team1,
			}
		}

		if _, ok := ma[team2]; !ok {
			ma[team2] = &result{
				name: team2,
			}
		}
		r1 := ma[team1]
		r2 := ma[team2]
		r1.played++
		r2.played++
		switch s[2] {

		case "win":
			{
				r1.won++
				r2.lost++
				r1.points += 3
			}
		case "loss":
			{
				r1.lost++
				r2.won++
				r2.points += 3
			}
		case "draw":
			{
				r1.draw++
				r2.draw++
				r1.points++
				r2.points++
			}
		default:
			return fmt.Errorf("error not win, draw or loss")
		}
	}

	teams := make([]result, 4, len(ma))
	for _, res := range ma {
		teams = append(teams, *res)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points != teams[j].points {
			return teams[i].points > teams[j].points
		}
		return teams[i].name < teams[j].name
	})
	fmt.Fprintf(writer, "%-31s| %2s | %2c | %2c | %2c | %2c\n", "Team", "MP", 'W', 'D', 'L', 'P')
	for _, res := range teams {
		if res.name == "" {
			continue
		}
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n", res.name, res.played, res.won, res.draw, res.lost,
			res.points)
	}
	return nil
}
