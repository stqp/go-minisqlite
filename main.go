package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/xerrors"
)

func printPrompt() {
	fmt.Print("sqlite> ")
}

func dbOpen(filename string) (fp *os.File) {
	fp, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func l(msg interface{}) {
	fmt.Println(msg)
}

func doMetaCommand(metaCommand string) error {
	if metaCommand == ".help" {
		l("help")
		return nil
	} else if metaCommand == ".exit" {
		os.Exit(0)
	}
	return xerrors.Errorf("Error: unknown command or invalid arguments:  \"%s.\" Enter \".help\" for help: %w", metaCommand)

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Must supply a database filename.\n")
		os.Exit(1)
	}

	//filename := os.Args[1]
	//file := db_open(filename)

	scan := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		scan.Scan()
		command := scan.Text()

		err := doMetaCommand(command)
		if err != nil {
			l(err)
		}
	}

}
