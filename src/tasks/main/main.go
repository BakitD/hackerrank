package main

import (
	"flag"
	"fmt"
	"tasks"
)

func main() {
	taskPtr := flag.String("task", "", "Specify task to run")
	flag.Parse()

	switch *taskPtr {
	case "warmup":
		fmt.Println(tasks.SimpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
	}
}
