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

For production, uncomment the "command:" line 

## Enviromentvariables needed:
- Database connection
- JWT-Secret

## TODO:
- Make middleware functional
- Add tracking interface for matches
- Add algorithm for creating lobby
- Add UML -> because of school project :)


## Endpoints:
