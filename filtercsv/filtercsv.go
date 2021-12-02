package filtercsv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
	
	"filtercsv/models"
	"filtercsv/utils"
)

func Run(o models.CLI) {
	column_index, err := strconv.Atoi(o.ColumnIndex)
	if err != nil {
		log.Fatalln("invalid column index")
	}
	
	delimiter, _ := utf8.DecodeRuneInString(o.Delimiter)
	column_index = column_index - 1
	
	// Making a hash map from --filter
	filters, err := utils.FileHashify(o.FiltersPath)
	if err != nil {
		log.Fatalln(err)
	}
	
	var is_header = true
	
	// Writer
	var writer = csv.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.Comma = delimiter
	
	// Reader
	var source *os.File
	if o.InputPath != "" && utils.FileExists(o.InputPath) {
		file, err := os.Open(o.InputPath)
		defer file.Close()
		if err != nil {
			log.Fatalln(err)
		}
		source = file
	} else {
		source = os.Stdin
	}
	
	reader := csv.NewReader(source)
	reader.Comma = delimiter
	reader.ReuseRecord = true
	reader.TrimLeadingSpace = true
	
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		
		// Keep header intact
		if is_header {
			is_header = false
			
			// Check if we have enough columns in the source csv file
			if len(line) <= column_index {
				log.Fatalln("there are not enough columns in the source csv")
			}
			
			if o.SkipHeader {
				writer.Write(line)
				continue
			}
		}
		
		_, found := filters[line[column_index]]
		if (o.Exclude && !found) || (!o.Exclude && found) {
			writer.Write(line)
		}
	}
}
