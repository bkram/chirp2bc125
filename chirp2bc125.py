#!/usr/bin/env python3
import argparse
import csv


def convert_chirp_to_bc125(input_file, output_file, start_number):
    with open(input_file, 'r') as input_csv, open(output_file, 'w',
                                                  newline='') as output_csv:
        csv_reader = csv.DictReader(input_csv)
        fieldnames = ["Channel", "Name", "Frequency", "Modulation",
                      "CTCSS/DCS", "Delay", "Lockout", "Priority"]
        csv_writer = csv.DictWriter(output_csv, fieldnames=fieldnames)
        csv_writer.writeheader()

        for i, row in enumerate(csv_reader, start=start_number):
            ctcss_dcs = convert_ctcss_dcs(row['Tone'], row['rToneFreq'])
            write_row_to_csv(csv_writer, str(
                i), row['Name'], row['Frequency'], row['Mode'], ctcss_dcs)


def convert_ctcss_dcs(tone, r_tone_freq):
    if tone in ['Tone', 'SQL'] and r_tone_freq:
        return f'{float(r_tone_freq):.1f} Hz'
    return 'none'


def write_row_to_csv(csv_writer, location, name, frequency, mode, ctcss_dcs):
    csv_writer.writerow({
        "Channel": location,
        "Name": name,
        "Frequency": frequency,
        "Modulation": mode,
        "CTCSS/DCS": ctcss_dcs,
        "Delay": "2",
        "Lockout": "no",
        "Priority": "no"
    })


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Convert Chirp CSV format to bc125 CSV format")

    parser.add_argument("--input", dest="input_file",
                        required=True,
                        help="Path to the input Chirp CSV file")
    parser.add_argument("--output", dest="output_file",
                        required=True,
                        help="Path to the output bc125 CSV file")
    parser.add_argument("--start-number", type=int, default=1,
                        help="Starting number for the first column")

    args = parser.parse_args()

    convert_chirp_to_bc125(
        args.input_file, args.output_file, args.start_number)
