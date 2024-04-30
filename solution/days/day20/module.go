package day20

import "strings"

type module interface {
	getName() string
	modList() []string
	receive(p pulse, from string, highCount, lowCount *int) pulse
	addIncoming(name string)
}

type pulse int

const (
	HIGH_PULSE pulse = 0
	LOW_PULSE  pulse = 1
	NO_PULSE   pulse = 2
)

func parseModules(lines []string) map[string]module {
	// populate all modules
	var mods = make(map[string]module)
	for _, line := range lines {
		mod := newModFromString(line)
		mods[mod.getName()] = mod
	}

	for name, mod := range mods {
		for _, outgoingModsName := range mod.modList() {
			_, ok := mods[outgoingModsName]
			if ok {
				mods[outgoingModsName].addIncoming(name)
			}
		}
	}

	return mods
}

func newModFromString(line string) module {
	parts := strings.Split(line, " ")
	first := parts[0][0]
	name := parts[0][1:]
	modsWithComma := parts[2:]
	var mods []string
	for _, str := range modsWithComma {
		mods = append(mods, strings.ReplaceAll(str, ",", ""))
	}
	if first == 'b' {
		return newBroadcaster(mods)
	} else if first == '%' {
		return newFlipFlop(name, mods)
	} else {
		return newConjunction(name, mods)
	}
}

type modPulse struct {
	modTo   string
	modFrom string
	pulse   pulse
}

func pressButton(modules map[string]module, highCount, lowCount *int) map[string]module {

	presses := []modPulse{{modTo: "broadcaster", modFrom: " ", pulse: LOW_PULSE}}

	for len(presses) > 0 {
		// pop the next item
		nextPress := presses[0]
		presses = presses[1:]

		// determine if in mapping
		mod, ok := modules[nextPress.modTo]
		if !ok {
			continue
		}

		// send the pulse
		outPulse := mod.receive(nextPress.pulse, nextPress.modFrom, highCount, lowCount)

		if outPulse == NO_PULSE {
			continue
		}

		for _, outMod := range mod.modList() {
			presses = append(presses, modPulse{modTo: outMod, modFrom: nextPress.modTo, pulse: outPulse})
		}
	}
	return modules
}

func pressButtonPart2(modules map[string]module, feederMod string, feederMap map[string]int, count int) map[string]module {

	presses := []modPulse{{modTo: "broadcaster", modFrom: " ", pulse: LOW_PULSE}}

	for len(presses) > 0 {
		// pop the next item
		nextPress := presses[0]
		presses = presses[1:]

		// determine if in mapping
		mod, ok := modules[nextPress.modTo]
		if !ok {
			continue
		}

		// send the pulse
        a, b := 0, 0 // don't care about high/low counts anymore
		outPulse := mod.receive(nextPress.pulse, nextPress.modFrom, &a, &b)

		if outPulse == NO_PULSE {
			continue
		}

		for _, outMod := range mod.modList() {
			presses = append(presses, modPulse{modTo: outMod, modFrom: nextPress.modTo, pulse: outPulse})

            // part 2 check "cycles"
            if outMod == feederMod && outPulse == HIGH_PULSE {
                feederMap[nextPress.modTo] = count
            }
		}
	}
	return modules
}
