package utils

import (
	"bufio"
	"log"
	"os"
)

// LoadPassword() takes a file name as an argument, opens the file, reads the file line by line, and
// returns a slice of strings containing the passwords
func LoadPassword(filaName string) ([]string, error) {
	passwords := []string{}
	readFile, err := os.Open(filaName); if err != nil {
        log.Fatal(err)
    }
    defer readFile.Close()


    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
		password := fileScanner.Text()
		passwords = append(passwords, password)
    }
	return passwords, nil
}