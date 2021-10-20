package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	inFile  = flag.String("i", "backlink.txt", "Okunacak dosyanın konumu")
	outFile = flag.String("o", "fixed_backlinks.txt", "Yazılacak dosyanın konumu")
)

func main() {
	flag.Parse()
	bLinks, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer bLinks.Close()

	fixedLinks, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fixedLinks.Close()

	scanner := bufio.NewScanner(bLinks)
	for scanner.Scan() {
		slashIndex := strings.Index(scanner.Text(), "/")
		if slashIndex == -1 {
			continue
		}

		fixText := scanner.Text()[:slashIndex] + "\r\n"
		_, err := fixedLinks.Write([]byte(fixText))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")

}
