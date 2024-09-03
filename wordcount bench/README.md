# Word Count

Can you improve the performance of this word count implementation?  

A simple timer has been used to track execution time.  

# Input file

Provide the name of the text file as a command line argument.

`go run main.go shakespeare.txt`  

Example file to use:
https://www.gutenberg.org/ebooks/100


# Sample execution time

`go run main.go shakespeare.txt`

```
ReadFile: "shakespeare.txt" 5442125 bytes
countWords: 966501 words, duration: 54305ns
byteCountWords: 741200 words, duration: 2414366ns
byteCountWordsV2: 0 words, duration: 2414414ns
byteCountWordsV3: 966503 words, duration: 18486ns
```

# To run benchmarks

`go test -bench=.`  

Sample Ouput  

```
BenchmarkCountWords-12                         1        2118500458 ns/op
BenchmarkByteCountWordsV2-12                   1        4715139125 ns/op
BenchmarkByteCountWordsV3-12                   1        1141539625 ns/op
```