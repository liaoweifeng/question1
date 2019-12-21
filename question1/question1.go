package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type food struct {
	name  string
	price int
}

var out string

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(8)

	menu1 := food{"小面", 6}
	menu2 := food{"饭团", 7}
	menu3 := food{"香菇滑鸡", 12}
	menu4 := food{"小炒肉", 15}
	menu5 := food{"黄焖鸡", 16}
	menu6 := food{"冒菜", 18}

	if a == 0 {
		sprce := strconv.Itoa(menu1.price)
		out = menu1.name + sprce
	} else if a == 1 {
		sprce := strconv.Itoa(menu2.price)
		out = menu2.name + sprce
	} else if a == 2 {
		sprce := strconv.Itoa(menu3.price)
		out = menu3.name + sprce
	} else if a == 3 {
		sprce := strconv.Itoa(menu4.price)
		out = menu4.name + sprce
	} else if a == 4 {
		sprce := strconv.Itoa(menu5.price)
		out = menu5.name + sprce
	} else if a == 5 {
		sprce := strconv.Itoa(menu6.price)
		out = menu6.name + sprce
	} else if a == 6 {
		sprce := strconv.Itoa(menu1.price + menu2.price)
		out = menu1.name + "和" + menu2.name + sprce
	} else if a == 7 {
		sprce := strconv.Itoa(menu3.price + menu2.price)
		out = menu3.name + "和" + menu2.name + sprce
	} else if a == 8 {
		sprce := strconv.Itoa(menu1.price + menu3.price)
		out = menu1.name + "和" + menu3.name + sprce
	}

	fmt.Println(out)
}
