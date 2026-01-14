package mission

import (
	"fmt"
	"sync"
)

type State int

const (
	STATE_IDLE State = iota
	STATE_TAKEOFF
	STATE_LAP
	STATE_SCAN
	STATE_STITCH_DELIVER
	STATE_LAND
)

type StateHandler func(*DroneFSM) State

type DroneFSM struct {
	currentState State
	numDelivered int
	handlers     map[State]StateHandler
}

func NewFSM() *DroneFSM {
	fsm := &DroneFSM{
		currentState: STATE_IDLE,
		numDelivered: 0,
	}

	fsm.handlers = map[State]StateHandler{
		STATE_IDLE:           (*DroneFSM).handleIdle,
		STATE_TAKEOFF:        (*DroneFSM).handleTakeoff,
		STATE_LAP:            (*DroneFSM).handleLap,
		STATE_SCAN:           (*DroneFSM).handleScan,
		STATE_STITCH_DELIVER: (*DroneFSM).handleStitchDeliver,
		STATE_LAND:           (*DroneFSM).handleLand,
	}

	return fsm
}

func (f *DroneFSM) Run() {
	for {
		handler, ok := f.handlers[f.currentState]
		if !ok {
			fmt.Printf("No handler for state: %d\n", f.currentState)
			return
		}

		nextState := handler(f)
		if nextState != f.currentState {
			fmt.Printf("State transition: %d -> %d\n", f.currentState, nextState)
		}
		f.currentState = nextState
	}
}

func (f *DroneFSM) handleIdle() State {
	// Add handling idle code here
	return STATE_TAKEOFF
}

func (f *DroneFSM) handleTakeoff() State {
	// Push takeoff to drone
	return STATE_LAP
}

func (f *DroneFSM) handleLap() State {
	// Add logic to push missions from yaml spec
	return STATE_SCAN
}

func (f *DroneFSM) handleScan() State {
	// Add scanning logic here
	return STATE_STITCH_DELIVER
}

func (f *DroneFSM) handleStitchDeliver() State {
	fmt.Println("Target Acquired. Pausing Scan...")

	var g errgroup.Group

	g.Go(func() error {
		return 
	})

	g.Go(func() error {
		return
	})

	if err := g.Wait(); err != nil {
        fmt.Printf("Stitch/deliver failed: %v\n", err)
        return STATE_LAND
    }

	f.numDelivered++
	if f.numDelivered >= 4 {
		return STATE_LAND
	}
	return STATE_SCAN
}

func (f *DroneFSM) handleLand() State {
	// Push land to drone
	return STATE_IDLE
}
