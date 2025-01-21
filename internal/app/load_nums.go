package app

import (
	"bufio"
	"context"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetSliceNums(ctx context.Context, path string) ([]int, error) {
	var nums []int

	file, err := os.Open(path)
	if err != nil {
		log.WithField("input file", path).
			WithError(err).
			Panic("read input file")
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.WithField("input file", path).
				WithField("input line", x).
				WithError(err)
			return nil, err
		}
		nums = append(nums, x)
	}

	return nums, nil
}
