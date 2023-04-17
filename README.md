# Notes

## Golang Tools
### Version
```
>> go version
```

### gofmt
Applied format to the whole project
```
>> gofmt -s -w .
```

### interface
The interface type that specifies zero methods is known as the empty interface:
```
interface{}
```
An empty interface may hold values of any type. (Every type implements at least zero methods.)

`any` is a new predeclared identifier and a type alias of `interface{}`.