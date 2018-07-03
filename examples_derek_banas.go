// to run go to the terminal and run: go run program_name.go
package main

import "fmt" //format package

/* hello
func main() {
 fmt.Println("Hello, World")
}
*/

func main() {
 
 var ID int = 121
 var weight float64 = 69.7
 fmt.Println(ID, "has", weight) // by default space added where comma(,) is
}
// len(stringName)
// str3 = str1 + str2
// fmt.Printf("%.3f kg\n", weight) // %T for type, your typical types: %d, %b(for binary), %c, %x(hex), %e(scientific)
// logical operators: &&, ||, !

/*
if condition == true {
    fmt.Println("Cond. true")
} else if {
    fmt.Println("Cond. false")
} else {
    something
}
*/

/*
var favNums2[5] float64

favNums2[0] = 163
favNums2[1] = 78557
favNums2[2] = 691
favNums2[3] = 3.141
favNums2[4] = 1.618

// You access the value by supplying the index number

fmt.Println(favNums2[3])

// Another way of initializing an array

favNums3 := [5]float64 { 1, 2, 3, 4, 5 }

// How to iterate through an array (Use _ if a value isn't used)

for i, value := range favNums3 {

	fmt.Println(value, i)

}
*/

// closure: function acts as repetitive code, which can access all the variables

// defer: execute after enclosing function is completed

// recover: error handling, defer + recover
/*
func safeDiv(num1, num2 int) int {
    defer func() {
        fmt.Println(recover())
    }()
    solution := num1 / num2
    return solution
}
*/

// pointer
/*
changeXValNow(&x)
// * signals that we are being sent a reference to the value
func changeXValNow(x *int){
 
	// Change the value at the memory address referenced by the pointer
	// * gives us access to the value the pointer points at
	
	*x = 2 // Store 2 in the memory address x refers to
 
}
*/


// create structures and assign methods to them
/*
package main
 
import "fmt"

func main() {
     
    // Define a rectangle
    rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}
     
    // Leave off attribute names if we know the order
    // rect1 := Rectangle{0, 50, 10, 10}
     
    // We access values with the dot operator
    fmt.Println("Rectangle is", rect1.width, "wide")
     
    // Call the method area for Rectangle
    fmt.Println("Area of the rectangle =", rect1.area())
     
}
 
// We can define our own types using struct
type Rectangle struct{
	leftX float64
	TopY float64
	height float64
	width float64
}
 
// We can define methods for our Rectangle by adding the receiver
// rect *Rectangle between func and the function name so we can 
// call it with the dot operator
func (rect *Rectangle) area() float64{
 
	return rect.width * rect.height
 
}
*/


// interface (polymorphism)

/*
package main
 
import "fmt"
import "math"
 
// STRUCTS AND INTERFACES
 
func main() {
 
	rect := Rectangle{20, 50}
	circ := Circle{4}
	
	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))
 
}
 
// An interface defines a list of methods that a type must implement
// If that type implements those methods the proper method is executed
// even if the original is referred to with the interface name
 
type Shape interface { // interface for area() function
	area() float64
}
 
type Rectangle struct{
	height float64
	width float64
}
 
type Circle struct{
	radius float64
}
 
func (r Rectangle) area() float64 { // area function attatched to Rectangle structure
	return r.height * r.width
}
 
func (c Circle) area() float64 { // area function attatched to Circle structure
	return math.Pi * math.Pow(c.radius, 2)
} 
 
func getArea(shape Shape) float64{ // call area function for type Shape(look for Shape interface)
 
	return shape.area()
 
}
*/


// strings, file io, exception handling, casting

