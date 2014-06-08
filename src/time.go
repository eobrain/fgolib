package time
import type (
	java.util.Date
)

var Millisecond = 1000000

func Now() {
	new Date()
}

func Sleep(nanosecs) {
	Thread::sleep(nanosecs/Millisecond)
}
