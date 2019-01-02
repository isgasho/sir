package main

import (
	"github.com/komuw/dbscan/heart"
)

/*
usage:
  go run -race heart/cmd/main.go
*/

func main() {
	heart.Run(
		6, 2,
		[]float64{
			1, 2,
			2, 2,
			2, 3,
			8, 7,
			8, 8,
			25, 80},
		3.0,
		1.0,
		false,
		"testHeartRun")

	// heart.Run(
	// 	10,
	// 	4,
	// 	[]float64{
	// 		5.1, 3.5, 1.4, 0.2,
	// 		4.9, 3.0, 1.4, 0.2,
	// 		4.7, 3.2, 1.3, 0.2,
	// 		4.6, 3.1, 1.5, 0.2,
	// 		5.0, 3.6, 1.4, 0.2,
	// 		5.4, 3.9, 1.7, 0.4,
	// 		4.6, 3.4, 1.4, 0.3,
	// 		5.0, 3.4, 1.5, 0.2,
	// 		4.4, 2.9, 1.4, 0.2,
	// 		4.9, 3.1, 1.5, 0.1,
	// 	},
	// 	3.0,
	// 	1.0,
	// 	false,
	// 	"testHeartRun")

	// heart.Run(7, 160,
	// 	[]float64{80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 107, 111, 109, 117, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 107, 111, 109, 117, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 106, 117, 109, 97, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 106, 117, 109, 97, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 106, 117, 109, 97, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 106, 111, 104, 110, 34, 125, 80, 79, 83, 84, 32, 47, 112, 111, 115, 116, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 32, 104, 116, 116, 112, 98, 105, 110, 46, 111, 114, 103, 13, 10, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55, 46, 53, 52, 46, 48, 13, 10, 97, 99, 99, 101, 112, 116, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101, 58, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58, 32, 49, 53, 13, 10, 13, 10, 123, 34, 110, 97, 109, 101, 34, 58, 34, 106, 111, 104, 110, 34, 125},
	// 	3.0, 1.0, false, "testHeartRun")

	/*
			or if you want Data to be autogenerated for you
		heart.Run(6, 2, []float64{1, 2, 2, 2, 2, 3, 8, 7, 8, 8, 25, 80}, 3.0, 1.0, true)
	*/
}

/*
debug:
	go build -o heat heart/cmd/main.go
	dlv exec ./heat
*/
