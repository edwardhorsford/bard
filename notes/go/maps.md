# Maps

```go
var my_map = make(map[string]int)
my_map["cat"] = "Mr.Meow"
```

A map is an unordered collection of key-value pairs. Also known as an
associative array, a hash table or a dictionary, maps are used to look up a
value by its associated key.

## Retrieving a value

```go
my_map["cat"] // "Mr.Meow"
```

## Deleting a value

```go
delete(my_map, "cat")
```
