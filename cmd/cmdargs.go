package cmd

import "flag"

type ArgParser interface {
	GetDay() uint
	GetPart() uint
    Parse()
}

type CmdParser struct {
    day *uint
    part *uint
}

func (p *CmdParser) GetDay() uint {
    return *p.day
}

func (p *CmdParser) GetPart() uint {
    return *p.part
}

func (p *CmdParser) Parse() {
    p.day = flag.Uint("d", 0, "day")
    p.part = flag.Uint("p", 0, "part")
    flag.Parse()
}
