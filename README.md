> this is a simple signup/login system, **using GO, GraphQL, Postgres and Docker**

**Functionalities:**  
- create account
- login to an account
- logout
- get the logged in user data

#### Development
first of all, clone the repo.  

cd into the `/server` and then run

```bash
go mod download
```

now cd into `/server/docker` and run

```bash
docker-compose up
```

if everything goes fine you can go to `http://localhost:8080` and you shall see the GraphQL interface

here is an example of queries/mutations available

```graphql
# create account / signup
mutation {
  createUser(data: {
    name: "yourname",
    email: "yourname@gmail.com",
    password: "12345"
  })
}
```

```graphql
# login
mutation {
  loginUser(data: {
    email: "yourname@gmail.com",
    password: "12345"
  })
}
```

```graphql
# logout
mutation {
  logout
}
```

```graphql
# Get logged in user data 
query {
  getAuthUser{
    id
    name
    email
    createdAt
  }
}
```

**The Front end part will come soon, but you can implement it by your own if you want**
