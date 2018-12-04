package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type guard struct {
	totalSleptMinutes        int
	minuteMap                map[string]int
	sleepiestMinute          string
	sleepTimeSleepiestMinute int
}

func (g *guard) registerSleep(fallAsleepTime time.Time, wakeUpTime time.Time) {
	for i := fallAsleepTime; i.Before(wakeUpTime); i = i.Add(time.Minute) {
		strMinute := i.Format("15:04")
		g.minuteMap[strMinute]++
		g.totalSleptMinutes++
	}
}

func (g *guard) registerSleepiness() {
	var sleepiestMinute string
	maxSleep := 0
	for i := range g.minuteMap {
		if g.minuteMap[i] > maxSleep {
			sleepiestMinute = i
			maxSleep = g.minuteMap[i]
		}
	}
	g.sleepiestMinute = sleepiestMinute
	g.sleepTimeSleepiestMinute = maxSleep
}

type event struct {
	timestamp time.Time
	text      string
	eventType string // This should be an enum but palla, this has the valid values "startshift", "fallasleep" and "wakeup"
}

func createEventFromString(s string) event {
	var splitted []string
	splitted = strings.Split(s, "]")
	dateStr := splitted[0][1:]
	dateTimeStamp, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	var eventType string
	switch splitted[1][1:6] {
	case "Guard":
		eventType = "startshift"
	case "falls":
		eventType = "fallasleep"
	case "wakes":
		eventType = "wakeup"
	}

	return event{timestamp: dateTimeStamp, text: splitted[1][1:], eventType: eventType}

}

func parseEvents(e []event, g map[int]*guard) {
	var activeGuard int
	var fallAsleepTime time.Time
	// var wakeUpTime time.Time
	for _, event := range e {
		switch event.eventType {
		case "startshift":
			fmt.Sscanf(event.text, "Guard #%d begins shift", &activeGuard)
			if _, ok := g[activeGuard]; !ok {
				g[activeGuard] = &guard{minuteMap: make(map[string]int)}
			}
		case "fallasleep":
			fallAsleepTime = event.timestamp
		case "wakeup":
			g[activeGuard].registerSleep(fallAsleepTime, event.timestamp)

		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var events []event
	guards := make(map[int]*guard)
	for {
		if !scanner.Scan() {
			break
		}
		events = append(events, createEventFromString(scanner.Text()))
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].timestamp.Before(events[j].timestamp)
	})

	parseEvents(events, guards)

	// EXTRACTING SOLUTION FOR PART 1

	maxSleep := 0
	sleepyGuardID := 0
	for i, g := range guards {
		if g.totalSleptMinutes > maxSleep {
			maxSleep = g.totalSleptMinutes
			sleepyGuardID = i

		}
	}
	sleepyMinute := ""
	minutesSlept := 0
	for k := range guards[sleepyGuardID].minuteMap {
		if guards[sleepyGuardID].minuteMap[k] > minutesSlept {
			sleepyMinute = k
			minutesSlept = guards[sleepyGuardID].minuteMap[k]
		}

	}
	fmt.Printf("A: Guard #%d Slept most. S?he slept %d times at %v\n", sleepyGuardID, minutesSlept, sleepyMinute)

	// EXTRACTING SOLUTION FOR PART 2
	for _, g := range guards {
		g.registerSleepiness()
	}
	var maxSleepTime int
	var worstGuard int
	for i, g := range guards {
		if g.sleepTimeSleepiestMinute > maxSleepTime {
			maxSleepTime = g.sleepTimeSleepiestMinute
			worstGuard = i
		}
	}
	fmt.Printf("B: Guard #%d , slept %d times, during %v\n", worstGuard, maxSleepTime, guards[worstGuard].sleepiestMinute)

}
