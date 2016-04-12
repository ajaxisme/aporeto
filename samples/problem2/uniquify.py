#!/usr/bin/python
import optparse
import mmh3

class Uniquify:

	def __init__(self, inputfile, outputfile):
		self.unique_hash_list = []
		self.inputfile = inputfile
		self.outputfile = outputfile
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
		hash_val = mmh3.hash64(line)
		if hash_val in self.unique_hash_list:
			return True
		else:
			self.unique_hash_list.append(hash_val)
			return False

	def write_to_output(self, line):
		# write contents to outputfile
		self.outfd.write(line)

	def __del__(self):
		# Destructor, close file descriptors
		self.infd.close()
		self.outfd.close()

if __name__ == "__main__":
	uniq = Uniquify("big_file.txt", "output.txt")
