// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type emailDomain struct {
	domain string
	count  int
}

// EmailDomainReport represents a sorted slice containing email domains count.
type EmailDomainReport []emailDomain

func (report EmailDomainReport) Len() int           { return len(report) }
func (report EmailDomainReport) Less(i, j int) bool { return report[i].count < report[j].count }
func (report EmailDomainReport) Swap(i, j int)      { report[i], report[j] = report[j], report[i] }

// NewEmailDomainReportFromMap returns a pointer to EmailDomainReport.
// It allows you to create the report from a given map.
// Input sample: map[gmail.com:10 hotmail.com:7]
func NewEmailDomainReportFromMap(mapReport map[string]int) *EmailDomainReport {
	report := EmailDomainReport{}

	for domain, count := range mapReport {
		emailDomain := emailDomain{
			domain: domain,
			count:  count,
		}

		report = append(report, emailDomain)
	}

	sort.Sort(report)
	return &report
}

// ImportFromCSV concurrently reads csv file located in p
// and returns EmailDomainReport.
//
// It uses two goroutines:
// - One reads the rows of the csv and emit them through rowChan.
// - Second one parse a given csv row and tries to get email domain.
func ImportFromCSV(p string) (*EmailDomainReport, error) {
	reportMap := make(map[string]int)

	rowChan := make(chan string)
	domainChan := make(chan string)

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	// Read csv goroutine
	// It reads row by row
	go func() {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			rowChan <- scanner.Text()
		}

		close(rowChan)
		// send signal close
		domainChan <- ""
	}()

	// Extract email domain goroutine
	// Input sample: Anna,Rivera,ariverag@whitehouse.gov,Female,105.158.80.2
	// Output sample: whitehouse.gov
	go func() {
		for {
			row := <-rowChan

			// signal close received
			if row == "" {
				break
			}

			// Get row columns
			data := strings.Split(row, ",")
			if len(data) == 5 {
				email, err := NewEmail(data[2])

				if err != nil {
					fmt.Println(err)
					continue
				}

				domain, err := email.Domain()
				if err != nil {
					fmt.Println(err)
					continue
				}

				domainChan <- domain
			}
		}

		close(domainChan)
	}()

	for domain := range domainChan {
		count, ok := reportMap[domain]
		if !ok {
			reportMap[domain] = 1
		} else {
			reportMap[domain] = count + 1
		}
	}

	return NewEmailDomainReportFromMap(reportMap), nil
}
