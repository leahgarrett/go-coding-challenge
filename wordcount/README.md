# Word Count

Can you improve the performance of this word count implementation?  

A simple timer has been used to track execution time.  

<br />  


# Input file

Provide the name of the text file as a command line argument.

`go run main.go sherlock.txt`  

Example file used:
https://www.gutenberg.org/ebooks/100


<br />  


# Slides for further background and description
https://docs.google.com/presentation/d/1M49tQDJVtRIbB-Do75S0_9LhUmt6Zv1q2SBE-kcOd84/edit?usp=sharing  


<br />  

# Some ideas from the discussion
Discussion at August Meetup: https://www.meetup.com/golang-mel/events/299861478/  

Suggestions  
- use go's benchmarking to look at the versions side by side
- could look at more efficient file loading
- can we simplify the code in the loop