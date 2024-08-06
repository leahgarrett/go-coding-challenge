# Word Count

Can you improve the performance of this word count implementation?  

A simple timer has been used to track execution time.  

# Input file

Provide the name of the text file as a command line argument.

`go run main.go shakespeare.txt`  

Example file to use:
https://www.gutenberg.org/ebooks/100


# Execution time

For base branch:
`"shakespeare.txt": 741200 words, duration: 3422897ms`

For this branch:
`"shakespeare.txt": 378586 words, duration: 30136ms`

Compare to word count:

`time wc -w shakespeare.txt `  

`378588 shakespeare.txt`    
`wc -w shakespeare.txt  0.01s user 0.00s system 76% cpu 0.016 total`  