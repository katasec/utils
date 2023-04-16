# Create a list of unique items

`Set` is a list of unique items.

## Create a list of unique `ints`

```go
// Create the set
myset := set.NewSet[int]()

// Add some items
myset.Add(5)
myset.Add(6)
myset.Add(7)

// Remove some items
myset.Del(7)

// Output list
for i := range myset.Items {
    fmt.Printf("item=%d\n", i)
}
```
