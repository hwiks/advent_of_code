package aoc_2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Run_circuits() {

	var circuits_commands []string = extract_lines("inputs/circuits_test.txt")

	var circuits = make(map[string]string, 0)

	for i := range circuits_commands {

		var circuits_components []string = strings.Split(circuits_commands[i], "->")

		circuits[circuits_components[1]] = circuits_components[0][:len(circuits_components[0])-1]

	}

	fmt.Println(circuits)
	// simplify all simplifiable expressions, then solve for desired variables
	// this would be my human approach but there is def better ways out there
	for variable, circuit := range circuits {
		if num_value, err := strconv.Atoi(circuit); err == nil {
			fmt.Println("Found an int ", variable, num_value)

			//todo: only save the ints
			//      then run as a function smtg like below for each of the ints (might need some trimming)
			// this problem needs to be recursive somehow
			// I also need to look into the gates: there must be libraries for the gates commands

			for var_nest, circuit_nest := range circuits {
				if strings.Contains(circuit_nest, variable) {
					fmt.Println("found the var here ", circuit_nest)
					var new_circuit = strings.Replace(circuit_nest, variable, " "+strconv.Itoa(num_value), -1)

					circuits[var_nest] = new_circuit

				}

			}
			delete(circuits, variable)
			fmt.Println(circuits)
		}

	}

}
