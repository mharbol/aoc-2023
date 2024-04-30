package day20

type broadcaster struct {
    name string
	mods []string
}

func newBroadcaster(mods []string) *broadcaster {
    return &broadcaster{
        name: "broadcaster",
        mods: mods,
    }
}

func (b *broadcaster) getName() string {
	return b.name
}

func (b *broadcaster) modList() []string {
	return b.mods
}

func (b *broadcaster) receive(p pulse, from string, highCount, lowCount *int) pulse {
	if p == LOW_PULSE {
        *lowCount++
		*lowCount += len(b.mods)
		return LOW_PULSE
	} else {
		return HIGH_PULSE
	}
}

func (b *broadcaster) addIncoming(name string) {}
