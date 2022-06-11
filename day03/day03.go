package day03

import (
	"bufio"
	"fmt"
	"github.com/akyrey/aoc-2021/utils"
	"math"
	"strconv"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func calculateRate(totalCount int, reportBitSum []int, mostCommon bool) []int {
	rate := make([]int, len(reportBitSum), cap(reportBitSum))

	for i := 0; i < len(reportBitSum); i++ {
		var value int
		if mostCommon {
			value = getMostCommonBit(totalCount, reportBitSum[i])
		} else {
			value = getLeastCommonBit(totalCount, reportBitSum[i])
		}

		rate[i] = value
	}

	return rate
}

func calculateDecimalRate(rate []int) int {
	binaryValue := arrayToString(rate[:], "")
	decimalValue, err := strconv.ParseInt(binaryValue, 2, 64)
	utils.CheckError(err)

	return int(decimalValue)
}

func getMostCommonBit(totalCount int, value int) int {
	zeroValueLimit := int(math.Round(float64(totalCount) / 2.0))

	if value < zeroValueLimit {
		return 0
	}

	return 1
}

func getLeastCommonBit(totalCount int, value int) int {
	zeroValueLimit := int(math.Round(float64(totalCount) / 2.0))

	if value < zeroValueLimit {
		return 1
	}

	return 0
}

func sumReportValues(reports [][]int, length int) []int {
	reportBitSum := make([]int, length, length)

	for _, v := range reports {
		for i := 0; i < length; i++ {
			reportBitSum[i] += v[i]
		}
	}

	return reportBitSum
}

func filterReportsByBitInPosition(reports [][]int, position int, mostCommon bool) [][]int {
	result := [][]int{}
	length := len(reports)
	// fmt.Printf("length %d ", length)
	sumReport := sumReportValues(reports, len(reports[0]))
	// fmt.Printf("sum report %v ", sumReport)
	filter := calculateRate(length, sumReport, mostCommon)
	// fmt.Printf("filter %v", filter)

	for _, v := range reports {
		if v[position] == filter[position] {
			result = append(result, v)
		}
	}

	return result
}

func filterByBit(reports [][]int, length int, mostCommon bool) []int {
	result := reports[:]

	for position := 0; position < length; position++ {
		if len(result) != 1 {
			result = filterReportsByBitInPosition(result, position, mostCommon)
		}
		// fmt.Printf("Position: %d, values: %v\n", position, result)
	}

	return result[0]
}

func Day03(test bool) {
	f, err := utils.GetFileToReadFrom(3, test)
	utils.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalCount := 0
	var reportBit [][]int
	var reportBitSum []int
	var length int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		length = len(split)
		if reportBitSum == nil {
			reportBitSum = make([]int, length, length)
		}
		currentReport := make([]int, length, length)

		fmt.Printf("Split %v\n", split)
		for i := 0; i < len(split); i++ {
			value, err := strconv.Atoi(split[i])
			utils.CheckError(err)

			reportBitSum[i] += value
			currentReport[i] = value
		}

		reportBit = append(reportBit[:], currentReport)
		totalCount++
	}

	gammaRateBinary := calculateRate(totalCount, reportBitSum, true)
	epsilonRateBinary := calculateRate(totalCount, reportBitSum, false)

	gammaRate := calculateDecimalRate(gammaRateBinary)
	epsilonRate := calculateDecimalRate(epsilonRateBinary)

	fmt.Printf("Total lines: %d, reportBitSum: %v, gammaRate: %v, gammaRateDecimal: %d, epsilonRate: %v, epsilonRateDecimal: %d, power consumption: %d\n",
		totalCount,
		reportBitSum,
		gammaRateBinary,
		gammaRate,
		epsilonRateBinary,
		gammaRate,
		gammaRate*epsilonRate,
	)

	oxygenRateBinary := filterByBit(reportBit, length, true)
	co2RateBinary := filterByBit(reportBit, length, false)

	oxygenRate := calculateDecimalRate(oxygenRateBinary)
	co2Rate := calculateDecimalRate(co2RateBinary)

	fmt.Printf("Oxygen generator rating: %v, decimal: %d, CO2 scrubber rating: %v, decimal: %d, life support rating: %d\n", oxygenRateBinary, oxygenRate, co2RateBinary, co2Rate, oxygenRate*co2Rate)
}
