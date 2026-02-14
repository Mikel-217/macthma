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

For testing edit the Dockerfile and uncomment the "RUN ["go run main.go --testing"]"

## Enviromentvariables needed:
- Database connection
- JWT-Secret

## TODO:
- Add db-setup service
- Add test interface where random users are created
- Add user-login
- Make middleware functional
- Add tracking interface for matches
- Add algorithm for creating lobby
- Add UML -> because of school project :)
