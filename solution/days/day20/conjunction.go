package day20

type conjunction struct {
	modList   []*string
	lastPulse pulse
	name      string
}

func newConjunction(name string, modList []*string) *conjunction {
	return &conjunction{
		modList:   modList,
		lastPulse: LOW_PULSE,
		name:      name,
	}
}

func (c *conjunction) getName() string {
	return c.name
}

func (c *conjunction) receive(p pulse, lowCount *int, highCount *int) pulse {
	last := c.lastPulse
	c.lastPulse = p
	if last == LOW_PULSE {
		*highCount++
		return HIGH_PULSE
	} else {
		*lowCount++
		return LOW_PULSE
	}
}

func (c *conjunction) moduleList() []*string {
	return c.modList
}
