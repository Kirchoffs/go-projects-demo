# Crawler

## Create the Project
```
>> mkdir crawler && cd crawler
>> go mod init crawler
```

## Add Dependencies
```
>> go get golang.org/x/net/html/charset
>> go get golang.org/x/text/encoding
```

## Run the Project
```
>> go build main.go
>> ./main
```

```
>> go run main.go > thepaper.html
```

## Code Details
### strings & bytes
```
import strings

body := io.ReadAll(resp.Body)

numLinks := strings.Count(string(body), "<a")
exist := strings.Contains(string(body), "<a")
```

```
import strings

body := io.ReadAll(resp.Body)

numLinks := bytes.Count(body, []byte("<a"))
exist := bytes.Contains(body, "<a")
```

## Golang Details
### Format Print
- __%v__ is a placeholder that represents the value of an operand. It is used to interpolate the value of a variable into a string. The behavior of %v depends on the type of the operand being printed.