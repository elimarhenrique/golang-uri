package main

import "fmt"

func main() {
	var p1, p2, p3 string

	fmt.Scanln(&p1)
	fmt.Scanln(&p2)
	fmt.Scanln(&p3)

	if p1 == "vertebrado" {
		if p2 == "ave" {
			if p3 == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if p3 == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if p2 == "inseto" {
			if p3 == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if p3 == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
