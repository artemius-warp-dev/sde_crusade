package main

import ( "fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.or")
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(time)
	

}
