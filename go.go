package main

import (
	"fmt"
	"strconv"
	"time"
)

func random(time time.Time) int32 {
	var secondsStr string = time.Format(".000");
	seconds, err := strconv.ParseFloat(secondsStr, 32);
	if err!= nil {
        panic(err);
    }

	randomNum := int32(seconds * 10)
	if randomNum == 0 {
		randomNum = 1
    }
	return randomNum;

}

func main() {
	var input string;
	fmt.Print("Enter a string: ");
	fmt.Scan(&input);
	presentTime := time.Now()
	randomNum := random(presentTime)

	for i := 0; i < len(input); i++ {
		fmt.Printf("%c", input[i] + byte(randomNum));
	}
	println()
}
