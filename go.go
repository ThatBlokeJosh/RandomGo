package main

import (
	"fmt"
	"time"
	"strconv"
)

func random() int32 {
	time := time.Now();
	var secondsStr string = time.Format(".000");
	seconds, err := strconv.ParseFloat(secondsStr, 32);
	if err!= nil {
        panic(err);
    }

	randomNum := int32(seconds * 100)
	return randomNum;

}

func main() {
	// slice of the alphabet
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	randomNum := random();

	var i int32 = 0;
	for i < randomNum {
		if randomNum >= 26 {
			randomNum -= 26;
		}
		i++;
	}

    fmt.Println(alphabet[randomNum])


}
