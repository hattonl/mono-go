package xfile

import (
	"bufio"
	"os"
	"strings"
)

func ReadLinesWithFunc(fileName string, lineOffset int, lines int, fn func(line string)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	index, scanned := 0, 0

	for scanner.Scan() {
		if index >= lineOffset {
			fn(scanner.Text())
			scanned++
		}
		index++
		if scanned >= lines {
			break
		}
	}
	return nil
}

func ReadLines(fileName string, lineOffset int, lines int) ([]string, error) {
	var linesBuffer []string
	err := ReadLinesWithFunc(fileName, lineOffset, lines, func(line string) {
		linesBuffer = append(linesBuffer, line)
	})
	return linesBuffer, err
}

func ReadFileCSV(fileName string, lineOffset int, lines int) ([][]string, error) {
	var values [][]string
	err := ReadLinesWithFunc(fileName, lineOffset, lines, func(line string) {
		vs := strings.Split(line, ",")
		values = append(values, vs)
	})
	return values, err
}

func ReadFileTSV(fileName string, lineOffset int, lines int) ([][]string, error) {
	var values [][]string
	err := ReadLinesWithFunc(fileName, lineOffset, lines, func(line string) {
		vs := strings.Split(line, "\t")
		values = append(values, vs)
	})
	return values, err
}
