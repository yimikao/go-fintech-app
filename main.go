package main

import "go-fintech-app/migrations"

func appendInt(curSlice []int, ele int) []int {

	var newSlice []int
	newLen := len(curSlice) + 1
	if newLen <= cap(curSlice) {
		newSlice = curSlice[:len(curSlice)+1]
	} else {

	}
	return newSlice[]
}

func main() {
	migrations.Migrate()

}
