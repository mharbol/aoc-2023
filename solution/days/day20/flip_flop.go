package day20

type flipFlop struct {
	modList []*string
	st      state
	name    string
}

func newFlipFlop(name string, modList []*string) *flipFlop {
	return &flipFlop{
		modList: modList,
		st:      OFF,
		name:    name,
	}
}

func (ff *flipFlop) getName() string {
	return ff.name
}

func (ff *flipFlop) moduleList() []*string {
	return ff.modList
}

func (ff *flipFlop) receive(p pulse, lowCount *int, highCount *int) pulse {
	if p == LOW_PULSE {
		if ff.st == OFF {
			ff.st = ON
			*highCount++
			return HIGH_PULSE
		} else {
			ff.st = OFF
			*lowCount++
			return LOW_PULSE
		}
	} else { // HIGH_PULSE
		return NO_PULSE
	}
}
