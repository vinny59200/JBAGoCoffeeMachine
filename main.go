package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// write your code here
	//blabla()
	//caluclateCupQuantities()
	//inferFromQuantities()
	var waterMl *int
	waterMl = new(int)
	*waterMl = 400
	var waterMlFilled *int
	waterMlFilled = new(int)
	*waterMlFilled = 0
	var isWaterMlPrompted *bool
	isWaterMlPrompted = new(bool)
	*isWaterMlPrompted = false
	var isWaterMlFilled *bool
	isWaterMlFilled = new(bool)
	*isWaterMlFilled = false

	var milkMl *int
	milkMl = new(int)
	*milkMl = 540
	var milkMlFilled *int
	milkMlFilled = new(int)
	*milkMlFilled = 0
	var isMilkMlPrompted *bool
	isMilkMlPrompted = new(bool)
	*isMilkMlPrompted = false
	var isMilkMlFilled *bool
	isMilkMlFilled = new(bool)
	*isMilkMlFilled = false

	var beansGr *int
	beansGr = new(int)
	*beansGr = 120
	var beansGrFilled *int
	beansGrFilled = new(int)
	*beansGrFilled = 0
	var isBeansGrPrompted *bool
	isBeansGrPrompted = new(bool)
	*isBeansGrPrompted = false
	var isBeansGrFilled *bool
	isBeansGrFilled = new(bool)
	*isBeansGrFilled = false

	var cupsCount *int
	cupsCount = new(int)
	*cupsCount = 9
	var cupsCountFilled *int
	cupsCountFilled = new(int)
	*cupsCountFilled = 0
	var isCupsCountPrompted *bool
	isCupsCountPrompted = new(bool)
	*isCupsCountPrompted = false
	var isCupsCountFilled *bool
	isCupsCountFilled = new(bool)
	*isCupsCountFilled = false

	var money *int
	money = new(int)
	*money = 550

	var previousCommand *string
	previousCommand = new(string)
	*previousCommand = ""

	//display(*waterMl, *milkMl, *beansGr, *cupsCount, *money)

	fmt.Printf("Write action (buy, fill, take, remaining, exit):\n")
loop:
	for {
		var menu string
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if menu == "back" {
			fmt.Printf("Write action (buy, fill, take, remaining, exit):\n")
			continue loop
		}
		/*
		   ______
		   | ___ \
		   | |_/ /_   _ _   _
		   | ___ \ | | | | | |
		   | |_/ / |_| | |_| |
		   \____/ \__,_|\__, |
		                 __/ |
		                |___/
		*/
		if menu == "buy" {
			fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
			*previousCommand = "buy"
			continue loop
		}
		if menu == "1" && *previousCommand == "buy" {
			fetch(waterMl, 250, milkMl, 0, beansGr, 16, money, 4, cupsCount)
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			continue loop
		} else if menu == "2" && *previousCommand == "buy" {
			fetch(waterMl, 350, milkMl, 75, beansGr, 20, money, 7, cupsCount)
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			continue loop
		} else if menu == "3" && *previousCommand == "buy" {
			fetch(waterMl, 200, milkMl, 100, beansGr, 12, money, 6, cupsCount)
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			continue loop
		}

		/*
		   ______ _ _ _
		   |  ___(_) | |
		   | |_   _| | |
		   |  _| | | | |
		   | |   | | | |
		   \_|   |_|_|_|
		*/
		if menu == "fill" {
			*previousCommand = "fill"

		fill:
			for {
				var fillCommand = ""
				if *isWaterMlFilled && *isMilkMlFilled && *isBeansGrFilled && *isCupsCountFilled {
					fill(waterMl, *waterMlFilled, milkMl, *milkMlFilled, beansGr, *beansGrFilled, cupsCount, *cupsCountFilled)
					*isWaterMlPrompted = false
					*isWaterMlFilled = false
					*isMilkMlPrompted = false
					*isMilkMlFilled = false
					*isBeansGrPrompted = false
					*isBeansGrFilled = false
					*isCupsCountPrompted = false
					*isCupsCountFilled = false
					*waterMlFilled = 0
					*milkMlFilled = 0
					*beansGrFilled = 0
					*cupsCountFilled = 0
					continue loop
				}
				if "fill" == *previousCommand && !*isWaterMlPrompted {
					fmt.Println("Write how many ml of water you want to add:")
					_, err := fmt.Scanln(&fillCommand)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return
					}
					*waterMlFilled, _ = strconv.Atoi(fillCommand)
					*isWaterMlPrompted = true
					*isWaterMlFilled = true
					continue fill
				} else if "fill" == *previousCommand && !*isMilkMlPrompted {
					fmt.Println("Write how many ml of milk you want to add:")
					_, err = fmt.Scanln(&fillCommand)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return
					}
					*milkMlFilled, _ = strconv.Atoi(fillCommand)
					*isMilkMlPrompted = true
					*isMilkMlFilled = true
					continue fill
				} else if "fill" == *previousCommand && !*isBeansGrPrompted {
					fmt.Println("Write how many grams of coffee beans you want to add:")
					_, err = fmt.Scanln(&fillCommand)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return
					}
					*beansGrFilled, _ = strconv.Atoi(fillCommand)
					*isBeansGrPrompted = true
					*isBeansGrFilled = true
					continue fill
				} else if "fill" == *previousCommand && !*isCupsCountPrompted {
					fmt.Println("Write how many disposable cups of coffee you want to add:")
					_, err = fmt.Scanln(&fillCommand)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return
					}
					*cupsCountFilled, _ = strconv.Atoi(fillCommand)
					*isCupsCountPrompted = true
					*isCupsCountFilled = true
					continue fill
				}
			}

		}

		/*
		    _____ _   _
		   |  _  | | | |
		   | | | | |_| |__   ___ _ __
		   | | | | __| '_ \ / _ \ '__|
		   \ \_/ / |_| | | |  __/ |
		    \___/ \__|_| |_|\___|_|
		*/
		if menu == "take" {
			fmt.Println("I gave you $" + string(*money))
			*money = 0
			*previousCommand = "take"
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			continue loop
		} else if menu == "remaining" {
			display(*waterMl, *milkMl, *beansGr, *cupsCount, *money)
			*previousCommand = "remaining"
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			continue loop
		} else if menu == "exit" {
			break loop
		}

	}

}

