# Word Count

Can you improve the performance of this word count implementation?  

A simple timer has been used to track execution time.  

# Input file

Provide the name of the text file as a command line argument.

`go run main.go moby.txt`  
or  
`go run main.go shakespeare.txt`  

Example file to use:
https://www.gutenberg.org/cache/epub/2701/pg2701.txt


# To view profile

`go tool pprof -http=:8080 cpuprofile.pprof`  