package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"regexp"
//	"strings"
//)
//
//func main() {
//	fmt.Println(dayFourA(parseTxt()))
//}
//
//func dayFourA(input []string) int {
//	count := 0
//
//	for _, passport := range input {
//		if strings.Contains(passport, "byr:") &&
//			strings.Contains(passport, "iyr:") &&
//			strings.Contains(passport, "eyr:") &&
//			strings.Contains(passport, "hgt:") &&
//			strings.Contains(passport, "hcl:") &&
//			strings.Contains(passport, "ecl:") &&
//			strings.Contains(passport, "pid:") {
//			count++
//		}
//	}
//
//	return count
//}
//
//func dayFourB(input []string) int {
//	count := 0
//
//	for _, passport := range input {
//		re := regexp.MustCompile("byr:([\\d]{4})")
//
//		if match, _ := regexp.MatchString(`byr:([\d]{4})`, passport); match {
//
//		}
//	}
//
//	return count
//}
//
//func parseTxt() []string {
//	rows, err := ioutil.ReadFile("./input.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return strings.Split(string(rows), "\n\n")
//}