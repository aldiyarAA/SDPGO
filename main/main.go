package main

import "fmt"

func main() {
	start()
}

var sum = 20

func start() {
	var variant int

	var username string
	fmt.Println("Wellcome to the game\nEnter Your name:")
	fmt.Scanf("%s", &username)

	fmt.Println("Hello,", username, "!\n"+
		"Now you can start your mission of resque the world.\n"+
		"You must won the villain \"Covid\" that want to destroy the world.\n"+
		"Whole world in your hand!\n"+
		"You have a Robot with the default configurations that have a transform mode. You"+
		"\nneed to use this function, because Covid has pa(power amount ) - 60. And your "+
		"\nrobot has only 20pa. To achive more than covid you need to pass the quest."+
		"\n"+
		"1.Start game.\n"+
		"2.I could't save the world. I have no mask.")
	fmt.Scan(&variant)
	if variant == 1 {
		bridge()
	} else {
		fmt.Println("it is so pity")
		return
	}

}
func result() {
	if sum > 60 {
		fmt.Println("Congrats, You won the Covid."+
			"\n You scored", sum)
	} else {
		fmt.Println("The world destroyed By Covid."+
			"\n Your score was ", sum, "that is less than Covid's pa")
	}

}
