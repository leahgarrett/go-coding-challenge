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

For this branch with buffering:
`"shakespeare.txt": 741200 words, duration: 97406ms`


# To view profile

`go tool pprof -http=:8080 cpuprofile.pprof`  
