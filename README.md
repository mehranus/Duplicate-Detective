# Duplicate Detective

Duplicate Detective is a simple command-line tool that finds duplicate files in a directory. It uses the MD5 hashing algorithm to compute a unique hash value for each file in the directory. If two files have the same hash value, Duplicate Detective assumes that they have the same contents and are therefore duplicates.
there is also a step-to-step guide blog post [here](https://mehranbehnam77.medium.com/duplicate-detective-a-step-by-step-guide-to-identifying-duplicate-files-with-golang-d1fdca5149f6)  

## Usage

To use Duplicate Detective, run the program with a single command-line argument specifying the directory to search:


```$ duplicatedetective /path/to/directory```

Duplicate Detective will then search the specified directory and its subdirectories for duplicate files, and print a message for each duplicate file it finds.

## Installation

To install Duplicate Detective, you need to have Go installed on your system. Then, simply run:


```$ go get github.com/mehrunus/duplicate-detective```

You can then run Duplicate Detective from the command line as described above.

## License

Duplicate Detective is licensed under the MIT License. See LICENSE for more information.