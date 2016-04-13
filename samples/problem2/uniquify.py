#!/usr/bin/python
from optparse import OptionParser
import hashlib

class Uniquify:

	def __init__(self, options):
		self.unique_hash_list = []
		self.inputfile = options.filename
		self.outputfile = options.output_filename
		self.verbose = options.verbose
		self.prepare_files()
		self.remove_duplicates()

	def prepare_files(self):
		# open file descriptors
		self.infd = open(self.inputfile, "r")
		self.outfd = open(self.outputfile, "w")

	def remove_duplicates(self):
		# Open inputfile, check for uniqueness, write to output
	    for line in self.infd:
			if not self.already_seen(line):
				self.write_to_output(line)

	def already_seen(self, line):
		# Calculate hash of line, check if already seen, return boolean
		hash_val = hashlib.md5(line).hexdigest()
		if hash_val in self.unique_hash_list:
			return True
		else:
			self.unique_hash_list.append(hash_val)
			return False

	def write_to_output(self, line):
		# write contents to outputfile
		if self.verbose:
			print "Writing: %s"%line
		self.outfd.write(line)

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
