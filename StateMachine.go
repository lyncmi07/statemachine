package statemachine

type StateMachine struct {
	currentState    int
	changeStateFunc func()
	allStates       map[int](map[int](func() bool))
}

//NewStateMachine creates a new state machine and returns it.
func NewStateMachine(changeStateFunc func()) *StateMachine {
	stateMachine := new(StateMachine)
	stateMachine.allStates = make(map[int](map[int](func() bool)), 0)
	stateMachine.changeStateFunc = changeStateFunc

	return stateMachine
}

//Update performs all logic for changing states as required.
//It returns the ID of the current state it is in.
func (sm *StateMachine) Update() int {
	for destinationState, changeStateFunc := range sm.allStates[sm.currentState] {
		if changeStateFunc() {
			sm.currentState = destinationState
			sm.changeStateFunc()
		}
	}

	return sm.currentState
}

//AddState adds a new state to the StateMachine.
//The first state to be added is the start state of the machine.
func (sm *StateMachine) AddState(stateID int) *StateMachine {
	if len(sm.allStates) == 0 {
		sm.currentState = stateID
	}

	sm.allStates[stateID] = make(map[int](func() bool), 0)

	return sm
}

//Creates a state change between the fromState and the toState when the changeStateFunc evaluates to true.
func (sm *StateMachine) AddStateChange(fromState, toState int, changeStateFunc func() bool) *StateMachine {
	state := sm.allStates[fromState]

	state[toState] = changeStateFunc
	return sm
}

func (sm StateMachine) CurrentState() int {
	return sm.currentState
}
