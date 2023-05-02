package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName="Go conference "
const conferenceTickets int=50
var remainingTickets uint=50

// Slice 
var bookings= make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg= sync.WaitGroup{}

func main()  {
	
	greetUser()

	

	
	
	firstName,lastName,email,userTickets :=getUserInput()


	isValidName,isValidEmail,isValidTicketNumber :=ValidUserInput(firstName,lastName,email,userTickets,remainingTickets )

	if isValidEmail && isValidName && isValidTicketNumber {
	
		bookTicket( userTickets ,firstName ,lastName ,email  )	

		wg.Add(1)
		//sets the number of goroutines to wait for 
		go sendTicket(userTickets ,firstName ,lastName ,email )

	fmt.Printf("First names of booking: %v",getFirstNames())
	



	if  remainingTickets ==0{
		fmt.Println("Our conference is booked out.Please come back next year.")
		// break
	}
		
	}else{
		if !isValidEmail{
			fmt.Printf("Entered email doesn't contains @ sign\n")
		}
		if !isValidName{
			fmt.Printf("First name or last name you enter is too short\n")
		}
		if !isValidTicketNumber{
			fmt.Printf("Number of tickets you enter is invalid\n")
		}
	}
	wg.Wait()

	//BLocks until the WaitGroup Counter is 0
	
	
}

func greetUser(){
	fmt.Printf("Welcome to %v booking application \n",conferenceName)
	fmt.Println("We have total of ",conferenceTickets,"Tickets and",remainingTickets,"are still available")
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string{
	var firstNames []string
	for _,booking :=range bookings{

		
		firstNames = append(firstNames, booking.firstName)
	}
	
	return firstNames
}



func getUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name")
	fmt.Scan(&email)

	fmt.Println("Enter No. of ticket you wants")
	fmt.Scan(&userTickets)

	return firstName,firstName,email,userTickets
}


func bookTicket(userTickets uint,firstName string,lastName string,email string){
	remainingTickets=remainingTickets- userTickets

	//create a map for a user 

	var userData=UserData{

		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets:userTickets ,
	}


	
	bookings = append(bookings,userData )
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for buying %v tickets .You will received confirmation email at %v\n",firstName,lastName,userTickets,email)	

	fmt.Printf("%v tickets are remaining for %v\n",remainingTickets,conferenceName)

}

func sendTicket(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(10*time.Second)
	var ticket=fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n",ticket,email)
	fmt.Println("##################")

	wg.Done()
	//Decrements the waitGroup counter by 1.
}