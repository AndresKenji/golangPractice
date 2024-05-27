package main

import (
    "context"
    "fmt"
    "net/url"
    "os"

    "github.com/vmware/govmomi"
)

func main() {
    // Creating a connection context
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Parsing URL
	vcurl := ""
    url, err := url.Parse(vcurl)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
	fmt.Println(url)

    // Connecting to vCenter
    client, err := govmomi.NewClient(ctx, url, true)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }

    // vCenter version
    info := client.ServiceContent.About
    fmt.Printf("Connected to vCenter version %s\n", info.Version)
}