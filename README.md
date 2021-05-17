# Daily Challenge
This repository was initially started to contain the code challenge solutions to the questions from *Daily Interview Pro* (https://www.techseries.dev/daily)

Later on it also started to host many solutions to problems on https://leetcode.com/problemset/all/

Leetcode nowadays has more than 2000 problems, it's challenging to solve them all. The pick of the problems to solve is heavily influenced by https://github.com/labuladong/fucking-algorithm 
### Code Structure
Each file is an executable program containing the *main* function. This makes it easy to test/run individual solution easily.
Some of the files do have the dependency on the _utils_ package which provides some basic data structures.    

### Run in IDE / CLI
I use IntelliJ IDEA to do edits. IntelliJ recognizes the go modules and packages nicely. Click on each file and select _Run 'go Build...'_ and it runs the program.

Alternatively, you can also go into this project root directory in CLI and run 
> go build --work you_chose_of_file.go 

Go automatically detects the dependency package and builds the file into an executable in the current directory. Then just run the executable. (IntelliJ does this better as it puts the build outputs into a temporary directory)

### Others
Not that it matters really, but here is the environment variables I use in my Mac:
> export GOPATH=~/go
>
> export GOBIN=~/go/bin