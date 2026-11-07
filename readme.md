```
# initialize a folder as a mod
cd cards

go mod init cards
```


# Maps

## Maps vs Structs

![alt text](image.png)


# Interfaces

Makes it easier to reuse code in Go.

# Channels and go routines

<b>Concurrency</b> = how you arrange the work. 

<b>Parallelism</b> = how many workers actually do it at the same moment.

## Think of a coffee shop

Concurrency is one barista taking an order, putting the milk to steam, and while it heats, going back to the register to take the next order. Nothing happens at the same instant — but things overlap, and nobody stands idle. One person, well organized.

Parallelism is three baristas, three espresso machines. Now three drinks are genuinely being made at the same second. More people.

One barista can be very concurrent. 
One barista can never be parallel.

## So what does Go give you?

Go gives you the barista's organization — not the extra baristas.


go func(){} gives you the well-organized barista: while goroutine A waits on the network, goroutine B runs. 

That costs nothing and doesn't need extra cores.

The extra baristas are your 4 physical cores, and Go's runtime borrows them from the OS. You don't create them; you only get to use the ones the machine has.


Put together: you write concurrency; the machine decides if it becomes parallelism.

## Why you should care

Say you have 20 API calls, each taking 200ms.

Sequential: 4 seconds. You're standing there doing nothing while each call waits.

With go: ~220 milliseconds. But here's the catch — this happens even on one core. 

Nothing ran at the same time. All you did was stop waiting.

That's the win most people actually want, and it's concurrency doing it, not parallelism.


## The trap

The mistake is thinking go = faster because "it runs in parallel."
Waiting on the network, a database, a file? → go helps a lot. The waiting overlaps.
Pure number crunching in a loop? → go doesn't help. You have 4 cores, and adding 1,000 goroutines just makes 996 of them queue up.

## Goroutine syntax

**Start a goroutine** by putting `go` before a function call:
 
```go
go myFunction()              // named function
go myFunction(5, "hi")       // with arguments
go func() {                  // anonymous function
    fmt.Println("running")
}()                          // the () at the end calls it
```
 
Passing values into an anonymous function is the safe way to avoid surprises with loop variables:
 
```go
go func(n int) {
    fmt.Println(n)
}(i)
```
 
**Waiting for goroutines:** `main` doesn't wait for them, so if `main` ends, they die. A `sync.WaitGroup` is the usual fix:
 
```go
var wg sync.WaitGroup
 
wg.Add(1)              // "one more worker to wait for"
go func() {
    defer wg.Done()    // "I'm finished"
    fmt.Println("work")
}()
 
wg.Wait()              // block until all are done
```
 
## Channel syntax
 
**Create:**
 
```go
ch := make(chan int)        // unbuffered: send waits for a receiver
ch := make(chan int, 3)     // buffered: holds up to 3 values before blocking
```
 
The type after `chan` is what the channel carries (`int`, `string`, a struct, etc.).
 
**Send and receive** (the arrow shows data direction):
 
```go
ch <- 42          // send 42 into ch
x := <-ch         // receive from ch into x
<-ch              // receive and discard
```
 
**Close** (only the sender should do this):
 
```go
close(ch)
```
 
**Loop until closed:**
 
```go
for v := range ch {
    fmt.Println(v)
}
```
 
**Check if closed:**
 
```go
v, ok := <-ch     // ok is false if ch is closed and empty
```
 
**Directional channels** (in function parameters, to limit what a function can do):
 
```go
func producer(out chan<- int) { out <- 1 }   // send-only
func consumer(in <-chan int)  { x := <-in }  // receive-only
```
 
**Select** (wait on several channels at once):
 
```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case ch2 <- 5:
    fmt.Println("sent to ch2")
case <-time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("nothing ready")   // optional, makes it non-blocking
}
```
 
## Quick reference
 
| Syntax | Meaning |
|---|---|
| `go f()` | run `f` concurrently |
| `make(chan T)` | unbuffered channel of `T` |
| `make(chan T, n)` | buffered channel, capacity `n` |
| `ch <- v` | send `v` |
| `v := <-ch` | receive |
| `close(ch)` | no more values will be sent |
| `for v := range ch` | receive until closed |
| `chan<- T` / `<-chan T` | send-only / receive-only |
| `select { ... }` | wait on multiple channel operations |
 
**Memory trick:** the arrow points the way the data flows. `ch <- v` pushes into the channel, and `<-ch` pulls out of it.
 
 ## Await goroutine

 ### sync.WaitGroup

 #### Why use `defer wg.Done()`?

You're right, it works without `defer` in that example. `defer` doesn't change the result there; it's about **safety**, and it only matters once the function gets more complex.

Same thing, without defer

```go
go func(n int) {
    fmt.Println("worker", n)
    wg.Done()           // called manually at the end
}(i)
```

In this tiny function, nothing can go wrong between the start and the `Done()` call, so both versions behave identically.

#### Where defer starts to matter

**1. Early returns**

```go
go func(n int) {
    if n == 2 {
        return              // Done() never runs!
    }
    fmt.Println("worker", n)
    wg.Done()
}(i)
```

Worker 2 exits early, the counter never reaches 0, and `wg.Wait()` blocks forever (deadlock).

With `defer wg.Done()` at the top, it runs on every exit path, including that early `return`.

**2. Panics**

```go
go func(n int) {
    defer wg.Done()         // still runs
    doRiskyThing()          // might panic
}(i)
```

If a panic is recovered somewhere, deferred calls still run. A manual `wg.Done()` at the bottom would be skipped.

**3. Future edits**

Someone adds an `if err != nil { return }` in the middle six months from now. With `defer`, the code stays correct. Without it, you've introduced a hang.

#### Rule of thumb

Put `defer wg.Done()` as the **first line** of the goroutine. That way you can't forget it, and it works no matter how the function exits.

So: not required for correctness in a trivial function, but a good habit because it protects you when the code grows.