#!/bin/bash

USAGE="Usage: $(basename "$0") --create-filename=<filename> [--no-prompt] [--verbose]"

INVALID=YES #Show usage prompt
VERBOSE=NO #Default
for i in "$@"
do
case $i in
	--create-file=*)
	FILENAME="${i#*=}"
    INVALID=NO
	shift
	;;
	--no-prompt)
	NO_PROMPT=YES
	shift
	;;
	--verbose)
	VERBOSE=YES
	shift
	;;
	--help|-h)
	HELP=YES
    INVALID=NO
	shift
	;;
	*)

	;;
esac
done

#Check for Invalid command
if [ "$INVALID" = "YES" ]; then
	echo "Invalid Command"
	echo $USAGE
	exit 1 
fi

#Check for Help command
if [ "$HELP" = "YES" ]; then
	echo $USAGE
	echo "where: "
	echo "<filename> - file to write the contents to"
	echo "--no-prompt - No prompt if <filename> exists"
	echo "--verbose - Verbose"
	exit 0
fi

# Hardcode all states to a list
STATES=( Alabama Alaska Arizona Arkansas California Colorado Connecticut Delaware Florida Georgia Hawaii Idaho Illinois Indiana Iowa Kansas Kentucky Louisiana Maine Maryland Massachusetts Michigan Minnesota Mississippi Missouri Montana Nebraska Nevada 'New Hampshire' 'New Jersey' 'New Mexico' 'New York' 'North Carolina' 'North Dakota' Ohio Oklahoma Oregon Pennsylvania 'Rhode Island' 'South Carolina' 'South Dakota' Tennessee Texas Utah Vermont Virginia Washington 'West Virginia' Wisconsin Wyoming )

function write_to_file {
    # Function that writes to the output file
	if [ "$1" = "YES" ]; then
		echo "File removed"
	fi

	if [ "$VERBOSE" = "YES" ]; then
	    echo "New file created"
    fi

	( IFS=$'\n'; echo "${STATES[*]}" ) > $FILENAME # Write states to file line by line
	exit 0
}

if [ -f $FILENAME ]; then
	if [ "$VERBOSE" = "YES" ]; then
		echo "File already exists"
	fi
	if [ "$NO_PROMPT" = "YES" ]; then
		# No prompt, overwrite
		write_to_file YES
	else
		while true; do # do this while response is not y/n 
			read -p "File exists. Overwrite (y/n) ?" prompt
			case $prompt in
				[y]) echo ; write_to_file YES;; # overwrite the file
				[n]) echo ; exit 0 ;; # do not overwrite, exit
				*) echo "Invalid input" # Invalid input
			esac
		done 
	fi
else
	write_to_file NO # File doesnt exist, create new	
fi

