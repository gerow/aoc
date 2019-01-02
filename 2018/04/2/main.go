package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type EventType int

const (
	ShiftChange EventType = iota
	Awake
	Asleep
)

type Event struct {
	Time    time.Time
	Type    EventType
	GuardID GuardID
}

type GuardID int

type Guard struct {
	ID            GuardID
	MinutesAsleep int
	TimesAsleep   [60]int
}

const TimeLayout = "2006-01-02 15:04"

var re = regexp.MustCompile(`^\[([^\]]*)\] (.*)$`)

func parse(s string) *Event {
	m := re.FindStringSubmatch(s)
	if m == nil {
		log.Fatalf("failed to parse event %q")
	}
	ds, es := m[1], m[2]

	var e Event
	var err error
	if e.Time, err = time.Parse(TimeLayout, ds); err != nil {
		log.Fatalf("failed to parse time %q: %v", ds, err)
	}

	prefix := strings.Fields(es)[0]
	switch prefix {
	case "falls":
		e.Type = Asleep
	case "wakes":
		e.Type = Awake
	case "Guard":
		e.Type = ShiftChange
		if _, err := fmt.Sscanf(es, "Guard #%d begins shift", &e.GuardID); err != nil {
			log.Fatalf("failed to parse guard from event %q: %v", es, err)
		}
	default:
		log.Fatalf("unknown event prefix %s", prefix)
	}

	return &e
}

func main() {
	var events []*Event
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		e := parse(s.Text())
		log.Printf("parsed event %+v", e)
		events = append(events, e)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}
	// order our events chronologically
	log.Print("sorting events chronologically")
	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})
	log.Print("sorted")
	for _, e := range events {
		log.Printf("%+v", e)
	}

	// play back the events and keep track of times asleep
	guards := make(map[GuardID]*Guard)

	var currentGuard *Guard
	var sleepTime int
	for _, e := range events {
		switch e.Type {
		case ShiftChange:
			if _, ok := guards[e.GuardID]; !ok {
				guards[e.GuardID] = &Guard{ID: e.GuardID}
			}
			currentGuard = guards[e.GuardID]
		case Awake:
			currentGuard.MinutesAsleep += e.Time.Minute() - sleepTime
			for i := sleepTime; i < e.Time.Minute(); i++ {
				currentGuard.TimesAsleep[i] += 1
			}
		case Asleep:
			sleepTime = e.Time.Minute()
		default:
			log.Fatalf("unknown event type %v", e.Type)
		}
	}

	log.Printf("%+v", guards)

	best := &Guard{}
	var bestMinute int
	var nAsleepAtTime int
	for _, g := range guards {
		for m, n := range g.TimesAsleep {
			if n > nAsleepAtTime {
				best = g
				bestMinute = m
				nAsleepAtTime = n
			}
		}
	}
	log.Printf("guard %d slept a total of %d minutes, best minute was %d during which they slept %d total minutes",
		best.ID, best.MinutesAsleep, bestMinute, nAsleepAtTime)
	fmt.Println(int(best.ID) * bestMinute)
}
