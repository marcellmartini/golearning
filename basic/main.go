package main

import (
	"basic/array"
	"basic/concurrency"
	"basic/concurrency/channel"
	"basic/errorhandling"
	"basic/function"
	"basic/json"
	"basic/maps"
	"basic/pointer"
	"basic/slice"
	"basic/stdout"
	"basic/structe"
	"basic/testing"
)

func main() {
	array.Examples()
	channel.Examples()
	concurrency.Examples()
	errorhandling.Examples()
	function.Examples()
	json.Examples()
	maps.Examples()
	pointer.Examples()
	testing.Examples()
	slice.Examples()
	stdout.Examples()
	structe.Examples()
}
