package aoc_2015

import (
	"fmt"
	"strconv"
	"strings"
)

func initialize_circuit_map(commands []string) map[string]string {
	var circuits = make(map[string]string, 0)

	for i := range commands {

		var circuits_components []string = strings.Split(commands[i], "->")

		var circuit_var string = strings.Replace(circuits_components[1], " ", "", -1)
		var circuit_comm string = strings.Replace(circuits_components[0], " ", "", -1)

		circuits[circuit_var] = circuit_comm

	}
	return circuits
}

func solve_solvable(circuit_map map[string]string) int {

	var solved map[string]string = find_and_execute_solvable(circuit_map)

	fmt.Println("solved:", solved)

	sub_solved(solved, circuit_map)
	fmt.Println("circuit map after substitution: ", circuit_map)

	return len(solved)

}

func find_and_execute_solvable(circuit_map map[string]string) map[string]string {

	var solvable map[string]uint16 = make(map[string]uint16)

	for key, circuit := range circuit_map {
		if num_value, err := strconv.Atoi(circuit); err == nil {
			solvable[key] = uint16(num_value)
		}

		if is_solvable(circuit) {
			fmt.Println("It's solvable!", circuit)
			circuit_map[key] = solve_circuit(circuit)
			fmt.Println(solvable[key])
		}

	}

	return circuit_map

}

func solve_circuit(circuit string) string {

	var result uint16
	if strings.Contains(circuit, "AND") {
		var variables []string = strings.Split(circuit, "AND")
		result1, err1 := strconv.Atoi(variables[0])
		result2, err2 := strconv.Atoi(variables[1])

		result = uint16(result1) & uint16(result2)
		fmt.Println("and result", result)

		check_error(err1)
		check_error(err2)

	} else if strings.Contains(circuit, "OR") {
		var variables []string = strings.Split(circuit, "OR")
		result1, err1 := strconv.Atoi(variables[0])
		result2, err2 := strconv.Atoi(variables[1])

		result = uint16(result1) | uint16(result2)

		check_error(err1)
		check_error(err2)

	} else if strings.Contains(circuit, "RSHIFT") {
		var variables []string = strings.Split(circuit, "RSHIFT")
		result1, err1 := strconv.Atoi(variables[0])
		result2, err2 := strconv.Atoi(variables[1])

		result = uint16(result1) >> uint16(result2)

		check_error(err1)
		check_error(err2)

	} else if strings.Contains(circuit, "LSHIFT") {
		var variables []string = strings.Split(circuit, "LSHIFT")
		result1, err1 := strconv.Atoi(variables[0])
		result2, err2 := strconv.Atoi(variables[1])

		result = uint16(result1) >> uint16(result2)

		check_error(err1)
		check_error(err2)

	} else if strings.Contains(circuit, "NOT") {
		var variables []string = strings.Split(circuit, "NOT")
		result1, err1 := strconv.Atoi(variables[0])

		result = ^uint16(result1)

		check_error(err1)

	}

	return strconv.Itoa(int(result))

}

func all_ints(circuit string, command string) bool {

	var all_ints bool

	if strings.Contains(circuit, command) {
		var variables = strings.Split(circuit, command)
		if is_int(variables[0]) && is_int(variables[1]) {
			all_ints = true
		}

	}
	return all_ints
}

func is_solvable(circuit string) bool {

	var excluding_not bool = all_ints(circuit, "AND") || all_ints(circuit, "OR") || all_ints(circuit, "RSHIFT") || all_ints(circuit, "LSHIFT")
	var only_not bool

	if strings.Contains(circuit, "NOT") {
		var variables = strings.Split(circuit, "NOT")
		if is_int(variables[0]) {
			only_not = true
		}

	}

	return excluding_not || only_not
}

func is_int(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	} else {
		return false
	}

}

func sub_solved(solvable map[string]string, circuit_map map[string]string) {

	for solv_key, solv_value := range solvable {
		for key, circuit := range circuit_map {

			if strings.Contains(circuit, solv_key) {

				var new_value = solv_value
				fmt.Println(circuit_map[key])
				fmt.Println(strings.Replace(circuit, solv_key, new_value, -1))

				circuit_map[key] = strings.Replace(circuit, solv_key, new_value, -1)
				fmt.Println(circuit_map[key])

			}

		}
	}

}

func Run_circuits() {

	var circuits_commands []string = extract_lines("inputs/circuits_test.txt")
	var circuit_map map[string]string = initialize_circuit_map(circuits_commands)

	for {
		var new_found int = solve_solvable(circuit_map)
		fmt.Println(new_found)
	}

}
