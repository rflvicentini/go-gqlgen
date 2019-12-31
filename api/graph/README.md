# go-gqlgen
Golang GraphQL API using gqlgen

## GraphQL Queries

### User
```
query GetUsers {
  users {
    id,
    username,
    email,
    meetups {
      name
    }
  }
}
```

### Meetup

```
query GetMeetups {
  meetups {
    id,
    name,
    description,
    user{
      username,
      meetups {
        name
      }
    }
  }
}
```
