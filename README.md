# chirp2bc125

Convert a Chirp csv file to bc125 csv file so you can import it on a Bearcat scanner.

Chirp <https://chirp.danplanet.com/projects/chirp/wiki/Home>

bc125 <https://github.com/fdev/bc125csv>

## Command flags

```bash
./chirp2bc125.py                                                   
usage: chirp2bc125.py [-h] --input INPUT_FILE --output OUTPUT_FILE [--start-number START_NUMBER]
chirp2bc125.py: error: the following arguments are required: --input, --output
```

## Example conversion

```bash
./chirp2bc125.py --input PA3RD_repeaterlist.csv --output PA3RD_repeaterlist_bc125.csv
```
