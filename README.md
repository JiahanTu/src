# src
This script tool could audit all of the command in the log

## Run
To trigger the code, simply run:
```
go run my_ssh/script.go
```

Then the command history would be stored in command_results.txt

Hint: The GOPATH need to be src's parent repo, to set gopath, run: `export GOPATH=[paht/to/src's/parent/repo]`

## Examples
```
jiahan.tu@Jiahans-MacBook-Pro src % go run my_ssh/script.go
jiahan.tu@Jiahans-MacBook-Pro src % pwd
/Users/jiahan.tu/Documents/Project/go_project/src
jiahan.tu@Jiahans-MacBook-Pro src % cd my_ssh 
jiahan.tu@Jiahans-MacBook-Pro my_ssh % ls
main.go         pkg             pty             script.go       tools
jiahan.tu@Jiahans-MacBook-Pro my_ssh % cd ..
jiahan.tu@Jiahans-MacBook-Pro src % exit
jiahan.tu@Jiahans-MacBook-Pro src % cat command_results.txt 
jiahan.tu@Jiahans-MacBook-Pro src % pwd
/Users/jiahan.tu/Documents/Project/go_project/src
jiahan.tu@Jiahans-MacBook-Pro src % cd my_ssh 
jiahan.tu@Jiahans-MacBook-Pro my_ssh % ls
main.go         pkg             pty             script.go       tools
jiahan.tu@Jiahans-MacBook-Pro my_ssh % cd ..
jiahan.tu@Jiahans-MacBook-Pro src % exit
```