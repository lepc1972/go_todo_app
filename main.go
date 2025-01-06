package main

import "fmt"

func main() {
	fullGolang := "watch TWN golang full course"
	shortGolang := "watch our go crash course"
	rewardCoffee := "Reward myself with a cup of coffee ☕"

	taskItems := []string{fullGolang, shortGolang, rewardCoffee} // slice
	//var taskItems = [20]string{fullGolang, shortGolang, rewardCoffee} array
	fmt.Println("##### Welcome to our todolist app #####")

	fmt.Println("List of my todos")

	for i, task := range taskItems {
		fmt.Println(i+1, task)

	}

}
