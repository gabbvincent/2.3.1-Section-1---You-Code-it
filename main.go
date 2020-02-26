// Programmer name: Vincent Gabb
// Date completed:  2/25/2020
// Description: 2.3.1 Section 1

package main

import (
    "fmt"
    "math/rand"
    
    
) //adding the ability to do random numbers

func main() {
    //create two variables - one for the computer and one for the user
    
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 0=rock, 1=scissors, 2=paper

    var user int 
    var computer = rand.Intn(3)

    //prompt the user for an integer value representing the player's choice

    fmt.Println("1 = scissors 2 = Rock 3 = Paper")

    fmt.Println()

    fmt.Println("Please enter a 1, 2, or 3.")
    fmt.Scanln(&user)
    
    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    //You will need to use decisions for this

    if user == 1 {
      fmt.Println("You chose Scissors.")
    } else if user == 2 {
      fmt.Println("You chose Rock.")
    } else {
      fmt.Println("You chose Paper.")
    }

    if computer == 0 {
      fmt.Println("The Computer chose Rock.")
    } else if computer == 1 {
      fmt.Println("The Computer chose Scissors.")
    } else {
      fmt.Println("The Computer chose Paper.")
    }
    
}