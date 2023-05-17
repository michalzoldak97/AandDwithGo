package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var order *product

func listCommands() {
	fmt.Println("commands -> lists all available commands\ncategories -> lists all categories\ncategory {{category}} -> lists all products in category\norder {{amount}} x {{product}} -> create order\npay {{cash amount}} -> pay for order\nq -> exit shop")
}

func payForOrder(s *shop, cash float64) {
	if order == nil {
		fmt.Println("Thanks for your payment, but what are you paying for?")
		return
	}

	if cash < order.price {
		fmt.Println("Wut?? Gimme money")
		return
	}

	success := s.buyProduct(order.name, order.amount)
	if !success {
		log.Fatal("Someting is not yes, failure")
	}

	fmt.Printf("Heres your %v x %v and the change is %v\n", order.amount, order.name,
		fmt.Sprintf("%.2f", cash-order.price))

	order = nil
}

func isPayQuery(phrase string) bool {
	f := strings.Fields(phrase)

	if len(f) != 2 ||
		f[0] != "pay" ||
		f[1] == "" {
		return false
	}

	return true
}

func createOrder(s *shop, p string, a int) {

	if a < 1 {
		fmt.Println("How do you imagine buying negative or 0 number of products? xD")
		return
	}

	exists, prod := s.getProduct(p)
	if !exists {
		fmt.Printf("Product %v not found\n", p)
		return
	}

	if a > prod.amount {
		fmt.Println("Hold up, we don't have that much")
		return
	}

	order = &product{
		name:   prod.name,
		amount: a,
		price:  prod.price * float64(a),
	}

	fmt.Printf("Thats %v\n", order.price)
}

func isOrderQuery(phrase string) bool {
	f := strings.Fields(phrase)

	if f[0] != "order" ||
		f[1] == "" ||
		f[2] != "x" ||
		f[3] == "" {
		return false
	}

	return true
}

func listProducts(s *shop, section string) {
	_, ok := s.products[section]
	if !ok {
		fmt.Printf("We have no %v section in our shop\n", section)
		return
	}

	for _, p := range s.products[section] {
		fmt.Printf(`"%v" amount: %v price: %v`, p.name, p.amount, p.price)
		fmt.Println()
	}
}

func isSectionQuery(phrase string) bool {
	f := strings.Fields(phrase)

	if len(f) != 2 ||
		f[0] != "category" ||
		f[1] == "" {
		return false
	}

	return true
}

func listSections(s *shop) {
	for _, sec := range s.getSections() {
		fmt.Println(sec)
	}
}

func runShop(s *shop, c string) {
	switch {
	case c == "categories":
		listSections(s)
	case isSectionQuery(c):
		listProducts(s, strings.Fields(c)[1])
	case isOrderQuery(c):
		{
			f := strings.Fields(c)

			a, err := strconv.Atoi(f[1])
			if err != nil {
				fmt.Println("You can't have non integer amount of our goods")
				break
			}

			n := strings.Join(f[3:], " ")

			createOrder(s, n, a)
		}
	case isPayQuery(c):
		{
			f := strings.Fields(c)

			a, err := strconv.ParseFloat(f[1], 64)
			if err != nil {
				fmt.Println("Money is countable, use dot separated, numeric value to count it")
				break
			}

			payForOrder(s, a)
		}
	case c == "commands":
		listCommands()
	case c == "q":
		{
			fmt.Println("Bye bye")
			os.Exit(0)
		}
	default:
		fmt.Println("I don't get it, can you repeat?")
	}
}

func handleInput(s *shop) {
	fmt.Println("Welcome!")
	reader := bufio.NewReader(os.Stdin)
	runtimeName := runtime.GOOS

	for {
		fmt.Print("-> ")

		txt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
		}

		if runtimeName == "windows" {
			txt = strings.Replace(txt, "\r\n", "", -1)
		} else if runtimeName == "linux" {
			txt = strings.Replace(txt, "\n", "", -1)
		}

		runShop(s, txt)
	}
}

func main() {
	s := newShop()

	handleInput(&s)
}
