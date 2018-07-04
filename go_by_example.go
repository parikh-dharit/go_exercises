package main

import (
    "os"
    "fmt"
    //"io"
    //"io/ioutil"
    //"bufio"
    //"os/exec"
    //"os/signal"
    //"syscall"
    //"encoding/json"
    //"bytes"
    //"regexp"
    //s "strings"
//import "math" // if not used will give error
// or you can do import _ "time"
//import "errors"
    //"sort"
    //"time"
    //"sync/atomic"
    //"math/rand" // pseudorandom number generation
    //"sync"
    //"strconv" // convert strings
    //"net"
    //"net/url"
    //"crypto/sha1" // there are other hash functions in crypto/ 
    //b64 "encoding/base64"
    //"flag"
)
//var p = fmt.Println
func main(){
    // 1. Hello World
    //fmt.Println("hello world");
    
    // 2. Values
    /*
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println("T & F =", (true && false))
    fmt.Println("T | F =", (true || false))
    fmt.Println("~T =", (!true))
    */
    
    // 3. Variables
    /*
    var a = "initial" // type will be infered from the assigned value
    fmt.Println(a)
    
    var b, c int = 1, 2
    fmt.Println(b, c)
    
    var d = true
    fmt.Println(d)
    
    var e int // initialiazed to default values
    fmt.Println(e)
    
    f:="short" // declare and assign
    fmt.Println(f)
    */
    
    // 4. Constants
    /*
    const s string = "constant"
    fmt.Println(s)
    
    const n = 500000000
    const d = 3e20 / n
    
    fmt.Println(d)
    fmt.Println(int64(d)) //numeric constant has no type until it’s given one
    
    fmt.Println(math.Sin(n))
    */
    
    // 5. For
    /*
    i := 1
    for i<= 3 { // with single condition
        fmt.Println(i)
        i = i+1
    }
    
    for j:=7; j<=9; j++{
        fmt.Println(j)
    }
    
    for { // forever loop
        fmt.Println("loop")
        break;
    }
    
    for n:=0; n<=5; n++ {
        if n%2==0 {
            continue
        }
        fmt.Println(n)
    }
    */
    
    
    // 6. If Else
    /*
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else { // not necessary to write else
        fmt.Println("7 is odd")
    }
    
    if num:=9; num<0 {
        fmt.Println(num, "is negative")
    } else if num < 10 { // cannot write in the new line
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
    */
    
    // 7. Switch
    /*
    i := 2
    // various types of print: Print, Printf, Println
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday: // multiple conditions
        fmt.Println("It's the weekend'")
        fmt.Println("Hoorray!")
    default:
        fmt.Println("It's a weekday")
    }
    
    t := time.Now()
    switch { // can be used instead of if-else
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
    
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't konw type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
    */
    
    // 8. Arrays
    /*var a [5] int
    fmt.Println("emp: ", a)
    
    a[4] = 100
    fmt.Println("set: ", a)
    fmt.Println("get: ", a[4])
    
    fmt.Println("len: ", len(a))
    
    b := [5] int {1, 2, 3, 4, 5}
    fmt.Println("dcl: ", b)
    
    var twoD [2][3] int
    
    for i:=0; i<2; i++ {
        for j:=0; j<3; j++ {
            twoD[i][j] = i+j
        }
    }
    fmt.Println("2d: ", twoD)    
    */
    
    // 9. Slices
    /*s := make([]string, 3)
    fmt.Println("emp: ", s)
    
    s[0] = "a" // single column does not work ' '
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set: ", s)
    fmt.Println("get: ", s[2])
    
    fmt.Println("len: ", len(s))
    
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd: ", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy: ", c)
    
    l := s[2:5]
    fmt.Println("sl1: ", l)
    
    l1 := s[:5]
    fmt.Println("sl2: ", l1)

    l2 := s[2:]
    fmt.Println("sl3: ", l2)

    t := []string{"g", "h", "i"}
    fmt.Println("dccl: ", t)
    
    twoD := make([][]int, 3)
    
    for i:=0; i<3; i++ {
        innerLen := i+1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++{
            twoD[i][j] = i+j
        }
    }
    fmt.Println("2d: ", twoD)
    */
    
    // 10. Maps
    /*m := make(map[string]int)
    
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map: ", m)
    
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    
    fmt.Println("len: ", len(m))
    
    delete (m, "k2")
    fmt.Println("map: ", m)
    
    _, prs := m["k2"] // blank identifier, ignores the value
    fmt.Println("prs:", prs) //check if present
    
    n:= map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map: ", n)
    */
    
    // 11. Range
    /*nums := []int{2,3,4}
    sum := 0
    for _, num := range nums{
        sum += num
    }
    fmt.Println("sum: ", sum)
    
    for i, num := range nums{
        if num == 3{
            fmt.Println("index: ", i)
        }
    }
    
    kvs := map[string]string{"a": "apple", "b": "banana"}
    
    for k,v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    
    for k := range kvs {
        fmt.Println("key: ", k)
    }
    
    for i, c := range "go" {
        fmt.Println(i, c    )
    }*/

    // 12. Functions
    /*
    res := plus(1,2)
    fmt.Println("1+2 =", res)
    
    res = plusPlus(1,2,3)
    fmt.Println("1+2+3 =", res)
    */
    
    // 13. Multiple return values
    /*a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
        
    _, c := vals()
    fmt.Println(c)
    */
    
    // 14. Variadic Functions
    /*sums(1, 2)
    sums(1, 2, 3) // arbitrary number of arguments
    n := []int{1, 3, 5, 7}
    sums(n...) // we are sending a slice as multiple arguments, does not work without ...
    */
    
    // 15. Closure
    /*nextInt := intSeq()
    
    fmt.Println(nextInt()) // closes over i of intSeq for nextInt
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInt := intSeq() // new i and new nextInt
    fmt.Println(newInt())
    */
    
    // 16. Recursion
    /*
    fmt.Println(fact(7))
    */
    
    // 17. Pointers
    /*i:=1
    fmt.Println("initial: ", i)
    
    zeroval(i)
    fmt.Println("zeroval: ", i)
    
    zeroptr(&i)
    fmt.Println("zeroptr: ", i)
    fmt.Println("ptr: ", &i)*/
    
    // 18. Structs
    /*fmt.Println(person{"Bob", 20})
    fmt.Println(person{name:"Alice", age:30})
    
    fmt.Println(person{name:"Fred"}) //other field initialiazed to zero
    fmt.Println(&person{name:"Ann", age:40})
    
    s:= person{name:"Sean", age:50}
    fmt.Println(s.name)
    
    sp := &s
    fmt.Println(sp.age)
    
    sp.age = 51
    fmt.Println(sp.age)*/
    
    // 19. Methods
    /*r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim: ", r.perim())
    
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim: ", rp.perim())
    */
    
    // 20. Intefaces
    //Interfaces are named collections of method signatures.
    /*r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    
    measure(r)
    measure(c)
    */
    
    // 21. Errors
    /*for _, i := range []int{7, 42} {
        if r, e:= f1(i); e!=nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }

    for _, i := range []int{7, 42} {
        if r, e:= f2(i); e!=nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }
    
    _, e := f2(42)
    // working with errors manually
    if ae, ok := e.(*argError); ok{
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }*/
    
    // 22. Goroutines , lighweight-threads
    /*f("direct")
    go f("goroutine")
    //fmt.Println("after go f")
    
    go func (msg string){
        fmt.Println(msg)
    }("going")
    
    fmt.Scanln() //enter(accept) value from the terminal
    fmt.Println("done")*/
    
    // 23. Channels , IPC pipes to connect concurrent goroutines
    /*messages := make(chan string) // channel can only be created using make, make can only create slices, maps and channels
    // unbuffered channel, channel size 0
    
    go func() { messages <- "ping"}()// sending signal to messages channel with message "ping"
    
    msg := <- messages // receive signal from messages channels
    fmt.Println(msg)*/
    
    // 24. Channel Buffering
    /*messages := make(chan string, 2)
    messages <- "hello"
    messages <- "world" // not necessary to send two messages
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    */
    
    // 25. Channel Synchronization
    /*done := make(chan bool, 1) // not necessary to make buffer of size 1 in this case
    go worker(done)
    <-done// not using the value just waiting for it to complete
    */
    
    // 26. Channel directions , only send or receive from the channel
    /*pings := make(chan string, 1)
    pongs := make(chan string, 1)
    
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)*/
    
    // 27. Select , wait for multiple channel operations
    /*c1 := make(chan string)
    c2 := make(chan string)
    
    go func(){
        time.Sleep(time.Second*2)
        c1 <- "one"
    }()
    go func(){
        time.Sleep(time.Second*1)
        c2 <- "two"
    }()
    // difference between switch and select
    // A select blocks until one of its cases can run, then it executes that case. 
    // It chooses one at random if multiple are ready.
    for i:=0; i<2; i++ {
        select{
            case msg1 := <-c1:
                fmt.Println("received", msg1)
            case msg2 := <-c2:
                fmt.Println("received", msg2)
        }
    }*/
    
    // 28. Timeout
    /*c1 := make(chan string, 1)
    go func(){
        time.Sleep(time.Second*2)
        c1 <- "result 1"
    }()
    
    select{
        case res := <-c1:
            fmt.Println(res)
        case <-time.After(1*time.Second):
            fmt.Println("timeout 1")
    }
    
    c2 := make(chan string, 1)
    go func(){
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    
    select {
        case res:= <-c2:
            fmt.Println(res)
        case <-time.After(3 * time.Second):
            fmt.Println("timeout 2")
    }*/
    
    // 29. Non-Blocking channel operations
    //Basic sends and receives on channels are blocking. 
    //However, we can use select with a default clause to implement non-blocking sends, receives, 
    //and even non-blocking multi-way selects.
    
    /*messages := make(chan string)
    signals := make(chan bool)
    
    select {
        case msg := <-messages:
            fmt.Println("received msg", msg)
        default:
            fmt.Println("no messages received") 
    }
    
    msg:="hi"
    select {
        case messages <- msg:// channel is unbuffered and there is no receiver waiting so it will be blocked here
            fmt.Println("sent message", msg)
        default:
            fmt.Println("no messages sent") 
    }
    
    select {
        case msg := <-messages:
            fmt.Println("received msg", msg)
        case sig := <-signals:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity") 
    }*/
    
    // 30. Closing Channels
    /*
    jobs := make(chan int, 5)
    done := make(chan bool)
    
    go func(){
        for {
            j, more := <- jobs// when the channel closes more will be false, given all the values in the buffer have been received
            //fmt.Println("more", more)
            if more{
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
    
    for j:=1; j<=3; j++{
        jobs<-j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
    
    <-done
    */
    
    // 31. Range over channel
    /*queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    
    close(queue)
    
    for elem:= range queue {
        fmt.Println(elem)
    }*/
    
    // 32. Timers , when you want to do something once, after some time
    /*timer1 := time.NewTimer(2*time.Second)
    
    <- timer1.C //channel C blocks until timer1 expires
    fmt.Println("Timer 1 expired")
    
    timer2 := time.NewTimer(time.Second)
    go func(){
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }*/
    
    // 33. Tickers , do something repeatedly at a time interval
    /*ticker := time.NewTicker(500 * time.Millisecond)
    
    go func(){
        for t:= range ticker.C {
            fmt.Println("tick at", t)
        }
    }()
    
    time.Sleep(1600*time.Millisecond)
    ticker.Stop()
    fmt.Println("Ticker stopped")
    */
    
    // 34. Worker Pools
    /*jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    for w:=1; w<=3; w++{
        go worker(w, jobs, results)
    }
    
    for j:=1; j<=5; j++{
        jobs <- j
    }
    
    close(jobs)
    
    for a:=1; a<=5; a++{
        <-results
    }*/
    
    // 35. Rate Limiting
    /*requests := make(chan int, 5)
    for i:=1; i<=5; i++{
        requests <- i;
    }
    close(requests)
    limiter := time.Tick(500*time.Millisecond)
    
    for req:= range requests{
        <- limiter
        fmt.Println("request", req, time.Now())
    }
    
    burstyLimiter := make(chan time.Time, 3)
    
    for i:=0; i<3; i++{
        burstyLimiter <- time.Now()//3 same time added
    }
    
    go func(){
        for t:=range time.Tick(1000*time.Millisecond){
            burstyLimiter <- t
        }
    }()
    
    burstyRequests := make(chan int, 5)
    for i:=1; i<=5; i++{
        burstyRequests <- i
    }
    close(burstyRequests)
    for req:= range burstyRequests{
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
    */
    
    // 36. Atomic counters
    /*var ops uint64
    
    for i:=0; i<50; i++{
        go func(){
            for {
                atomic.AddUint64(&ops, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    time.Sleep(time.Second)
    
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)*/
    
    // 37. Mutexes
    /*var state = make(map[int]int)
    
    var mutex = &sync.Mutex{}
    
    var readOps uint64
    var writeOps uint64
    
    for r:=0; r<100; r++{
        go func(){
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    
    for w:=0; w<10; w++{
        go func(){
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    
    time.Sleep(time.Second)
    
    mutex.Lock()
    fmt.Println("state1:", state)
    mutex.Unlock()
    
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
    
    time.Sleep(time.Millisecond * 1000)
    
    mutex.Lock()
    fmt.Println("state2:", state)
    mutex.Unlock()
    */
    
    // 38. Stateful Goroutines
    // instead of sharing a resource, the resource is owned by a goroutine and 
    //other goroutines will send msg to this goroutine and receive reply
    /*var readOps uint64
    var writeOps uint64
    
    reads := make(chan *readOp) // channel of structure, which has another channel as part of it
    writes := make(chan *writeOp)
    
    // go routine that owns the resource -> state variable
    go func(){
        var state = make(map[int]int)
        for{
            select{
                case read := <-reads:
                    read.resp <- state[read.key]
                case write := <-writes:
                    state[write.key] = write.val
                    write.resp <- true
            }
        }
    }()
    
    for r:=0; r<100; r++{
        go func(){
            for {
                read := &readOp{
                    key: rand.Intn(5),
                    resp: make(chan int)}
                reads <- read // sends read on the read channel of the resource holding goroutine
                <-read.resp // waits till resource holding go routine sends a msg on read.resp channel 
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    
    for w:=0; w<10; w++{
        go func(){
            for {
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <- write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    
    time.Sleep(time.Second)
        
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
    */
    
    // 39. Sort
    /*strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings: ", strs)
    
    ints := []int{7, 2, 4}    
    sort.Ints(ints)
    fmt.Println("Ints: ", ints)
    
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted:",s)*/
    
    // 40. Sorting by Functions
    /*fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)*/
    
    // 41. Panic
    /*panic("a problem")
    
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }*/
    
    // 42. Defer
    /*f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)*/
    
    // 43. Collection Functions
    /*var strs = []string{"peach", "apple", "pear", "plum"}
    
    fmt.Println(Index(strs, "pear"))
    
    fmt.Println(Include(strs, "grape"))
    
    fmt.Println(Any(strs, func (v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    
    fmt.Println(All(strs, func (v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    
    fmt.Println(Filter(strs, func (v string) bool {
        return strings.Contains(v, "e")
    }))
    
    fmt.Println(Map(strs, strings.ToUpper))
    */
    
    // 44. String Functions
    /*p("Contains: ", s.Contains("test", "es"))
    p("Count: ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index: ", s.Index("test", "e"))
    p("Join: ", s.Join([]string{"a","b"},"-"))
    p("Repeat: ", s.Repeat("a", 5))
    p("Replace: ", s.Replace("foo", "o", "0", -1))
    p("Replace: ", s.Replace("foo", "o", "0", 1))
    p("Split: ", s.Split("a-b-c-d-e", "-"))
    p("ToLower: ", s.ToLower("TEST"))
    p("ToUpper: ", s.ToUpper("test"))
    p()
    
    p("Len: ", len("hello"))    
    p("Char: ", "hello"[1])*/
    
    // 45. String Formatting
    /*p := point{1, 2}
    fmt.Printf("%v\n", p)
    
    fmt.Printf("%+v\n", p)
    
    //The %#v variant prints a Go syntax representation of the value    
    fmt.Printf("%#v\n", p)
    
    //To print the type of a value, use %T.
    fmt.Printf("%T\n", p)
    
    //Formatting booleans is straight-forward.
    fmt.Printf("%t\n", true)
    
    //There are many options for formatting integers. Use %d for standard, base-10 formatting.
    fmt.Printf("%d\n", 123)
    
    //This prints a binary representation.
    fmt.Printf("%b\n", 14)
    
    //This prints the character corresponding to the given integer.
    fmt.Printf("%c\n", 33)
    
    //%x provides hex encoding.
    fmt.Printf("%x\n", 456)
    
    //There are also several formatting options for floats. For basic decimal formatting use %f.
    fmt.Printf("%f\n", 78.9)
    
    //%e and %E format the float in (slightly different versions of) scientific notation.
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    
    //For basic string printing use %s.
    fmt.Printf("%s\n", "\"string\"")
    
    //To double-quote strings as in Go source, use %q.
    fmt.Printf("%q\n", "\"string\"")
    
    //As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
    fmt.Printf("%x\n", "hex this")

    //To print a representation of a pointer, use %p.
    fmt.Printf("%p\n", &p)
    
    //To specify the width of an integer, use a number after the % in the verb.
    //By default the result will be right-justified and padded with spaces.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    //specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    //To left-justify, use the - flag.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    //to control width when formatting strings, especially to ensure that they align in table-like output.
    //For basic right-justified width.
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    //To left-justify use the - flag as with numbers.
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
    
    //So far we’ve seen Printf, which prints the formatted string to os.Stdout. 
    //Sprintf formats and returns a string without printing it anywhere.
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // format+print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
    */
    
    // 46. Regular Expressions
    /*match, _ := regexp.MatchString("p([a-z])+ch", "peach") //matching returns bool
    fmt.Println(match)
    
    r, _ := regexp.Compile("p([a-z]+)ch") // have to compile for other operations
    
    fmt.Println(r.MatchString("peach"))
    
    fmt.Println(r.FindString("peach punch")) // finds the first match and returns it
    
    fmt.Println(r.FindStringIndex("peach punch")) // returns start and end index of the string
    
    fmt.Println(r.FindStringSubmatch("peach punch")) // returns the match for whole r and for [a-z]+ as well, [0 5 1 3]
    
    fmt.Println(r.FindStringSubmatchIndex("peach punch")) // indexes of above
    
    fmt.Println(r.FindAllString("peach punch pinch", -1)) // returns all matches
    
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // indexes of matches and submatches
    
    fmt.Println(r.FindAllString("peach punch pinch", 2)) // limits the number of matches to 2
    
    fmt.Println(r.Match([]byte("peach")))
    
    r = regexp.MustCompile("p([a-z]+)ch") // when you have to work on constants
    fmt.Println(r)
    
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // replace subset of strings
    
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
    */
    
    // 47. JSON
    /*bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
    
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    
    mapD := map[string]int{"apple":5, "lettuce":7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
    
    res1D := &response1{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
    
    res2D := &response2{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
    
    byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
    
    var dat map[string]interface{} // decoded json will be stored here
    
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)
    
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)
    
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{} // decode to custom data types
    
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
    
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple":5, "lettuce":7}
    enc.Encode(d) // encodes and puts into os.Stdout
    */
    
    // 48. Time
    /*p := fmt.Println
    now := time.Now()
    p(now)
    
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC) //using time structure, have to provide time zone 
    // Year, Month, Day, H, M, S, Nanosecond, Location
    p(then)
    
    p("Year", then.Year())
    p("Month", then.Month())
    p("Day", then.Day())
    p("Hour", then.Hour())
    p("Minute", then.Minute())
    p("Second", then.Second())
    p("Nanosecond", then.Nanosecond())
    p("Location", then.Location())
    
    p("Weekday", then.Weekday())
    
    p(then.Before(now)) //checks which one came first
    p(then.After(now))
    p(then.Equal(now))
    
    diff := now.Sub(then) // subtracts, gives durations between th two
    p(diff)
    
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
    
    p(then.Add(diff)) // add to the time, advance it by diff
    p(then.Add(-diff)) // or go back diff time amount
    */
    
    // 49. Epoch , time elapsed since January 1, 1970
    /*now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)
    
    millis := nanos/1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)
    
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
    */
    
    // 50. Time Formatting / Parsing
    /*p := fmt.Println
    
    t := time.Now()
    p(t.Format(time.RFC3339))
    
    t1, e := time.Parse(
        time.RFC3339,
        "2016-11-01T22:08:41+00:00")
    p(t1)
   
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)
    // custom layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern 
    // with which to format/parse a given time/string
    
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
    
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e) // error in parse
    */
    
    // 51. Random Numbers
    /*fmt.Print(rand.Intn(100), ",")
    fmt.Println(rand.Intn(100))
    
    fmt.Println(rand.Float64())
    
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Println((rand.Float64()*5)+5)
    
    s1 := rand.NewSource(time.Now().UnixNano()) // giving a seed that changes
    r1 := rand.New(s1)
    
    fmt.Print(r1.Intn(100), ",")
    fmt.Println(r1.Intn(100))
    
    // to reproduce the same sequence
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Println(r2.Intn(100))
    
    // source the same seed
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Println(r3.Intn(100))
    */
    
    // 52. Number Parsing
    /*f, _ := strconv.ParseFloat("1.234", 64) // 64 bits of precision to parse
    fmt.Println(f)
    
    i, _ := strconv.ParseInt("123", 0, 64) // should fit in 64 bit, 0 -> infer base from string
    fmt.Println(i)
    
    d, _ := strconv.ParseInt("0x1c8", 0, 64) // willl recognize hex formatted number
    fmt.Println(d)
    
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)
    
    k, _ := strconv.Atoi("134") //base-10 int parsing
    fmt.Println(k)
    
    _, e := strconv.Atoi("wud") // error on bad input
    fmt.Println(e)
    */
    
    // 53. URL Parsing
    /*s := "postgres://user:pass@host.com:5432/path?k=v#f"
    // scheme, authentication info, host, port, path, query params, query fragment
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(u.Scheme)
    
    fmt.Println(u.User) // user:pass
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)
    
    fmt.Println(u.Host) // host.com:5432
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)
    
    fmt.Println(u.Path) // /path
    fmt.Println(u.Fragment)// f
    
    fmt.Println(u.RawQuery) // k=v
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m) // [k:[v]]
    fmt.Println(m["k"][0]) // v
    */
    
    
    // 54. SHA1 Hashes
    /*s := "sha1 this string"
    
    h := sha1.New()
    
    h.Write([]byte(s)) // byte stores data, length and (not in string)capacity, it is not immutable(unlike string which can not be changed) 
    
    bs := h.Sum(nil)
    
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
    */
    
    // 55. Base64 Encoding
    /*data := "abc123!?$*&()'-=@~"
    
    // standard base64 encoding
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()
    
    // url compatible base64 encoding
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
    */
    
    // 56. Reading files
    /*dat, err := ioutil.ReadFile("/tmp/dat") // read the whole file into memory
    check(err)
    fmt.Print(string(dat))
    
    f, err := os.Open("/tmp/dat") // opening a file, getting os.File value
    check(err)
    
    b1 := make([]byte, 5)
    n1, err := f.Read(b1) // read upto 5 bytes if available
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    
    o2, err := f.Seek(6, 0) // go to a location and read from there
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
    
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2) //more robustly implemented
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
    
    _, err = f.Seek(0, 0) // similar to re-wind
    check(err)
    
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
    
    f.Close() // schedule immediately after Opening with defer
    */
    
    // 57. Writing Files
    /*d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)
    
    f, err := os.Create("/tmp/dat2")
    check(err)
    
    defer f.Close()
    
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
    
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)
    
    f.Sync() // flushes write to storage
    
    w := bufio.NewWriter(f) // buffered writer
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)
    
    w.Flush()
    */
    
    // 58. Line Filters , similar to grep, sed
    /*scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ucl := s.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }*/
    
    
    // 59. Command-Line Arguments
    // have to build the program first: go build program.go
    // and then: ./program arg1 arg2 arg3
    // instead of: go run program.go
    // os.Args
    /*argsWithProg := os.Args // first is path to the program
    argsWithoutProg := os.Args[1:]
    
    arg := os.Args[3]
    
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
    */
    
    // 60. Command-Line Flags
    /*wordPtr := flag.String("word", "foo", "a string")
    // defined a flag: word = foo [default], description= a string 
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")
    
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") // or we can pass a variable
    
    flag.Parse() // command line parsing
    
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
    // executed using: ./program -word=opt -numb=7 -fork -svar=flag
    // to see description: ./program -h
    */
    
    
    // 61. Environment Variables
    /*os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))
    
    fmt.Println()
    for _, e := range os.Environ() { // slice of all the envr variables
        pair := s.Split(e, "=")
        fmt.Println(pair[0])
    }
    */
    
    // 62. Spawning Processes
    /*dateCmd := exec.Command("date")
    
    dateOut, err := dateCmd.Output() 
    // helper that handles the common case of running a command,
    // waiting for it to finish, and collecting its output
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))
    
    grepCmd := exec.Command("grep", "hello")
    
    grepIn, _ := grepCmd.StdinPipe() // create input/output pipes
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start() // start the process
    grepIn.Write([]byte("hello grep\ngoodbye grep")) // give input to process
    grepIn.Close() // close process
    grepBytes, _ := ioutil.ReadAll(grepOut) // read output of the process
    grepCmd.Wait() // wait for everything to finish
    
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))
    
    // to spawn a full command with a string, use bash’s -c option
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
    */
    
    // 64. Executing Processes , replacce our go process
    // starting goroutines, spawning processes, and exec’ing processes covers most use cases for fork
    /*binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }
    
    args := []string{"ls", "-a", "-l", "-h"}
    
    env := os.Environ()
    
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }*/
    
    // 65. Signals
    // want a server to gracefully shutdown when it receives a SIGTERM or behave a particular way on some SIGNAL
    /*
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // receive notification for given signals
    // ctrl+c -> SIGINT
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()
    
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
    */
    
    // 66. Exit
    // defers will not be run when using os.Exit
    
    defer fmt.Println("Pi")
    
    //Exit with status 3
    os.Exit(3)
    // By building and executing a binary you can see the status in the terminal
    //$ go build program.go
    //$ ./program
    //$ echo $?
    
}

