#!/bin/bash 

files=$(ls data | grep pg ) #| sed s/[^0-9]//g )

files_csv=''
for file in $files
do
	#echo $file
	files_csv="$files_csv,./data/$file"
done

files_csv=$(echo $files_csv | sed 's/,//') #remove beginning comma
#echo $files_csv; exit

go run emerging.go cmap.go -chan -infiles="$files_csv" -readers=10 -askers=2 -askdelay=10 
