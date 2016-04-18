# aporeto

## Bash Script
1. Used a list to store all 50 states

## Uniquify - Python
#### Implementation
#### For O(1) lookup and accuracy, using a dictionary to store lines as key
#### Only thing we will be compromising on would be the order which is not restored
1. Read a line from the input file
2. If line already in dictionary: move to 4, else add line as key to dictionary
3. Repeat 1-2 till all lines are read
4. Write all keys in the dictionary to output file

## Word Count - Go
#### Implementation done for serial execution
1. Get list of urls from commandline
2. For each URL, do:
  1. Get data from URL
  2. Apply regex to data to extract all valid words
  3. Create a frequency table for all words
  4. Write to output file
