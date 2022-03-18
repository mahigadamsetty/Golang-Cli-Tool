package main

import (
	"flag"
	"log"
)

func main() {
	var favMovie string
	var wishMovie string

	flag.StringVar(&favMovie, "f", "", "favourite movies to be added")
	flag.StringVar(&wishMovie, "w", "", "wishlist movies to be added")

	wantWishList := flag.Bool("lw", false, "want wish list")
	flag.Parse()

	if wishMovie != "" {
		if err := WriteToFile("wish_list.txt", wishMovie); err != nil {
			log.Fatal(err)
		}
	}
	if favMovie != "" {
		if err := WriteToFile("favourite_list.txt", favMovie); err != nil {
			log.Fatal(err)
		}

	}

	if *wantWishList {
		if err := ReadFromFile("wish_list.txt"); err != nil {
			log.Fatal(err)
		}
	}
}