func display(waterMl int, milkMl int, beansGr int, cupsCount int, money int) (int, error) {
	return fmt.Printf(`The coffee machine has:
%[1]d of water
%[2]d of milk
%[3]d of coffee beans
%[4]d of disposable cups
%[5]d of money`, waterMl, milkMl, beansGr, cupsCount, money)
}

func fill(mainWater *int, water int, mainMilk *int, milk int, mainCoffee *int, coffee int, mainCups *int, cups int) {
	*mainWater += water
	*mainMilk += milk
	*mainCoffee += coffee
	*mainCups += cups
}

func fetch(mainWater *int, water int, mainMilk *int, milk int, mainCoffee *int, coffee int, mainMoney *int, money int, mainCups *int) {
	if *mainWater < water {
		fmt.Println("Sorry, not enough water!")
	} else if *mainMilk < milk {
		fmt.Println("Sorry, not enough milk!")
	} else if *mainCoffee < coffee {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if *mainCups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		*mainWater -= water
		*mainMilk -= milk
		*mainCoffee -= coffee
		*mainCups -= 1
		*mainMoney += money
	}
}

func inferFromQuantities() {
	fmt.Println("Write how many ml of water the coffee machine has:")
	var waterMl int
	_, err := fmt.Scanln(&waterMl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Write how many ml of milk the coffee machine has:")
	var milkMl int
	_, err = fmt.Scanln(&milkMl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	var beansGr int
	_, err = fmt.Scanln(&beansGr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Write how many cups of coffee you will need:")
	var cupsCount int
	_, err = fmt.Scanln(&cupsCount)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	var maxWaterUnit = waterMl / 200
	var maxMilkUnit = milkMl / 50
	var maxBeansUnit = beansGr / 15
	s := [3]int{maxWaterUnit, maxMilkUnit, maxBeansUnit}
	minUnits := s[0]
	for i := 1; i < len(s); i++ {
		if minUnits > s[i] {
			minUnits = s[i]
		}
	}
	//fmt.Println(s)
	//fmt.Println(minUnits)
	if cupsCount == minUnits {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if cupsCount > minUnits {
		fmt.Printf("No, I can make only %d cups of coffe\n", minUnits)
	} else {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", minUnits-cupsCount)
	}
}

func caluclateCupQuantities() {
	fmt.Println("Write how many cups of coffee you will need:")
	var cupsCount int
	_, err := fmt.Scanln(&cupsCount)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("For %d cups of coffee you will need:", cupsCount)
	fmt.Printf(`
%[1]d ml of water
%[2]d ml of milk
%[3]d g of coffee beans
`, cupsCount*200, cupsCount*50, cupsCount*15)
}

func blabla() (int, error) {
	return fmt.Println(`
Starting to make a coffee
Grinding coffee beans
Boiling water
Mixing boiled water with crushed coffee beans
Pouring coffee into the cup
Pouring some milk into the cup
Coffee is ready!\n
`)
}
