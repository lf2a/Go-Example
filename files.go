package main

import (
	"log"
	"os"
    "fmt"
    "time"
)

func createEmptyFile() {
    emptyFile, err := os.Create("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

func createDir() {
    _, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("new_dir", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

func renameFile() {
    oldName := "file1.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func moveFile() {
    oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile() {
    sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

func getFileInfo() {
    fileStat, err := os.Stat("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}

func deleteFile() {
    err := os.Remove("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func readChar() {
    filename := "test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
}

func appendFile() {
    message := "Add this content at end"
	filename := "test.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}

// Changing permissions, ownership, and timestamps
func change() {
    // Test File existence.
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// Change permissions Linux.
	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change file timestamps.
	addOneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

func main() {
    // createEmptyFile()
    // createDir()
    // renameFile()
    // moveFile()
    // copyFile()
    // getFileInfo()
    // deleteFile()
    // readChar()
    // appendFile()
    // change()
}
