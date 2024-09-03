# Word Count

Can you improve the performance of this word count implementation?  

A simple timer has been used to track execution time.  

This is an updated version of [this earlier version](../wordcount/README.md). This version has been refactored with code from branches [of the base version](../wordcount) included as separate functions. These functions have had benchmark tests implemented.

<br />  


# Input file

Provide the name of the text file as a command line argument.

`go run main.go shakespeare.txt`  

Example file used:
https://www.gutenberg.org/ebooks/100

<br />  



# Creating a large input file

Making a file that comprises of 64 x `shakespeare.txt` in a file name which has been added to `.gitignore`
`cat shakespeare.txt shakespeare.txt shakespeare.txt shakespeare.txt > s.txt`
`cat s.txt s.txt s.txt s.txt > test.txt`    
`cat test.txt test.txt test.txt test.txt > test_data.txt`


<br />  


# Sample execution time

`go run main.go shakespeare.txt`

```
ReadFile: "shakespeare.txt" 5442125 bytes
countWords: 966501 words, duration: 55374ns
byteCountWords: 741200 words, duration: 2372541ns
byteCountWordsV2: 741200 words, duration: 2450428ns
byteCountWordsV3: 966503 words, duration: 18934ns
```
Comparing this with the command line word count tool:  
`time wc shakespeare.txt `  

Output:  
```
  196389  966503 5442125 shakespeare.txt
wc shakespeare.txt  0.02s user 0.00s system 87% cpu 0.028 total
```

<br />  

# To run benchmarks

`go test -bench=.`  

Sample Output  

```
BenchmarkCountWords-12                         1        2118500458 ns/op
BenchmarkByteCountWordsV2-12                   1        4715139125 ns/op
BenchmarkByteCountWordsV3-12                   1        1141539625 ns/op
```

<br />  


# Slides for further background and description

https://docs.google.com/presentation/d/18cKx35_qqlfPLHyAkSBOAgg4tD05jfS7uptieK8tsxA/edit?usp=sharing

<br />  


# Some ideas from the discussion
Disucssion at September Meetup:  
https://www.meetup.com/golang-mel/events/301991465/


Author's concerns
- some refactoring to move code into benchmark tests but still need further testing

Concerns
- why word counts are different? Between versions and `wc`

Suggestions include 
- use go's benchmarking to look at the versions side by side
- could look at more efficient file loading
- how about running the tests on a smaller file multiple times to more accurately represent performance
- can we simplify the code in the loop
- can we use code from go's implementation of `strings.Fields` to detect whitespace (bit wise operators would be more efficient than using the `if`)