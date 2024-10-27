# Vote_Simple

## 什么是 Vote_simple？

Vote_Simple 是基于 GO 的 GraphQL 服务端。

该程序默认 2 秒钟生成一个随机的 ticket，每张 ticket 有效期为从服务端生成直到下一次生成新的 ticket 为止。每张 ticket 有最大使用次数，默认为 200 次。

服务端提供下列操作接口：

- **获取当前票证**

```
query {
  getTicket
}
```

- **投票**

```
mutation {
  vote(usernames: ["user1", "user2"], ticket: "the_ticket_id") 
}
```

- **查询特定用户票数**

```
query {
  queryVotes(username: "user1")
}
```

## Vote_Simple 有什么亮点？

* Vote_Simple 采用 badger 数据库，因此 Vote_Simple 简单轻便，能有效应对高并发场景，并且投票值也能持久化。
* Vote_Simple 采用 gqlgen 框架，提供了一个图形化的 GraphQL 客户端，仅需打开浏览器输入 `localhost:8080`即可。
* Vote_Simple 结构简单，适合简单进行 GraphQL 学习与二次开发。

## Vote_Simple 如何开始？

打开浏览器访问 `localhost:8080` 即可访问图形化的 GraphQL 客户端。

 

