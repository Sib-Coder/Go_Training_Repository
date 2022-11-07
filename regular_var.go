package main
 
import (
    "os"
    "regexp"
    "fmt"
)
 
 
func check(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
 
func main() {
    emails := []string{
        "brown@fox",
        "brown@fox.",
        "brown@fox.com",
        "br@own@fox.com",
    }
 
    pattern := `^\w+@\w+\.\w+$`
    for _, email := range emails {
        matched, err := regexp.Match(pattern, []byte(email))
        check(err)
        if matched {
            fmt.Printf("√ '%s' is a valid email\n", email)
        } else {
            fmt.Printf("X '%s' is not a valid email\n", email)
        }
    }
}