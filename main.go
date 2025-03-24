package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    // Input URL website yang akan diatasi
    var targetURL string
    fmt.Print("Masukkan URL website target: ")
    fmt.Scanln(&targetURL)

    // Input jumlah request yang dikirim
    var numRequests int
    fmt.Print("Masukkan jumlah request: ")
    fmt.Scanln(&numRequests)

    // Input interval antar request dalam milidetik
    var interval int
    fmt.Print("Masukkan interval antar request (dalam milidetik): ")
    fmt.Scanln(&interval)

    // User-Agents yang akan digunakan
    userAgents := []string{
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 Edg/89.0.774.68",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
    }

    // Melakukan pengiriman request secara berulang
    for i := 0; i < numRequests; i++ {
        // Membuat request baru
        req, err := http.NewRequest("GET", targetURL, nil)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }

        // Mengubah user-agent secara acak
        req.Header.Set("User-Agent", userAgents[i%len(userAgents)])

        // Melakukan request ke website target
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }

        // Menutup response body
        resp.Body.Close()

        // Menampilkan status request
        fmt.Printf("Request %d: %s\n", i+1, resp.Status)

        // Menunggu selama interval milidetik
        time.Sleep(time.Duration(interval) * time.Millisecond)
    }
}
