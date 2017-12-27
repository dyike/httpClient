## httpClient 

> http requset lib.

### GET 

```go
url := "http://example.com"

result, err := httpClient.DoRequest(httpClient.Request{
    Method: "GET",
    URL:    url,
})
fmt.Println(result, err)
```

### PUT 

```go
url := "http://example.com"

data := map[string]string{
    "go":   "golang",
}
body, _ := json.Marshal(data)

header := http.Header{}
header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4")

result, err := httpClient.DoRequest(httpClient.Request{
    Method: "PUT",
    URL:    url,
    Body:   body,
    Header: header,
})
fmt.Println(result, err)
```


### POST 

```go

url := "http://example.com"

data := map[string]string{
    "go":   "golang",
}
body, _ := json.Marshal(data)

header := http.Header{}
header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4")

result, err := httpClient.DoRequest(httpClient.Request{
    Method: "POST",
    URL:    url,
    Body:   body,
    Header: header,
})
fmt.Println(result, err)

```


### DELET 

```go

url := "http://example.com"

result, err := httpClient(httpClient.Request{
    Method: "DELETE",
    URL:    url,
})
fmt.Println(result, err)

```

