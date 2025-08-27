// Fanout demonstrates parallel steps, and the one place the determinism rule
// stops being obvious.
//
// Steps may run concurrently. What must stay deterministic is the *set* of
// step ids requested, not the order in which their results arrive. Deriving an
// id from completion order — "result:0", "result:1" — breaks replay, because
// the second execution will not complete in the same order.
package main

import "fmt"

func main() {
	fmt.Println("fanout: read the package comment")
}
