# blog

## JSON Objects returned by API:

Make sure the right content type like `Content-Type: application/json; charset=utf-8` is correctly returned.

### Authentication
``` JSON
{
    "status": 1,
    "message": "",
    "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhhcnZleWpodWFuZ0BnbWFpbC5jb20iLCJleHAiOjE1ODYwNzQ2NzAsImlhdCI6MTU4NjA3MTA3MCwiaWQiOiIxIn0.4kLiw0BBjc9Fl4rendjK0Av2Zi30zHVo7q1B9XThbhQ",
        "tokenType": "bearer",
        "expiresIn": 3600
    }
}
```


### User

```JSON
{
    "status": 1,
    "message": "",
    "data": {
        "username": "Harvey",
        "email": "harveyjhuang@example.com",
        "bio": "I have a dream!",
        "image": "https://example.com/my.jpg"
    }
}
```

### Single Article

```JSON

```

### Multiple Articles

```JSON

```


### List of Tags

```JSON

```

### Errors and Status Codes

If a request fails any validations, expect a 400 and errors in the following format:

```JSON
{
    "status": 0,
    "message": "record not found",
    "data": null
}
```

#### Other status codes:

401 for Unauthorized requests, when a request requires authentication but it isn't provided

403 for Forbidden requests, when a request may be valid but the user doesn't have permissions to perform the action

404 for Not found requests, when a resource can't be found to fulfill the request


## Endpoints:

### Authentication:

`POST /api/users/login`

Example request body:
```JSON
{
    "email": "harveyjhuang@example.com",
    "password": "test1234"
}
```

No authentication required, returns a [Authentication](#authentication)

Required fields: `email`, `password`


### Registration:

`POST /api/users`

Example request body:
```JSON
{
    "username": "Harvey",
    "email": "harveyjhuang@example.com",
    "password": "test1234"
}
```

No authentication required, returns a [User](#user)

Required fields: `email`, `username`, `password`



### Get Current User

`GET /api/user`

Authentication required, returns a [User](#user) that's the current user



### Update User

`PUT /api/user`

Example request body:
```JSON
{
  "user":{
    "email": "jake@jake.jake",
    "bio": "I like to skateboard",
    "image": "https://i.stack.imgur.com/xHWG8.jpg"
  }
}
```

Authentication required, returns the [User](#users-for-authentication)


Accepted fields: `email`, `username`, `password`, `image`, `bio`


### List Articles

`GET /api/articles`

Returns most recent articles globally by default, provide `tag`, `author` or `favorited` query parameter to filter results

Query Parameters:

Filter by tag:

`?tag=AngularJS`

Filter by author:

`?author=jake`

Favorited by user:

`?favorited=jake`

Limit number of articles (default is 20):

`?limit=20`

Offset/skip number of articles (default is 0):

`?offset=0`

Authentication optional, will return [multiple articles](#multiple-articles), ordered by most recent first


### Get Article

`GET /api/articles/:slug`

No authentication required, will return [single article](#single-article)

### Create Article

`POST /api/articles`

Example request body:

```JSON
{
  "article": {
    "title": "How to train your dragon",
    "description": "Ever wonder how?",
    "body": "You have to believe",
    "tagList": ["reactjs", "angularjs", "dragons"]
  }
}
```

Authentication required, will return an [Article](#single-article)

Required fields: `title`, `description`, `body`

Optional fields: `tagList` as an array of Strings


### Update Article

`PUT /api/articles/:slug`

Example request body:

```JSON
{
  "article": {
    "title": "Did you train your dragon?"
  }
}
```

Authentication required, returns the updated [Article](#single-article)

Optional fields: `title`, `description`, `body`

The `slug` also gets updated when the `title` is changed


### Delete Article

`DELETE /api/articles/:slug`

Authentication required


### Get Tags

`GET /api/tags`

No authentication required, returns a [List of Tags](#list-of-tags)

## TODO

### API
[x] Register
[x] Login
[ ] Profile
[ ] Article
    [ ] Create
    [ ] Update
    [ ] List
    [ ] Get
    [ ] Delete

### Error Handle
[ ] error definition 

### GCP Setting
[ ] TSL
[ ] Firewall
[ ] Cloud Run
