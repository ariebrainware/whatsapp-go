## Planning Project
Creating a JSON API that will allow users to create, read, update and delete events.

## JWT (json web token)
* Header: the type of token and the signing algorithm used.
* payload: the second part of the token which contains the claims. These claims include application specific data (e.g, user id, username), token expiration time (exp), issuer(iss), subject (sub) and so on.
* signature: the encoded header, encoded payload and a secret we provide are used to create the signature.
```javascript
Token = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM
```

## token types:
* access token: an access token in used for requests that require authentication. it is nomarlly added in the header of the request. it is recommended that an access token have a short lifespan (15 minutes).
* refresh token: a refresh token has a longer lifespan, usually 7 days. this token is used to generate new access and refresh tokens. In the event that the access token expires , new sets of access and refresh tokens are created when the refresh token route is hit (from our application).

## User table
DROP TABLE IF EXISTS `user`; 
    CREATE TABLE `user`( 
    `id` int(11) NOT NULL AUTO_INCREMENT, 
    `username` varchar(30) NOT NULL, 
    `password` varchar(30) NOT NULL, 
    PRIMARY KEY (`id`));