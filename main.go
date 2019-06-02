package golang_lesson1

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

const ntpServer = "0.pool.ntp.org"

func PrintTime() error {
	ntpTime, err := ntp.Time(ntpServer)

	if err != nil {
		return fmt.Errorf("error getting time from ntp server \"%s\": %s", ntpServer, err.Error())
	}

	fmt.Println("Local time: ", time.Now())
	fmt.Println("Ntp time: ", ntpTime)

	return nil
}
