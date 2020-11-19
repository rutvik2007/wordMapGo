## CS476 Assignment 2

Alejandro Servetto	aservet1@binghamton.edu

Rutvik Saptarshi	rsaptar2@binghamton.edu

Data collected by running the program with the different configurations on all the pg\*.txt files. USed a bash 
script to facilitate all that. Inlcuded in the tar.gz file runall.sh. Redirect stderr to your desired output file and you
will get a dump of the timing results (in different format than the table below. I reformatted that for this writeup)

See Bash Script: ./runall.sh

1) Time your results. Time results with settings:
	
	trial	askdelay	readers	askers	time(s)
	1  	10	1	1	0m20.799s
	2	10	16	2	0m35.042s
	3	10	4	8	0m34.195s
	4	10	16	32	11m18.747s
	5	10	64	64	61m33.074s <-- this is the time that trial 5 was cut off because it was taking way too long


2) Discuss your results

		It looks like a lot of readers and askers lead to horirbly long run times, as the 61 minutes listed above is
	the time that we interrupted the program to force it to finish. It was taking too long. Here there were more readers
	and askers that were necessary to read the amount of files available. So we had too many extra go routines running and
	the channels were constantly full and blocked, waiting to be used. The overhead of having so many readers and askers
	was too much for the requirements of this program, so it ended up just causing a lot of overhead for a lot of concurrent
	components, which was just a huge time cost. Waiting and context switch overheads, stuff like that. Too much overhead
	that just got in the way of each other.

	There seems to be a direct positive correlation between number of askers and number of readers.

	This implementation does not scale well with more readers and askers. You need to pick the appropriate balance
	of readers and askers for the CPU hardware you have (to assess what level of concurrency it can support) and
	for the size and amount of data you're using. You also have to configure your channel buffer sizes to work
	with these parameters as well to avoid too-often having to wait on full buffers, as that adds extra overhead.



3)
	8 core laptop

