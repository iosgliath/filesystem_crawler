package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"path/filepath"
	"regexp"
	"strings"
)


func plunder(p string, i os.FileInfo, e error) error {
	tmp := strings.Split(p, "/")
	here := tmp[len(tmp)-1]
	for x, n := range lkps {
		if n.MatchString(here) {
			fmt.Printf("[+] Hit\n")
			fmt.Printf("	.name 	%v\n", i.Name())
			fmt.Printf("	.perm 	%v\n", i.Mode())
			fmt.Printf("	.@	%s\n", p)
			fmt.Printf("	.for 	%v\n", lkps[x])
			fmt.Printf("	.size	%v\n", i.Size())
			fmt.Printf("	.mod 	%v\n", i.ModTime())
			if i.IsDir() {
				fmt.Printf("	.type	directory\n")
			}
			if filepath.Ext(p) != "" {
				fmt.Printf("	.type	%s\n", filepath.Ext(p))
			}
		}
	}
	return nil
}

func lookups(path string) []*regexp.Regexp {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lkp := []*regexp.Regexp{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lkp = append( lkp, regexp.MustCompile( scanner.Text() ) )
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lkp
}

var lkps = lookups(os.Args[2])

func main() {
	fmt.Printf("\n\nParsing :	%s\n", os.Args[1])
	fmt.Printf("For keys :	%s\n\n\n", lkps)
	if e := filepath.Walk(os.Args[1], plunder); e != nil {
		log.Panicln(e)
	}
	fmt.Printf("\n\n\n")
}
