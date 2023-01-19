#!/bin/awk -f
{
if(NR==1){

	if (val=""){
		print "One for you, one for me." }
	else
		print $0

}
}