// 12. Functions
/*
func plus(a int, b int) int {
    return a+b
}

func plusPlus(a, b, c int) int {
    return a+b+c
}
*/

// 13. Multiple return values
/*func vals() (int, int){
    return 3, 7
}*/


// 14. Variadic Functions
/*func sums (nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums{
        total += num
    }
    fmt.Println(total)
}*/

// 15. Closures
/*func intSeq() func() int {
    i := 0
    return func() int {//returns an anonymus function
        i++
        return i
    }
}*/

// 16. Recursion
/*func fact(n int) int {
    if n==0{
        return 1
    } else {
        return n*fact(n-1)
    }
}*/

// 17. Pointers
/*func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}*/

// 18. Structs
/*type person struct{
    name string
    age int
}*/

// 19. Methods
/*type rect struct{
    width, height int
}

func (r *rect) area () int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}*/

// 20. Interfaces
/*type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width*r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
// function (connected to structure) name (input parameters) return_variable_type {
//      return someVariable
//}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius //variables and methods with capital first letter can be accessed from other packages
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

type geometry interface {
    area() float64
    perim() float64
}

func measure (g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}*/

// 21. Errors
/*func f1(arg int) (int, error){// error is a built-in interface
    if arg == 42 {
        return -1, errors.New("can't work with 42")//constructs a new error value with this msg
    }
    return arg+3, nil
}

type argError struct{// custom type error with Error method attached to it
    arg int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg+3, nil
}*/


