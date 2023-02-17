package customerimporter

import (
	"testing"
)

func TestImportFromCSV(t *testing.T) {
	t.Run("Should read file and generated sorted report", func(t *testing.T) {
		report, _ := ImportFromCSV("customers.csv")
		data := *report

		if data[0].count != 1 {
			t.Errorf("Report order is wrong")
		}

		if data[len(data)-1].count != 14 {
			t.Errorf("Report order is wrong")
		}
	})

	t.Run("Should NOT generated report if csv NOT found", func(t *testing.T) {
		_, err := ImportFromCSV("non_existing_file.csv")

		if err == nil {
			t.Errorf("Report should not be generated when file does not exist, cause %v", err)
		}
	})
}
