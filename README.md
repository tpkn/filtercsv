# filtercsv


CLI filters the contents of the csv file according to the filters from the another file. Made to process big files by a lots of filters.

By default, the results include lines that match the filters. Add `--exclude` flag to exclude lines matching filters.

```text
+---------------------+---------------+------------------+
| <input.csv>         | <filters.csv> | <output>         |
+---------------------+---------------+------------------+
| ID,City             | 901007        | ID,City          |
| 901023,Chelyabinsk  | 901055        | 901007,Habarovsk |
| 901019,Ekaterinburg | 999999        | 901055,Izhevsk   |
| 901007,Habarovsk    |               |                  |
| 901055,Izhevsk      |               |                  |
| 901009,Kaliningrad  |               |                  |
| 901035,Krasnoyarsk  |               |                  |
+---------------------+---------------+------------------+
```

## Usage

```shell
cat <input.csv> | filtercsv [-options] > <output>
```



## Options

```text
  -i                   Source csv file path (if you are not satisfied with pipes)
  -d                   Fields delimiter (default: ",")
  -f, --filters        File with a list of filters
  -c, --column         Column index (starting from 1)
  -h, --skip-header    Keep header intact (default: false)
  -e, --exclude        Inversed filtering (default: false)
  --help               Help
  --cpu                Max CPU cores (default: max)
  -v, --version        Print current version number
```
