## basic api by graphql with golang 

通过graphiql 请求

![graphiql](../../assets/01-basic-api-graphiql.png)

通过postman 请求

![postman-graphql](../../assets/01-basic-api-postman-graphql.png)

curl 无参数请求

```sh
➜ curl -XGET http://localhost:3000/graphql -d'{"query":"{jobs{id,company}}"}'
{"data":{"jobs":[{"company":"Apple","id":1},{"company":"Google","id":2},{"company":"Microsoft","id":3},{"company":"Facebook","id":4},{"company":"Facebook","id":5}]}}     
```

curl 带参数请求

```sh
➜ curl -XGET http://localhost:3000/graphql -d'{"query":"{job(id:3){id,company}}'                                    
{"data":{"job":{"company":"Microsoft","id":3}}}
```