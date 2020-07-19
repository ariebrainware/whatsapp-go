## Planning Project
Creating a JSON API that will allow users to create, read, update and delete events.

## JWT (json web token)
* Header: the type of token and the signing algorithm used.
* payload: the second part of the token which contains the claims. These claims include application specific data (e.g, user id, username), token expiration time (exp), issuer(iss), subject (sub) and so on.
* signature: the encoded header, encoded payload and a secret we provide are used to create the signature.
```javascript
Token = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM
```