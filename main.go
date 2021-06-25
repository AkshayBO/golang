package main

import "fmt"

// function to add number
func add (x int,y int) int{
        var result int= x+y
        return result
}

// function to return two values
func add1 (x int,y int) (int,int){
        return x,y
}

func main(){
        fmt.Println("hello world")
        //type of variable declaration
        //var num1 = 2
        var num2 int = 3
        // if you want to exclude var keyword
        num1 := 2
        // declare first and then initialized
        //var num11 int
        //num11 := 11
        var result= num1+num2
        fmt.Println(result)

        //There is only one loop in golang i.e for loop
        i:=1
        for i<=5{
                fmt.Println("Akshay")
                i++
        }
        //same way to define for loop

        for i := 1; i<=5; i++ {
                fmt.Println("Akshay")
        }
        num3 := 3
        num4 := 6
        result1 := add(num3,num4)
        println(result1)
        // in golang you can return two values from function
        // function to return two values
        result2,result3 := add1(num3,num4)
        println(result2,result3)
        //if and else statement in golang
        test :=1
        if test <= 5 {
                println("Hi"}
        } else {
                println("Byeeee!")
        }
}
