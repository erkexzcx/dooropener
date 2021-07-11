package dooropener

import (
	"fmt"
	"os"
	"time"
)

func openDoor() {
	for i := 0; i < servoConfig.Pushes; i++ {
		setAngle(servoConfig.PushedAngle)
		time.Sleep(servoConfig.PushedWait)

		setAngle(servoConfig.ReleasedAngle)
		time.Sleep(servoConfig.ReleasedWait)
	}
	setAngle(servoConfig.AngleInactive)
}

// Move servo to requested angle
func setAngle(angle int) {
	f, err := os.Create("/dev/pi-blaster")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d=%f\n", servoConfig.Pin, angleToServo(angle)))
	f.Sync()
}

// helper function - converts angle to servo value (aka 'map' function)
func angleToServo(val int) float64 {
	return (float64(val)-0)*(0.25-0.05)/(180-0) + 0.05
}
