# About
A small Cli-Utility that will calculate for you the number of lines in all files and directories starting from your current directory.
You can use `alias countlines='pwd && echo "Total number of lines: $(cat $(fd -t file) | wc -l)"'`, and it will be even faster to count lines in all files, but my program has some svistoperdelki, for example:
- you can only count lines in files that match the desired extension.
- the program can display a table, list or tree of files with the number of lines and their size.
- by default hidden files are not taken into account by the program, but the `-h` flag will fix this.

For usage see `--help` flag.

# Installation
Build binary and copy it to your /usr/bin for linux
```sh
git clone https://github.com/PutaMadre1337/CountLines && cd CountLines && go build -o countlines $(find cmd/main.go) && sudo cp countlines /usr/bin && cd .. && rm -rf CountLines
```

# Configuration
```sh
export DEFAULT_FORM="table"
export ALLWAYS_DISPLAY_SIZE="false"
export ALLWAYS_SHOW_HIDDEN_FILES="false"
export LIST_ENUMERATOR="Roman"
export TREE_ENUMERATOR="RoundedEnumerator"
```
