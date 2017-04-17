package main

import (
	"flag"
	"github.com/nathmclean/aws_metadata/meta"
	"net/http"
	"fmt"
	"log"
)

func metadata(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ami-id\nami-launch-indexami-manifest-path" +
		"\nblock-device-mapping/\nhostname\ninstance-action\ninstance-id" +
		"\ninstance-type\nkernel-id\nlocal-hostname\nlocal-ipv4\nmac" +
		"\nnetwork/\nplacement/\npublic-hostname\npublic-ipv4\npublic-keys/\nreservation-id" +
		"\nsecurity-groups\nservices/")
}

func blockDevice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path %s", r.URL.Path)
}

func main() {
	var portPTR = flag.String("port", "8080", "Port to listen on")
	//var proxy = flag.Bool("proxy", false, "Whether to act as a proxy - only for use on ec2 instances")
	//var netIntName = flag.String("interface", "aws0", "The name to give to the interface")
	flag.Parse()

	meta.Load()
	log.Println("Setting up Server on port " + *portPTR)

	http.HandleFunc("/latest/meta-data/", metadata)
	http.HandleFunc("/latest/meta-data/*", metadata)
	http.HandleFunc("/latest/meta-data/block-device-mapping/", blockDevice)
	http.ListenAndServe(":" + *portPTR , nil)
}