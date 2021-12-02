package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	
	"filtercsv/filtercsv"
	"filtercsv/models"
	"filtercsv/utils"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(`filtercsv, v%v | (c) 2021 tpkn.me

Usage:
  cat source.csv | filtercsv [-options] > filtered.csv

Options:
  -i                   Source csv file path (if you are not satisfied with pipes)
  -d                   Fields delimiter (default: ",")
  -f, --filters        File with a list of filters
  -c, --column         Column index (starting from 1)
  -h, --skip-header    Keep header intact (default: false)
  -e, --exclude        Inversed filtering (default: false)
  --help               Help
  --version            Print current version number
  --cpu                Max CPU cores (default: max)
`, version)

func main() {
	var cli = models.CLI{}
	flag.StringVar(&cli.InputPath, "i", "", "Source csv file path")
	flag.StringVar(&cli.Delimiter, "d", ",", "Fields delimiter")
	flag.StringVar(&cli.FiltersPath, "filters", "", "File with a list of filters (one line - one filter)")
	flag.StringVar(&cli.ColumnIndex, "column", "", "Source file column index (starting from 1)")
	flag.BoolVar(&cli.SkipHeader, "skip-header", false, "Keep header intact")
	flag.BoolVar(&cli.Exclude, "exclude", false, "Exclude lines matching filters")
	flag.BoolVar(&cli.Help, "help", false, "Help")
	flag.BoolVar(&cli.Version, "version", false, "Print current version number")
	flag.IntVar(&cli.CPU, "cpu", runtime.NumCPU(), "Max CPU cores")
	
	flag.StringVar(&cli.FiltersPath, "f", "", "Short alias for --filters")
	flag.StringVar(&cli.ColumnIndex, "c", "", "Short alias for --column")
	flag.BoolVar(&cli.SkipHeader, "h", false, "Short alias for --skip-header")
	flag.BoolVar(&cli.Exclude, "e", false, "Short alias for --exclude")
	flag.Parse()
	
	runtime.GOMAXPROCS(cli.CPU)
	
	log.SetFlags(0)
	log.SetPrefix("Error: ")
	
	if cli.Help {
		fmt.Println(help)
		os.Exit(0)
	}
	
	if cli.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	
	if !utils.FileExists(cli.FiltersPath) {
		log.Fatalln("file with filters does not exist")
	}
	if cli.ColumnIndex == "" {
		log.Fatalln("column index is not set")
	}
	
	filtercsv.Run(cli)
}
