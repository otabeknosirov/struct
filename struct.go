package main

import "fmt"

type ContactMan struct{

	name string
	surname string
	gender string
}

type ContactWoman struct{

	name string
	surname string
	gender string
}

func main()  {
	var man ContactMan
	var woman ContactWoman

	/*Entering man's data*/
	man.name = "Bill"
	man.surname = "Gates"
	man.gender = "male"

	/*Entering woman's data*/
	woman.name = "Janna"
	woman.surname = "Dark"
	woman.gender = "female"

	var gender string

	fmt.Println("Enter the gender:")
	fmt.Scan(&gender)

	if(gender == man.gender){
		fmt.Printf("%s \t%s \t%s", man.name, man.surname, man.gender)
	}else{
		fmt.Printf("%s \t%s \t%s", woman.name, woman.surname, woman.gender)
		
	}
	fmt.Println()


}