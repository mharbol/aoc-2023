package day20

type flipFlop struct {
	name  string
	state state
	mods  []string
}

func (ff *flipFlop) getName() string {
	return ff.name
}

func (ff *flipFlop) modList() []string {
	return ff.mods
}

func (ff *flipFlop) receive(p pulse, from string, highCount, lowCount *int) pulse {
	if p == HIGH_PULSE {
		return NO_PULSE
	} else { // LOW_PULSE
		if ff.state == OFF {
			ff.state = ON
			*highCount += len(ff.mods)
			return HIGH_PULSE
		} else { // ON
			ff.state = OFF
			*lowCount += len(ff.mods)
			return LOW_PULSE
		}
	}
}

func newFlipFlop(name string, mods []string) *flipFlop {
	return &flipFlop{
		name:  name,
		state: OFF,
		mods:  mods,
	}
}

func (ff *flipFlop) addIncoming(name string) {}

type state bool

const (
	ON  state = true
	OFF state = false
)
