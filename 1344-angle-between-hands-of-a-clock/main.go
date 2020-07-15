package main

import "math"

func angleClock(hour int, minutes int) float64 {
	h, m := float64(hour), float64(minutes)

	time := 60*h + m

	minuteAngle := 360.0 / 60.0
	hourAngle := minuteAngle / 12.0

	hoursAngle := hourAngle * time
	minutesAngle := minuteAngle * m

	angle := math.Abs(minutesAngle - hoursAngle)
	if angle > 180 {
		angle = 360.0 - angle
	}

	return angle
}
