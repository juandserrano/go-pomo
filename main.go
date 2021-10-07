package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const banner string = `
######                      ###
#     # ######  ####  ##### ### 
#     # #      #        #   ### 
######  #####   ####    #    #  
#   #   #           #   #       
#    #  #      #    #   #   ### 
#     # ######  ####    #   ###`

func main() {
	if len(os.Args) != 1 {
		if os.Args[1] == "start" {
			if len(os.Args) == 2 {
				fmt.Printf("Defaulting to 20 minutes")
				time.Sleep(500 * time.Millisecond)
				startTimer("20")
			} else {
				startTimer(os.Args[2])
			}
		}

	}
}

func startTimer(timeString string) {
	timeInMinutes, err := strconv.ParseFloat(timeString, 32)
	if err == nil {
		timer := timeInMinutes * 60
		for timer > 0 {
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Time left: %d:%s\n", int(timer/60), secondsFormatter(int(timer)%60))
			time.Sleep(1 * time.Second)
			timer--
		}

		i := 0
		fmt.Print("\a")
		for i < 200 {
			fmt.Print("\033[H\033[2J")
			fmt.Print(banner)
			time.Sleep(1000 * time.Millisecond)
			fmt.Print("\033[H\033[2J")
			time.Sleep(500 * time.Millisecond)
			i++
		}
	}
}

func secondsFormatter(seconds int) string {
	if seconds < 10 {
		return "0" + strconv.Itoa(seconds)
	}
	return strconv.Itoa(seconds)
}
