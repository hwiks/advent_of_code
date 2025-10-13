package aoc_2015

import (
	"fmt"
	"strconv"
)

func Run_sleight() {

	var sleight_list []string = extract_lines("inputs/sleight.txt")

	var str_lit_len int = 0
	var total_n_characters int = 0
	var total_with_extra int = 0

	for _, present := range sleight_list {

		str_lit_len += len(present)

		var present_string, err = strconv.Unquote(present)
		check_error(err)

		var present_quoted = strconv.Quote(present)

		total_n_characters += len(present_string)

		total_with_extra += len(present_quoted)

	}

	fmt.Println("For the first riddle: ", str_lit_len-total_n_characters)
	fmt.Println("For the second riddle: ", total_with_extra-str_lit_len)

}
