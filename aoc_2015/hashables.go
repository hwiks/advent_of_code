package aoc_2015

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func begins_with_00000(hash string) bool {
	return hash[:6] == "000000"
}

func Run_hashables() {

	var input string = "bgvyzdsv"
	var found_hash bool = false
	var i int = 0

	fmt.Println("running")

	for !found_hash {
		var output string = GetMD5Hash(input + strconv.Itoa(i))

		if begins_with_00000(output) {
			found_hash = true
			fmt.Println("Found a hash beginning with 000000: ")
			fmt.Println(output)
			fmt.Println("with number: ")
			fmt.Println(i)

		}
		i = i + 1

	}

}
