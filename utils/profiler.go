package utils

import (
	"fmt"
	"time"
)

// Profiler is used to record execution time and support ultimately depth of children methods.
// This would help debugging performance right on the server without any tool.
// Receivers of this are nil-safe (if pointer is nil, can still call without any problem)
type Profiler struct {
	desc      string      // description, each profiler will have their own description
	constDesc string      // constant description, will be kept and copied to children profilers constantly
	start     int64       // epoch ms, immediately assigned = epoch when instance created
	duration  int64       // total execution time in milliseconds
	finalized bool        // finalized state, init <false>, will be changed to <true> when Finalize* methods get called
	level     int         // level of child, master = 0, the first tier (children of master) = 1, the second tier (children of first tier) = 2
	err       bool        // finalized with at least one error
	children  []*Profiler // children of current instance
}

// NewMasterProfiler returns a new profiler instance
func NewMasterProfiler(desc string, constDesc string, enable bool) *Profiler {
	if !enable {
		return nil
	}
	master := newProfiler(desc, 0)
	master.constDesc = constDesc
	return master
}

// newProfiler returns instance with default value
func newProfiler(desc string, level int) *Profiler {
	return &Profiler{
		desc:      desc,
		start:     time.Now().UnixMilli(),
		finalized: false,
		level:     level,
	}
}

// NewChild inits and returns a new child of the current instance
func (p *Profiler) NewChild(formatDesc string, a ...any) *Profiler {
	if p == nil {
		return nil
	}
	child := newProfiler(fmt.Sprintf(formatDesc, a...), p.level+1)
	child.constDesc = p.constDesc
	p.children = append(p.children, child)
	return child
}

// Finalize stops the execution time counter for the current instance and seals it (can not finalize again).
func (p *Profiler) Finalize() *Profiler {
	if p == nil {
		return nil
	}

	if !p.finalized {
		p.duration = time.Now().UnixMilli() - p.start
		p.finalized = true
	}
	return p
}

// FinalizeWithCheckErr stops the execution time counter for the current instance and seals it, later if provide err, the <err> state can be changed
func (p *Profiler) FinalizeWithCheckErr(err error) {
	if p == nil {
		return
	}
	if !p.finalized {
		p.Finalize()
	}
	if err != nil {
		p.err = true
	} else {
		// ignore
	}
}

// FinalizeWithErr stops the execution time counter for the current instance and seals it, later if provide err, the <err> state can be changed
func (p *Profiler) FinalizeWithErr(err error) error {
	if p == nil {
		return err
	}
	if !p.finalized {
		p.Finalize()
	}
	if err != nil {
		p.err = true
	} else {
		panic("error must not be nil")
	}
	return err
}

// Print does system print out the record data, returns if itself or any child has err
func (p *Profiler) Print() (anyError bool) {
	if p == nil {
		return false
	}
	bzPad := make([]byte, p.level*2)
	for i := 0; i < len(bzPad); i++ {
		bzPad[i] = 32 // space
	}
	duration := p.duration
	if !p.finalized {
		duration = time.Now().UnixMilli() - p.start
	}
	fmt.Printf("%s[L%2d] desc [%8s][%-30s] err [%5t] duration [%7d] ms children [%2d]", string(bzPad), p.level, p.constDesc, p.desc, p.err, duration, len(p.children))
	if len(p.children) > 1 {
		fmt.Printf(" avg [%2d]/child\n", int(duration)/len(p.children))
	} else {
		fmt.Printf("\n")
	}

	anyErr := false
	if len(p.children) > 0 {
		for _, child := range p.children {
			if child.Print() {
				anyErr = true
			}
		}
	}
	if p.level == 0 {
		fmt.Printf("Finished [%8s][%-30s] duration [%7d] ms", p.constDesc, p.desc, duration)
		if anyErr {
			fmt.Printf(" with at least one ERR\n")
		} else {
			fmt.Printf(" gracefully\n")
		}
	}

	if p.err {
		return true
	}

	return anyErr
}
