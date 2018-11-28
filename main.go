package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	date   time.Time //date of startDate, default is today
	delta  int       //window size
	offset int       //change the startDate
	mode   string    //pick which task

	flagMode   = flag.String("m", "", "decide pick which task") //*string
	flagDelta  = flag.Int("delta", 1, "window size")
	flagOffset = flag.Int("offset", 0, "change the startDate")
)

//naked function
func defaultDate(mode string) (d time.Time) {
	d = time.Now()
	switch {
	case mode == "mc.tp":
		d = d.AddDate(0, 0, -2)
		break
	case mode == "wm":
		d = d.AddDate(0, 0, -2)
		break
	case mode == "zd":
		d = d.AddDate(0, 0, -1)
		break
	}
	return
}

func init() {
	//parse command line
	flag.Parse()
	//check is any task is called, use nil instead "" because falgMode is pointer
	if flagMode != nil {
		date = defaultDate(*flagMode)
		fmt.Println(date)
	}
	//calcualte the startDate
	date = time.Now()

}
func main() {
	fmt.Println("running")
}
