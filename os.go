package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func chmod() {
	// chmod(filename, modo)
	if err := os.Chmod("filename", 0644); err != nil {
		log.Fatal(err)
	}
}

func env() {
	os.Setenv("NAME", "luiz")
	os.Setenv("AGE", "19")

	fmt.Printf("%s %s.\n", os.Getenv("NAME"), os.Getenv("AGE"))
}

func command() {
	cmd := exec.Command("ls", "-lha")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}

func output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func run() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func start() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func main() {
	// chmod()
	// env()
	// command()
	// output()
	// run()
	// start()

	fmt.Println(os.Getwd())
	fmt.Println(os.Getpid())
	fmt.Println(os.Hostname())
	fmt.Println(os.UserHomeDir())

	usr, _ := user.Current()
	fmt.Println(usr.Gid, usr.HomeDir, usr.Name, usr.Username, usr.Uid)
}
