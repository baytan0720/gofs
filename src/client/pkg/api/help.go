package api

import "fmt"

func Help() {
	logo()
	fmt.Println("Usage of gofs <command>:")
	fmt.Println("  gofs help\t'get usage of gofs'")
	fmt.Println("  gofs sysinfo\t'get system stat'")
	fmt.Println("")
	fmt.Println("  gofs ls <gofspath>\t'get file list on path'")
	fmt.Println("  gofs del <gofspath>\t'delete file on gofs'")
	fmt.Println("  gofs stat <gofspath>\t'get file stat'")
	fmt.Println("  gofs rename <gofspath>\t'rename file or directory'")
	fmt.Println("  gofs put <gofspath> <localpath>\t'put file to gofs'")
	fmt.Println("  gofs get <gofspath> <localpath>\t'get file to local'")
	fmt.Println("  gofs mkdir <gofspath> <dirname>\t'make directory on path'")
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
