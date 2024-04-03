package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Print("Enter password: ")
	godotenv.Load()

	var input string = ""
	fmt.Scanln(&input)
	if input != os.Getenv("PASSWD") {
		fmt.Println(fmt.Errorf("invalid password"))
		fmt.Println()
		return
	}

	cmd := exec.Command("powershell")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	os.Stdout.Write([]byte("! PIPED TEMRINAL - EXECUTE EXIT TO QUIT !\n"))

	cmd.Run()
}
