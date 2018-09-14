# Nulls
Nulls is a simple library that helps you to marshal null database values to there default zero values in golang.



# nulls.NullString
For string field that you know that may contain a null value use (nulls.NullString) instead of (sql.NullString)or string.
This will marshal  to  "" for all null string.You can easily change this behaviour to whatever you wish your null strings to marshal  to.

## Example
```go
type Pojo struct {
	ID   uint `json:"id"`
	Content  NullString `json:"content"`
	
}
```

### nulls.NullInt64
This will marshal  to 0 for all null Int64 fields.This is a replacement of the sql..NullInt64.If you want to change this behaviour you can easily change it.

### Example usage
```go
type Pojo struct {
	ID   uint `json:"id"`
  Complete    NullInt64   `json:"complete"`
	
	
}
```

### nulls.NullFloat64
This will return 0.0 for all null Float64  fields.

### Example usage
```go
type Pojo struct {
	ID   uint `json:"id"`
  Price    nulls.NullFloat64  `json:"complete"`
		
}
```

### nulls.NullBool
This will marshal to  false for all null bool  fields.

### Example usage
```go
type Pojo struct {
	ID   uint `json:"id"`
  Status    nulls.NullBool  `json:"status"`
		
}
```




