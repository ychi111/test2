package main

import "fmt"
import "math/rand"
import "time"


func main(){


	var a int = 0
	var g string
	var n,f,r int 
	
	fmt.Println("Учимся считать!!!")
	fmt.Print("Выбери способы оперции: ")		
	fmt.Scanln(&g)	
	for i := 0; a != 10; i++ {

			rand.Seed(time.Now().UnixNano())
			number1 := rand.Intn(490)

			rand.Seed(time.Now().UnixNano())
			number2 := rand.Intn(423)

		a = a + 1

// 1 - сложение

		if g == "+" {

		fmt.Println(number1,"+",number2)
		fmt.Scanln(&f)	
		n = number1 + number2
		fmt.Println("Вашь ответ:",f,"правильниый ответ:",n)

		

					} 
	
// 2 - вычитание

		if g == "-" {

		fmt.Println(number1,"-",number2)
		fmt.Scanln(&f)	
		n = number1 - number2
		fmt.Println("Вашь ответ:",f,"правильниый ответ:",n)

		

					} 	

// 3 - умножение

		if g == "*" {

		fmt.Println(number1,"*",number2)
		fmt.Scanln(&f)	
		n = number1 * number2
		fmt.Println("Вашь ответ:",f,"правильниый ответ:",n)

		

					} 	

// 4 - деление

		if g == "/" {

		fmt.Println(number1,"/",number2)
		fmt.Scanln(&f)	
		n = number1 / number2
		fmt.Println("Вашь ответ:",f,"правильниый ответ:",n)
		
					} 	

					if f == n {

			r = r + 1

		} else {

			a = a - 1	

		}
}


	fmt.Println()
	fmt.Println("верно:",r,"из:",a)



}