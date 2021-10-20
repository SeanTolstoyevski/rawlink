// Copyright (C) 2021 SeanTolstoyevski - mailto:seantolstoyevski@protonmail.com
//
// The source code of this project is licensed under the MIT license.
// You can find the license on the repo's main folder.
// Provided without warranty of any kind.


// Read the README for more information.
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
	inFile  = flag.String("i", "backlink.txt", "Location of the file to be read")
	outFile = flag.String("o", "fixed_backlinks.txt", "Name of the file to be written")
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
