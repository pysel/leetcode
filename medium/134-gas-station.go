package main

// Idea: brute force. If cannot reach from i to j, cannot reach j from any of (i, j)
// Use modules
// 99% 99%
func canCompleteCircuit(gas []int, cost []int) int {
	station := 0 // a potential index of a starting stations
	index := 0   // an index of current station
	curGas := 0  // current gas amount
	for {
		curGas += gas[index%len(gas)]

		// if we have sufficient amount of gas until the next station,
		// simply increment index and check if we made a loop
		if cost[index%len(gas)] <= curGas {
			curGas -= cost[index%len(gas)]
			index++
			if index%len(gas) == station {
				break
			}
			continue
		} else {
			// start from the next station
			index++

			curGas = 0
			station = index
			if station > len(gas) { // if next potential station is again the initial one, return -1
				return -1
			}
		}
	}

	return station
}
