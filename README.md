# aporeto

## Bash Script
1. Used a list to store all 50 states

## Uniquify - Python
#### Implementation
1. Read a line from the textfile
2. Generate a hash of the line
3. If hash already in hash_table: move to 4, else write the line to output, add hash to hash_table
4. Repeat 1-3 till all lines are read

## Word Count - Go
#### Implementation done for serial execution
1. Get list of urls from commandline
2. For each URL, do:
  1. Get data from URL
  2. Apply regex to data to extract all valid words
  3. Create a frequency table for all words
  4. Write to output file
