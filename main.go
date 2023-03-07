// package
package main

// import
//
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// main
func main() {

	// Scanner to get the input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")
	fmt.Println("Enter mail address:")

	// While loop
	for scanner.Scan() {
		address := scanner.Text()
		domain := strings.Split(address, "@")[1]
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	if !hasSPF && !hasMX && !hasDMARC {
		fmt.Println("\n\nEmail service is not a vaild one")
	} else {
		fmt.Println("\n\nEmail service is a vaild one")
	}
}
