package main

import "io/ioutil"
import "os"
import "strconv"
import "strings"
import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readint(f string) int64 {
	str, err := ioutil.ReadFile(f)
	check(err)
	i, err := strconv.ParseInt(strings.Trim(string(str), "\n "), 10, 0)
	check(err)
	return i
}

func main() {
	max_brightness := readint("/sys/class/backlight/intel_backlight/max_brightness")
	cur_brightness := readint("/sys/class/backlight/intel_backlight/brightness")
	if len(os.Args) == 1 {
		fmt.Printf("%.0f\n", float64(cur_brightness)/float64(max_brightness)*100)
	} else if len(os.Args) == 2 {
		arg, err := strconv.ParseFloat(strings.Trim(os.Args[1], "\n "), 64)
		check(err)
		new_brightness := int(float64(cur_brightness) + ((arg / 100.0) * float64(max_brightness)))
		if new_brightness > int(max_brightness) {
			new_brightness = int(max_brightness)
		} else if new_brightness < 0 {
			new_brightness = 0
		}
		f, err := os.OpenFile("/sys/class/backlight/intel_backlight/brightness", os.O_WRONLY, 0644)
		defer f.Close()
		check(err)
		fmt.Fprintf(f, "%d\n", new_brightness)
	} else if len(os.Args) == 3 {
		arg, err := strconv.ParseFloat(strings.Trim(os.Args[2], "\n "), 64)
		check(err)
		new_brightness := int((arg / 100.0) * float64(max_brightness))
		f, err := os.OpenFile("/sys/class/backlight/intel_backlight/brightness", os.O_WRONLY, 0644)
		defer f.Close()
		check(err)
		fmt.Fprintf(f, "%d\n", int(new_brightness))
	}
}