// 22. Goroutines
/*func f (from string) {
    for i:=0; i<3; i++{
        fmt.Println(from,":", i)
        time.Sleep(time.Millisecond*1000)//couldn't see going in between goroutines without this forced sleeping
    }
}*/

// 25. Channel Synchronization
/*func worker(done chan bool) {
    fmt.Println("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}*/

// 26. Channel directions
/*func ping(pings chan<- string, msg string){//sends on channel
    pings<-msg
}

func pong(pings <-chan string, pongs chan<- string){ // takes (receives) from pings channel and puts (sends) on pongs channel
    msg := <-pings
    pongs <- msg
}*/

// 34. Worker Pools
/*func worker(id int, job <-chan int, results chan<- int){
    for j := range job{
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("woker", id, "finished job", j)
        results <- j*2
    }
}*/

// 38. Stateful Goroutines
/*type readOp struct{
    key int
    resp chan int
}

type writeOp struct{
    key int
    val int
    resp chan bool
}*/

// 40. Sorting by Functions
/*type byLength []string

func (s byLength) Len() int {
    return len(s)
} 

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool{
    return len(s[i]) < len(s[j])
}*/

// 42. Defer
/*func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File){
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){
    fmt.Println("closing")
    f.Close()
}*/

// 43. Collection Functions
/*func Index(vs []string, t string) int {
    for i,v :=range vs{
        if v==t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool{
    return Index(vs, t) >=0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs{
        if f(v){
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs{
        if !f(v){
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs{
        if f(v){
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs{
        vsm[i] = f(v)
    }
    return vsm
}*/

// 45. String Fomatting
/*type point struct {
    x, y int
}*/


// 47. JSON
/*type response1 struct {
    Page int
    Fruits []string
}
// tagged structre for json keys, otherwise variable names will be taken as json keys
type response2 struct {
    Page int        `json:"page"`
    Fruits []string `json:"fruits"`
}*/


// 56. Reading Files
/*func check(e error){
    if e!=nil {
        panic(e)
    }
}*/



