#!/usr/bin/env bash
printout=""

if [ $# -eq 0 ] 
then
	exit 0
fi

if (( $1%3 == 0 )) 
then
	printout+="Pling"
fi

if (( $1%5 == 0 ))
then
	printout+="Plang"
fi
if (( $1%7 == 0 ))
then
	printout+="Plong"
fi

if [ "$printout" ] 
then
	echo $printout
	exit 0
fi
echo $1
