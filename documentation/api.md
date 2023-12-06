


# alchemy
  

## Informations

### Version

1.0

### Contact

  

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### BearerAuth (header: Authorization)



> **Type**: apikey

## All endpoints

###  credential

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/login | [post login](#post-login) | create Token |
  


###  spaceship

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/spaceship | [get spaceship](#get-spaceship) | get list spaceship |
| GET | /api/v1/spaceship/id/{id} | [get spaceship ID ID](#get-spaceship-id-id) | get spaceship |
| POST | /api/v1/spaceship | [post spaceship](#post-spaceship) | create spaceship |
| POST | /api/v1/spaceship/id/{id} | [post spaceship ID ID](#post-spaceship-id-id) | update spaceship |
  


## Paths

### <span id="get-spaceship"></span> get list spaceship (*GetSpaceship*)

```
GET /api/v1/spaceship
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BearerAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| class | `query` | string | `string` |  |  |  |  |
| name | `query` | string | `string` |  |  |  |  |
| status | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-spaceship-200) | OK | data |  | [schema](#get-spaceship-200-schema) |
| [400](#get-spaceship-400) | Bad Request | error message |  | [schema](#get-spaceship-400-schema) |

#### Responses


##### <span id="get-spaceship-200"></span> 200 - data
Status: OK

###### <span id="get-spaceship-200-schema"></span> Schema
   
  

[SpaceshipFindSpaceshipResponse](#spaceship-find-spaceship-response)

##### <span id="get-spaceship-400"></span> 400 - error message
Status: Bad Request

###### <span id="get-spaceship-400-schema"></span> Schema
   
  

[HTTPHTTPError](#http-http-error)

### <span id="get-spaceship-id-id"></span> get spaceship (*GetSpaceshipIDID*)

```
GET /api/v1/spaceship/id/{id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BearerAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | spaceship id |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-spaceship-id-id-200) | OK | data |  | [schema](#get-spaceship-id-id-200-schema) |
| [400](#get-spaceship-id-id-400) | Bad Request | error message |  | [schema](#get-spaceship-id-id-400-schema) |

#### Responses


##### <span id="get-spaceship-id-id-200"></span> 200 - data
Status: OK

###### <span id="get-spaceship-id-id-200-schema"></span> Schema
   
  

[SpaceshipSpaceshipResponse](#spaceship-spaceship-response)

##### <span id="get-spaceship-id-id-400"></span> 400 - error message
Status: Bad Request

###### <span id="get-spaceship-id-id-400-schema"></span> Schema
   
  

[HTTPHTTPError](#http-http-error)

### <span id="post-login"></span> create Token (*PostLogin*)

```
POST /api/v1/login
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [AuthLoginRequest](#auth-login-request) | `models.AuthLoginRequest` | | ✓ | | login request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-login-200) | OK | Data |  | [schema](#post-login-200-schema) |
| [400](#post-login-400) | Bad Request | Error message |  | [schema](#post-login-400-schema) |

#### Responses


##### <span id="post-login-200"></span> 200 - Data
Status: OK

###### <span id="post-login-200-schema"></span> Schema
   
  

[AuthLoginResponse](#auth-login-response)

##### <span id="post-login-400"></span> 400 - Error message
Status: Bad Request

###### <span id="post-login-400-schema"></span> Schema
   
  

[HTTPHTTPError](#http-http-error)

### <span id="post-spaceship"></span> create spaceship (*PostSpaceship*)

```
POST /api/v1/spaceship
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BearerAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [SpaceshipCreateSpaceshipRequest](#spaceship-create-spaceship-request) | `models.SpaceshipCreateSpaceshipRequest` | | ✓ | | login request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-spaceship-201) | Created | success true |  | [schema](#post-spaceship-201-schema) |
| [400](#post-spaceship-400) | Bad Request | success false |  | [schema](#post-spaceship-400-schema) |

#### Responses


##### <span id="post-spaceship-201"></span> 201 - success true
Status: Created

###### <span id="post-spaceship-201-schema"></span> Schema
   
  

[HTTPHTTPResponse](#http-http-response)

##### <span id="post-spaceship-400"></span> 400 - success false
Status: Bad Request

###### <span id="post-spaceship-400-schema"></span> Schema
   
  

[HTTPHTTPResponse](#http-http-response)

### <span id="post-spaceship-id-id"></span> update spaceship (*PostSpaceshipIDID*)

```
POST /api/v1/spaceship/id/{id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BearerAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | spaceship id |
| request | `body` | [SpaceshipCreateSpaceshipRequest](#spaceship-create-spaceship-request) | `models.SpaceshipCreateSpaceshipRequest` | | ✓ | | login request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-spaceship-id-id-200) | OK | success true |  | [schema](#post-spaceship-id-id-200-schema) |
| [400](#post-spaceship-id-id-400) | Bad Request | success false |  | [schema](#post-spaceship-id-id-400-schema) |

#### Responses


##### <span id="post-spaceship-id-id-200"></span> 200 - success true
Status: OK

###### <span id="post-spaceship-id-id-200-schema"></span> Schema
   
  

[HTTPHTTPResponse](#http-http-response)

##### <span id="post-spaceship-id-id-400"></span> 400 - success false
Status: Bad Request

###### <span id="post-spaceship-id-id-400-schema"></span> Schema
   
  

[HTTPHTTPResponse](#http-http-response)

## Models

### <span id="auth-login-request"></span> auth.LoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="auth-login-response"></span> auth.LoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| token | [AuthToken](#auth-token)| `AuthToken` |  | |  |  |



### <span id="auth-token"></span> auth.Token


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_token | string| `string` |  | |  |  |
| refresh_token | string| `string` |  | |  |  |



### <span id="http-http-error"></span> http.HTTPError


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | [interface{}](#interface)| `interface{}` |  | |  |  |



### <span id="http-http-response"></span> http.HTTPResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| success | boolean| `bool` |  | |  |  |



### <span id="spaceship-armament"></span> spaceship.Armament


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| qty | string| `string` |  | |  |  |
| title | string| `string` |  | |  |  |



### <span id="spaceship-create-spaceship-request"></span> spaceship.CreateSpaceshipRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| armament | [][SpaceshipArmament](#spaceship-armament)| `[]*SpaceshipArmament` | ✓ | |  |  |
| class | string| `string` | ✓ | |  |  |
| crew | integer| `int64` |  | |  |  |
| image | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| status | string| `string` | ✓ | |  |  |
| value | number| `float64` |  | |  |  |



### <span id="spaceship-data-spaceship"></span> spaceship.DataSpaceship


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |



### <span id="spaceship-find-spaceship-response"></span> spaceship.FindSpaceshipResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [][SpaceshipDataSpaceship](#spaceship-data-spaceship)| `[]*SpaceshipDataSpaceship` |  | |  |  |



### <span id="spaceship-spaceship-response"></span> spaceship.SpaceshipResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| armament | [][SpaceshipArmament](#spaceship-armament)| `[]*SpaceshipArmament` |  | |  |  |
| class | string| `string` |  | |  |  |
| crew | integer| `int64` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| image | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| value | number| `float64` |  | |  |  |


