package main
import (
	"flag"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
	"strings"
)

var (
	morse = map[string]string{}
	pin = rpio.Pin(3) // CHANGE GPIO PIN!!
	unit int
)

func init() {
	morse["a"] = "._"
	morse["b"] = "_..."
	morse["c"] = "_._."
	morse["d"] = "_.."
	morse["e"] = "."
	morse["f"] = ".._."
	morse["g"] = "__."
	morse["h"] = "...."
	morse["i"] = ".."
	morse["j"] = ".___"
	morse["k"] = "_._"
	morse["l"] = "._.."
	morse["m"] = "__"
	morse["n"] = "_."
	morse["o"] = "___"
	morse["p"] = ".__."
	morse["q"] = "__._"
	morse["r"] = "._."
	morse["s"] = "..."
	morse["t"] = "_"
	morse["u"] = ".._"
	morse["v"] = "..._"
	morse["w"] = ".__"
	morse["x"] = "_.._"
	morse["y"] = "_.__"
	morse["z"] = "__.."
}

func main() {
	message := flag.String("m", "sos sos", "Message to encode")
	verbose := flag.Bool("v", false, "Enable logging")
	flag.IntVar(&unit,"u",200, "Time unit in millisecond")
	flag.Parse()
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	pin.Output()
	defer rpio.Close()
	words := strings.Split(*message, " ")
	for i,_ := range words {
		if *verbose {
			log.Printf("enconding %s\n", words[1])
		}
		sing(words[i])
		pin.Low()
		if *verbose {
			log.Printf("pause between words\n")
		}
		time.Sleep( time.Duration(7 * unit) * time.Millisecond)
	}
}

func sing(msg string) {
	pin.Low()
	x := 0
	for _,l := range msg {
		//fmt.Printf("%s\n",morse[string(l)])
		for _,e := range morse[string(l)] {
			if string(e) == "." {
				x=1
			}
			if string(e) == "_" {
				x=3
			}
			pin.High()
			time.Sleep( time.Duration(x*unit) * time.Millisecond)
			pin.Low()
			time.Sleep( time.Duration(unit) * time.Millisecond)
		}
		pin.Low()
		time.Sleep( time.Duration(3*unit) * time.Millisecond)
	}
	pin.Low()
}
	morse["l"] = "._.."
	morse["m"] = "__"
	morse["n"] = "_."
	morse["o"] = "___"
	morse["p"] = ".__."
	morse["q"] = "__._"
	morse["r"] = "._."
	morse["s"] = "..."
	morse["t"] = "_"
	morse["u"] = ".._"
	morse["v"] = "..._"
	morse["w"] = ".__"
	morse["x"] = "_.._"
	morse["y"] = "_.__"
	morse["z"] = "__.."
}

func main() {
	msg := "sos"
	sing(msg)
}

func sing(msg string) {
	pin := rpio.Pin(3)
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	pin.Output()
	defer rpio.Close()
	pin.Low()
	x := 0
	for _,l := range msg {
		fmt.Printf("%s\n",morse[string(l)])
		for _,e := range morse[string(l)] {
			if string(e) == "." {
				x=1
			}
			if string(e) == "_" {
				x=3
			}
			pin.High()
			time.Sleep( time.Duration(x) * time.Second)
			pin.Low()
			time.Sleep( 1 * time.Second)
		}
		pin.Low()
		time.Sleep( 3 * time.Second)
	}
	pin.Low()

}
