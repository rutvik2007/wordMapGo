#!/bin/bash 

files=$(ls data | grep pg ) 

for file in $files
do
	if [ "$files_csv" == '' ]
	then
		files_csv="./data/$file"
	else
		files_csv="$files_csv,./data/$file"
	fi
done

#echo $files_csv; exit

readerVals=(1 16 4 16 64)
askerVals=( 1  2 8 32 64)

for index in ${!readerVals[*]}
do
	rs=${readerVals[$index]}
	as=${askerVals[$index]}
	>&2 echo " --> STARTING TRIAL WITH $rs READERS, $as ASKERS"
	time go run emerging.go cmap.go -chan -infiles="$files_csv" \
		-readers=${readerVals[$index]} \
		-askers=${askerVals[$index]} \
		-askdelay=10\
		> /dev/null
	echo " <-- done with iteration $(( $index + 1 )) out of 5. readers=$rs ; askers=$as"
done
