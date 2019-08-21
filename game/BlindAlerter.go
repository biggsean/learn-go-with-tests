package poker

import "fmt"
import "os"
import "time"

// BlindAlerter is an interface to schedule an alert to raise the blinds
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc implements the BlindAlerter interface
type BlindAlerterFunc func(duration time.Duration, amoount int)

// ScheduleAlertAt method for the BlindAlerterFunc
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter is a implementation of BlindAlerter
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
