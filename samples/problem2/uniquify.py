#!/usr/bin/python
""" Program to read a file line by line and Remove duplicates
and write back to another file.

Approach Used:
1. Read input file one line at a time
2. Check for line in unique_list:
	a. If line present, move to next line
	b. If not present, write line to output file, add line to unique_list
"""
from optparse import OptionParser

class Uniquify:

	def __init__(self, options):
		""" Using a dictionary to store unique lines as key,
		    key lookup in dictionary is O(1)
		"""
		self.unique_line_dict = {} # Storing all unique lines as key
		self.inputfile = options.filename
		self.outputfile = options.output_filename
		self.verbose = options.verbose
		self.prepare_files()
		self.create_unique_line_dict()
		self.write_to_file()

	def prepare_files(self):
		# open file descriptors
		self.infd = open(self.inputfile, "r")
		self.outfd = open(self.outputfile, "w")

	def create_unique_line_dict(self):
		# Open inputfile, check for uniqueness, add to unique_line_dict
	    for lines in self.infd:
			# Replacing "\r\n" or "\r" by "\n" and treating them as separated \n
            # terminated lines
			lines = lines.replace("\r", "\n").split("\n")
			for line in lines:
			     if len(line) != 0 and not self.unique_line_dict.has_key(line):
			         self.unique_line_dict[line] = True

	def write_to_file(self):
		# write contents to outputfile
		for line in self.unique_line_dict:
			if self.verbose:
			     print "Writing: %s\n"%line
			self.outfd.write(line + "\n")

	def __del__(self):
		# Destructor, close file descriptors
		self.infd.close()
		self.outfd.close()

if __name__ == "__main__":
	
	usage = "usage: %prog --file=<filename> --output=<output-filename> [-v|--verbose]"
	parser = OptionParser(usage)
	parser.add_option("--file", dest="filename", help="input filename", metavar="FILE")
	parser.add_option("--output", dest="output_filename", help="output filename to write to", metavar="FILE")
	parser.add_option("-v", "--verbose", action="store_true", dest="verbose", default=False, help="Write to stdout")
	(options, args) = parser.parse_args()

	if not options.filename or not options.output_filename:
		parser.error("Incorrect Usage")
	uniq = Uniquify(options)
