package model

import "fmt"

func Help() {
	logo()
	fmt.Println("Usage of gofs:")
	fmt.Println("  gofs help\tget usage of gofs")
	fmt.Println("  gofs dninfo\tget datanode info")
	fmt.Println("")
	fmt.Println("  gofs put [gofspath] [inputpath]\tput file to gofs")
	fmt.Println("  gofs get [gofspath] [outputpath]\tget file to local")
}

func logo() {
	fmt.Println("  _____  ____  ______ _____ ")
	fmt.Println(" / ____|/ __ \\|  ____/ ____|")
	fmt.Println("| |  __| |  | | |__ | (___  ")
	fmt.Println("| | |_ | |  | |  __| \\___ \\ ")
	fmt.Println("| |__| | |__| | |    ____) |")
	fmt.Println(" \\_____|\\____/|_|   |_____/ ")
	fmt.Println("")
}
