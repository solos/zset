[![Go Report Card](https://goreportcard.com/badge/github.com/solos/zset)](https://goreportcard.com/report/github.com/solos/zset)
# zset
Implementing sorted set in Redis with golang.

## Installation
```bash
go get -u github.com/solos/zset
```

## Usage
Removed RWLock in the SortedSet. 
Just implement it yourself if needed.
```go

s := zset.New()
// add data
s.Set(0.99, "600600", nil)
s.Set(-0.52, "600519", nil)
s.Set(-2.13, "601398", nil)

stocks := []string{"600600", "600519", "601398"}

for index := range stocks {
    stock_code := stocks[index]
    rank, score, _ := s.GetRank(stock_code, false)
    fmt.Println(rank)
    fmt.Println(score)
}


```
