package main

import "fmt"

func main(){
    // make function allocates and initializes the map
    // map[key]value -- this defines the acceptable types for a maps
    // ex: map[string]int -- the key is a string and the value is an int in this case

    states := make(map[string]string)
    fmt.Println(states)

    states["WA"] = "Washington"
    states["OR"] = "Oregon"
    states["CA"] = "California"
    fmt.Println(states)

    // you can retrieve item based on its key
    california := states["CA"]
    fmt.Println(california)

    // to delete an item from the map, use the built-in delete function
    delete(states, "OR")
    fmt.Println(states)

    states["NY"] = "New York"

    // this loops through the map, assigns the two variables
    // to the key and value of the current entry
    // maps iteration order is not guaranteed -- the last item added might appear first after this for loop
    for k, v := range states {
        fmt.Printf("%v: %v\n", k, v)
    }

}
