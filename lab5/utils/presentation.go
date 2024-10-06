package utils

import (
	"bufio"
	"math/big"
	"os"
	"strings"
)

func ReadPointsFromFile(filename string) ([]Point, error) {
	var points []Point
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			continue
		}
		x := new(big.Int)
		y := new(big.Int)
		if _, ok := x.SetString(strings.TrimSpace(coords[0]), 10); !ok {
			continue
		}
		if _, ok := y.SetString(strings.TrimSpace(coords[1]), 10); !ok {
			continue
		}
		points = append(points, Point{X: x, Y: y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}
