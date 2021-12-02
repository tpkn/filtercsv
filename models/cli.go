package models

type CLI struct {
	InputPath   string
	Delimiter   string
	FiltersPath string
	ColumnIndex string
	SkipHeader  bool
	Exclude     bool
	
	Help    bool
	Version bool
	CPU     int
}
