# GlobalWebIndex Engineering Challenge

This is a service for signing up and logging in users. Logged in users may create, read, update or delete their relevant asset data. This service provides basic auth functionality and persists in a postgres sql database built for this reason in ```https://api.elephantsql.com/```.

## How to run

Assuming that docker is already installed, download (clone) the source code.
```https://github.com/Thodoras/platform2.0-go-challenge/```

After getting in the folder that the source code is downloaded build the docker image. Below is an example:
```sudo docker build -t go-project-app .```

Now run the docker image. Following from the previous example run:
```sudo docker run -it -p 8002:8001  go-project-app```

The above example listens to:
```127.0.0.1:8002/```

## API

If it's the first time to use this service. First use

```POST /users/signup```

with a message body:

```
{
  "name": "foo",
  "password": "Fo0b4r1!@"
}
```

It should be noted that a name of 3 to 20 characters is needed and a password from 6 to 16 with at least one lower case letter, one upper case letter, one digit and one special character (!@#$%^&*).

Then login:

```POST /users/login```

with the same message body used to signup. The response will contain user ```id``` and a ```token```.
The token should be used in the headers (with key: Token, and value: the value of the 'token' field in the response), and the id should be used in the ```user_id``` section of the url.

**The rest of the api is:**

```GET /assets/{user_id}``` Returns all the user relevant data.

The response looks like this:
```
{
    "user_id": 1,
    "audiences": [
        {
            "id": 0,
            "user_id": 1,
            "gender": "m",
            "birth_country": "uk",
            "age_groups": "20-30",
            "hours_spent": 3,
            "num_of_purchases_per_month": 1
        },
        {
            "id": 2,
            "user_id": 1,
            "gender": "f",
            "birth_country": "gr",
            "age_groups": "20-30",
            "hours_spent": 2,
            "num_of_purchases_per_month": 0
        },
    ],
    "charts": [
        {
            "id": 1,
            "user_id": 1,
            "title": "foo",
            "axis_x_title": "foox",
            "axis_y_title": "fooy",
            "data": "fooData"
        },
    ],
    "insights": [
        {
            "id": 1,
            "user_id": 1,
            "text": "some text"
        }
    ]
}
```
Alternativelly in order to secure that the request will take reasonably long the client may call the following endpoint providing ```limit``` and ```offset``` as parameters in the url.

```GET /assets/paginated/{user_id}```</br>

The following are simillar with the above two endpoints but return only the favoured by the user assets.</br>
```GET /assets/{user_id}/favourites``` </br>
```GET /assets/paginated/{user_id}/favourites```


The response will be similar with the above.


```POST /assets/audiences/{user_id}``` To add an audience with request:
```
{
	"gender": "f",
	"birth_country": "gr",
	"age_groups": "40-50",
	"hours_spent": 1,
	"num_of_purchases_per_month": 1
}
```
```POST /assets/charts/{user_id}``` To add a chart with request:
```
{
	"title": "ba",
	"axis_x_title": "barx",
	"axis_y_title": "bary",
	"data": "barData"
}
```
```POST /assets/insights/{user_id}``` To add am insight with request:
```
{
	"text": "some other text"
}
```

For the above three the response is an asset ```id``` number.

```PUT /assets/audiences/{user_id}``` To update am audience (a similar request with POST) </br>
```PUT /assets/charts/{user_id}``` To update am chart (a similar request with POST) </br>
```PUT /assets/insights/{user_id}``` To update am insight (a similar request with POST) </br>

The response for the above three requests is a number which indicate the rows that were altered (should be 1).

```DELETE /assets/audiences/{user_id}/delete/{id}``` To delete an audience with a given ```id``` </br>
```DELETE /assets/charts/{user_id}/delete/{id}``` To delete a chart with a given ```id``` </br>
```DELETE /assets/insights/{user_id}/delete/{id}``` To delete an insight with a given ```id``` </br>

Again the response for the above three requests is a number which indicate the rows that were deleted (should be 1).

##Dependencies

```github.com/dgrijalva/jwt-go```  for generating tokens </br>
```go get github.com/gorilla/mux```  for creating routing points with restful verbs </br>
```go get github.com/lib/pq```  for connecting with postgres sql </br>
```go get github.com/subosito/gotenv```  for creating a config file with globals </br>
```go get golang.org/x/crypto/bcrypt``` for hashing passwords to store in database and verification between hashed and hashed passwords </br>
