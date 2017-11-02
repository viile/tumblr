package main

import (
    "fmt";
    "github.com/tumblr/tumblrclient.go"
)

func main() {
    fmt.Printf("heelo world")

    client := tumblrclient.NewClient(
        "aMKe1gA6lqjXEy7gETNlxboKaq9V6el1Nj2kKPDbZdIKfkVvuE",
        "W3kG5fs0sgMTa0ZWepabN0BMJLAL6eq05Y4r22KUAFSdQlgWTc",
    )

    _,err := client.Get("viile-things.tumblr.com")
    if err != nil {
        
    }
    fmt.Printf("%s","123")
}
