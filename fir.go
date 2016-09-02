package firClient

import (
	"flag"
	"fmt"
)

const (
	libraryVersion = "0.1.0"
	userAgent      = "fir_client/" + libraryVersion
	mediaType      = "application/json"
)

func main() {
	hostPtr := flag.String("host", "localhost", "Hostname of the FIR instance")
	flag.Parse()
	fmt.Println("host:", *hostPtr)
}
