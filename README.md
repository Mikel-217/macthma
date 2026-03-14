# About
This is a school project.
It's a matchmaking server where user can connect to and get transfared to a lobby with up to 20 Players.
The users get ranked from:
- Total hours
- Kills in the last game
- Last win
- Last Position 

## Deployment / Testing

``` bash

docker-compose up

```

For production, follow the instructions in the [docker-compose](docker-compose.yml) file

## Enviromentvariables needed:
- Database connection
- JWT-Secret

In the [docker-compose](docker-compose.yml) are examples for both.

## TODO:
- Add algorithm for creating lobby
- Add UML -> because of school project :)
- Make db_excecuter performant -> some querys aren´t executed
  - Batching Updates
  - Bulk Updates

## Endpoints:
### Post endpoints:
- /login -> for sending the user an accesstoken
- /register -> for a user to register
- /match-data -> for creating new data based on the match
### Get endpoints:
- /join-match -> for creating a websocket
