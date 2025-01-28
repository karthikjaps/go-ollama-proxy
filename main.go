package main

import (
    "io"
    "log"
    "net/http"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
    client := &http.Client{}

    // Forward the request to Ollama (adjust port if needed)
    targetURL := "http://localhost:11434" + r.URL.Path

    // Create a new request with the same method and body
    req, err := http.NewRequest(r.Method, targetURL, r.Body)
    if err != nil {
        http.Error(w, "Error creating request", http.StatusInternalServerError)
        return
    }

    // Copy headers from the incoming request
    req.Header = r.Header

    // Send the request to Ollama
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Error forwarding request", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Write Ollama's response back to the client
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

func main() {
    http.HandleFunc("/", proxyHandler)
    log.Println("Proxy server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
