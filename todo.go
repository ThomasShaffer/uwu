package main

import (
    "os/exec"
    "strings"
    "bytes"
)

type entry struct {
    file        string
    filepath    string
    priority    int
    message     string
    line        string
}

func NewEntry(filepath string, priority int, message string, line string) (entry) {
    filepathSplit := strings.Split(filepath, "/")
    return entry{
        file:   strings.TrimSpace(filepathSplit[len(filepathSplit)-1]),
        filepath: strings.TrimSpace(filepath),
        priority:   priority,
        message:    strings.TrimSpace(message), 
        line:       line,
    }

}

func GetTodos() ([]entry) {
    var out bytes.Buffer
    var result []entry
    grep := exec.Command("grep", "TODO*:", "-Hrni", ".")
    for dir := range config.IgnoreDir {
        grep.Args = append(grep.Args, "--exclude-dir="+config.IgnoreDir[dir])
    }
    for file := range config.IgnoreFile {
        grep.Args = append(grep.Args, "--exclude="+config.IgnoreFile[file])
    }
    grep.Stdout = &out
    grep.Run()
    split := strings.Split(string(out.Bytes()), "\n")
    split = split[:len(split)-1]
    for line := range split {
        colonSplit := strings.Split(split[line], ":")
        filepath := colonSplit[0]

        line := colonSplit[1]

        message := colonSplit[3]
        entry := NewEntry(filepath, 0, message, line)
        result = append(result, entry)
    }
    return result
}


func FormatTodos(todos []entry) []string {
    var result []string
    var maxFileLength int
    var maxLineLength int
    for entry := range todos {
        maxFileLength = max(maxFileLength, len(todos[entry].file))
        maxLineLength = max(maxLineLength, len(todos[entry].line))
    }

    for entry := range todos {
        todo := todos[entry]
        numFileSpaces := strings.Repeat(" ", maxFileLength - len(todo.file) + 1)
        numLineSpaces := strings.Repeat(" ", maxLineLength - len(todo.line) + 0)
        result = append(result, todo.file + numFileSpaces + ":" + todo.line + numLineSpaces + ": " + todo.message + "\n")
    }
    return result
}