/*
package main
 
import ("fmt"
"strings"
"sort"
"os"
"log"
"io/ioutil"
"strconv")
 
func main() {
 
	// STRING FUNCTIONS
 
	sampString := "Hello World"
 
	// Returns true if phrase exists in string
	fmt.Println(strings.Contains(sampString, "lo"))
	
	// Returns the index for the match
	fmt.Println(strings.Index(sampString, "lo"))
	
	// Returns the number of matches for the string
	fmt.Println(strings.Count(sampString, "l"))
	
	// Replaces the first letter with the second as many times as you define
	fmt.Println(strings.Replace(sampString, "l", "x", 3))
	
	// Return a list separating with the defined separator 
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)
	
	// Returns a string using the values passed in separated with separator
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ");
	
	fmt.Println(listOfNums);
	
	// FILE I/O
	
	// Create a file
	file, err := os.Create("samp.txt") 
	
	// Output any errors
	if err != nil { 
		log.Fatal(err)
	} 
	
	// Write a string to the file
	file.WriteString("This is some random text")
	
	// Close the file
	file.Close() 
	
	// Try to open the file
	stream, err := ioutil.ReadFile("samp.txt")
	
	if err != nil {
		log.Fatal(err)
	}
	
	// Convert into a string
	readString := string(stream)
	
	fmt.Println(readString)
	
	// EXCEPTING INPUT
	
	fmt.Println("What is your name? ")
	
	var name string
	
	fmt.Scan(&name)
	
	fmt.Println("Hello", name)
	
	// CASTING
	
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"
	
	// Convert numbers types
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	
	// Convert a string into an int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
    fmt.Println(newInt)
    
    // Convert a string into a float
    newFloat, _ := strconv.ParseFloat(randString2, 64)
    fmt.Println(newFloat)
 
}
*/



// create http server
/*
package main
 
import (
    "fmt"
    "net/http"
)
 
// CREATE A HTTP SERVER
 
// http.ResponseWriter assembles the servers response and writes to 
// the client
// http.Request is the clients request
func handler(w http.ResponseWriter, r *http.Request) {
 
	// Writes to the client
    fmt.Fprintf(w, "Hello World\n")
}
 
func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}
 
func main() {
 
	// Calls for function handlers output to match the directory /
    http.HandleFunc("/", handler)
    
    // Calls for function handler2 output to match directory /earth
    http.HandleFunc("/earth", handler2)
    
    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)
}
*/


// Go Routines

/*
package main
 
import "fmt"
import "time"
 
// GO ROUTINES
 
func count(id int) {
	for i := 0; i < 10; i++ { 
		fmt.Println(id, ":", i) 
		
		// Pause the function for 1 second to allow other functions to execute
		time.Sleep(time.Millisecond * 1000)
	} 
} 
 
func main() { 
 
	// A go routine is a function that runs in parallel with other functions
	// We define one by using go followed by the function name
	
	for i := 0; i < 10; i++ {
		go count(i) // issues go routines, lighweight thread, for concurrency
	}
	
	// Wait for the timer to make sure the go routine has time to 
	// finish otherwise the program would end before that happens
	time.Sleep(time.Millisecond * 11000)
	
	// better option would be to use wg sync.WaitGroup, where we do wg.Add(1) for each new routine and wg.Done() 
	// when a routine(function) ends. And we do wg.Wait() to let them synchronize
}
*/


// Go channels

/*
package main
 
import "fmt"
import "time"
import "strconv"
 
// CHANNELS
// Channels allow us to pass data between go routines
 
var pizzaNum = 0
var pizzaName = ""
 
func makeDough(stringChan chan string){
 
	pizzaNum++
 
	// Convert int into a string
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
 
	fmt.Println("Make Dough and Send for Sauce");
	
	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName
	
	time.Sleep(time.Millisecond * 10)
 
}
 
func addSauce(stringChan chan string){
 
	// Receive the value passed on the channel
	pizza := <- stringChan
	
	fmt.Println("Add Sauce and Send", pizza, "for Toppings")
	
	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName
	
	time.Sleep(time.Millisecond * 10)
 
}
 
func addToppings(stringChan chan string){
 
	// Receive the value passed on the channel
	pizza := <- stringChan
	
	fmt.Println("Add Toppings to", pizza, "and Ship")
	
	time.Sleep(time.Millisecond * 10)
 
}
 
func main() { 
 
	// Make creates a channel that can hold a string
	// int channel intChan := make(chan int)
	stringChan := make(chan string) // chan is channel data type
	
	// Cycle through and make 3 pizzas
	for i := 0; i < 3; i++{
	
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		
		time.Sleep(time.Millisecond * 5000)
	
	}
 
}
*/
