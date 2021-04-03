### README

## how to run this application

You first need to build the application with the command below.
``` bash
make build
```

Subsequently, you need to feed the mock data to the database.
``` bash
make migrate
```

Then, you can start up the server with this command.
``` bash
make up
```

``` gql
query {
  scooter (latitude: 1.39, longitude: 103.9, limit:20, distance: 5000) {
      title
      latitude
      longitude
    }
}
```
