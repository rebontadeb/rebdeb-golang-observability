
### Run Redis Based on Docker

```
docker run --name some-redis -p 6379:6379 -v /Users/path/Documents/redispvc:/data -d redis redis-server --save 60 1 --loglevel warning
```
