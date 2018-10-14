package main

import (
	"fmt"
	"ml253/statemachine"
)

type testStateMachine struct {
	counter      int
	stateMachine statemachine.StateMachine
}

const (
	stateA = iota
	stateB
)

func main() {
	sm := testStateMachine{counter: 0}
	sm.stateMachine = *statemachine.NewStateMachine(func() { sm.counter = 0 })

	sm.stateMachine.
		AddState(stateA).
		AddState(stateB).
		AddStateChange(stateA, stateB, func() bool { return sm.counter == 2 }).
		AddStateChange(stateB, stateA, func() bool { return sm.counter == 3 })

	for i := 0; i < 20; i++ {
		fmt.Println(sm.stateMachine.Update())
		sm.counter++
	}
}
