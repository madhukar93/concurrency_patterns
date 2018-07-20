---

---
# CSP

what ?
formal description of real time behaviour of sytems

notation

- vendingMachine!coin (input)
- vendingMachine?chocolate (output)

this one's from the youtube vid - seems kind of diff.
- ⍺ Process = {event1, event2}
- x -> P (x (event) then P (event))
- recursion didn't understand
- μ (recursion)
- | (choice)
- || (concurrence)

you don't have to understand ^ to do csp, don't panic. The math just demonstrates
rigor, which we don't have to get into.

concurrency -

- a pattern of interaction of processes in real time.
- simultaneous execution.

## features

http://www.cs.cmu.edu/afs/cs.cmu.edu/user/brookes/www/papers/csp25.pdf

- sequential imperitive processes (procedural blocks ?) execute concurrently and communicate by synchronized input and output.
- no shared memory access ? Processes must have "disjoint local state"
- processes interact solely by message passing
- use guarded commands to handle concurrency

## implementation - python using gevent maybe.

## implementation - go channels

## applications

- simple producer consumer
- readers writers problem
- dining philosphers problem
- non blocking IO
- simple cli chat app over raw sockets or something ?

## comparison with actor model ?
- https://en.wikipedia.org/wiki/Communicating_sequential_processes#Comparison_with_the_actor_model
- https://engineering.universe.com/introduction-to-concurrency-models-with-ruby-part-ii-c39c7e612bed
- https://news.ycombinator.com/item?id=7612775
- https://www.quora.com/What-are-the-differences-between-the-actor-model-and-Communicating-Sequential-Processes-CSP-and-when-should-each-be-used
- https://www.quora.com/How-are-Akka-actors-different-from-Go-channels-How-are-two-related-to-each-other

## limitations ?

## links

 - https://golang.org/doc/codewalk/sharemem/
 - https://www.youtube.com/watch?v=3gXWA6WEvOM : Eric Shul: CSP 2015
 - http://divan.github.io/posts/go_concurrency_visualize/