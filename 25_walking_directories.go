package main

import(
    "fmt"
    "os"
    "path/filepath"
)

func main(){

    // this returns the absolute path of a package by looking at
    // the relative path information
    root, _ := filepath.Abs(".")
    fmt.Println("Processing path:", root)

    // processPath will be called for each directory and each
    // file in the directory structure
    err := filepath.Walk(root, processPath) // processPath is a function which needs to have
                                            // a specific function signature to be recursively
                                            // called by the caller

    if err != nil {
        fmt.Println("error:", err)
    }

}

// filepath.Walk uses this function for each file and path it encounters with.
// It only accepts a function with a specific signature
// -- path --> specifies current directory or filename
// -- info --> specifies the the file information
// -- err  --> this is the error object
func processPath(path string, info os.FileInfo, err error) (error){

    // if the function receives an error from the Walk function
    // it return that error and it halts the calling Walk function
    if err != nil {
        return err
    }

    // "." is here to skip doing something
    // with the filesystem root
    if path != "." {
        if info.IsDir() {
            fmt.Println("Directory:", path)
        } else {
            fmt.Println("File:", path)
        }
    }

    return nil
}