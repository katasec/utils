# Errx

Use :

```go
errx.PanicOnError(err)
```

Instead of:

```go
if err !=nil {
    panic(err)
}
```
