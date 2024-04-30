package day20

type conjunction struct {
	name     string
	mods     []string
	incoming map[string]pulse
}

func newConjunction(name string, mods []string) *conjunction {
	return &conjunction{
		name:     name,
		mods:     mods,
		incoming: make(map[string]pulse),
	}
}

func (c *conjunction) getName() string {
	return c.name
}

func (c *conjunction) modList() []string {
	return c.mods
}

func (c *conjunction) receive(p pulse, from string, highCount, lowCount *int) pulse {
	c.incoming[from] = p
	if c.allHigh() {
		*lowCount += len(c.mods)
		return LOW_PULSE
	} else {
		*highCount += len(c.mods)
		return HIGH_PULSE
	}
}

func (c *conjunction) allHigh() bool {
	for _, val := range c.incoming {
		if val == LOW_PULSE {
			return false
		}
	}
	return true
}

func (c *conjunction) addIncoming(name string) {
	c.incoming[name] = LOW_PULSE
}
