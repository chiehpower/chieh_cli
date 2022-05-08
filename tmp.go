package main

import (
	"fmt"
	"os/exec"
	"log"

)

func main() {
	// cmd := exec.Command("ls", "-lah")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// fmt.Printf("combined out:\n%s\n", string(out))

	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println("Start install it")

		// log.Fatal(err)
	}
	
	log.Println(path) // bin/ls
}

// as util
// func commandExists(cmd string) bool {
// 	_, err := exec.LookPath(cmd)
// 	return err == nil
// }

