package filtercsv

import (
	"testing"
	
	"filtercsv/models"
)

var options = models.CLI{
	InputPath:   "./../_/test.csv",
	FiltersPath: "./../_/test_hashify.csv",
	ColumnIndex: "1",
	SkipHeader:  true,
	Exclude:     true,
}

func Test_Run(t *testing.T) {
	Run(options)
}
