package main

import (
	"flag"
	"fmt"
	"time"
	"xuan/test2/go_library/utils/files"
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
	case mode == "wm":
		d = d.AddDate(0, 0, -2)
	case mode == "zd":
		d = d.AddDate(0, 0, -1)
	}
	return
}

func init() {
	//parse command line
	flag.Parse()
	//check is any task is called, use nil instead "" because falgMode is pointer
	if flagMode != nil {
		date = defaultDate(*flagMode)
	}
	//calcualte the startDate
	date = date.AddDate(0, 0, -1**flagOffset)
	delta = *flagDelta

}
func main() {
	if flagMode != nil {
		switch {
		case *flagMode == "":
			fmt.Println("No Mode")
		case *flagMode == "env":
			fmt.Println(date)
			fmt.Println(delta)
		case *flagMode == "files":
			files.ParseFile()
		}
	}
}
