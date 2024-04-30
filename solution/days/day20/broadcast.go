package day20

type broadcast struct {
	modList []*string
	name    string
}

func newBroadcast(name string, modList []*string) *broadcast {
	return &broadcast{
		modList: modList,
		name:    name,
	}
}

func (b *broadcast) getName() string {
	return b.name
}

func (b *broadcast) receive(p pulse, lowCount *int, highCount *int) pulse {
	if p == LOW_PULSE {
		*lowCount++
	} else {
		*highCount++
	}
	return p
}

func (b *broadcast) moduleList() []*string {
	return b.modList
}
