# gopherscom

**gopherscom** is a Gophers coding solution, trend in go and knowledge sharing community especially for Golang developers. It is implemented using [Go](https://golang.org/) and **Go's** framework [Gin](https://github.com/gin-gonic) and [Postgresql](https://www.postgresql.org/) as persistant database. **gopherscom** provides a [RESTful Service](https://restfulapi.net/)
 and features like [Geojson](http://geojson.io/), [Email
auth](https://mail.google.com/), [JWT with refreshToken](http://www.passportjs.org/) and [Redis Performance Caching](https://redis.io/) are included currently and more will be
added soon.Here is a deployed link [HEROKU](https://gopherscom.herokuapp.com/)


##
##


## API LIST

##

```go
const PUBLIC_URL = "https://gopherscom.herokuapp.com/public"
const PROTECTED_URL = "https://gopherscom.herokuapp.com/protected"
const BASE_URL = "https://gopherscom.herokuapp.com/"
```
##


### User
**[User](./docs/user.md)** is a about authetication and credentials. It primary have create account feature, `accessToken` and `refreshToken` feature.
##


### Profile
**[Profile](./docs/profile.md)** route is mostly concerned with user detail information and it provides data from both `Postgresql database` and `Redis cache`.
##


### Blog
**[Blog](./docs/blog.md)** route is mostly concerned with `Blog` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Blog Service** is one of the main features of **GophersCom**.
##


### Apptype
**[Apptype](./docs/apptype.md)** route is mostly concerned with `Apptype tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **apptype tag items**.
##


### Library
**[Library](./docs/library.md)** route is mostly concerned with `Library tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **library tag items**.
##


### Other
**[Other](./docs/other.md)** route is mostly concerned with `Other tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **other tag items**.
##


### Platform
**[Platform](./docs/platform.md)** route is mostly concerned with `Platform tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **platform tag items**.
##


### Tag
**[Tag](./docs/tag.md)** route is mostly concerned with `tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **tag items**.
##


### Language
**[Language](./docs/language.md)** route is mostly concerned with `language tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **language tag items**.
##

### Framework
**[Framework](./docs/framework.md)** route is mostly concerned with `framework tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **framework tag items**.
##

### Database
**[Database](./docs/database.md)** route is mostly concerned with `database tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **database tag items**.
##

### Bootcamp
**[Bootcamp](./docs/bootcamp.md)** route is mostly concerned with `Bootcamp` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Bootcamp Service** is one of the main features of **GophersCom**.
##

### Company
**[Company](./docs/company.md)** route is mostly concerned with `Company` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Company Service** is one of the main features of **GophersCom**.
##

### Branch
**[Branch](./docs/branch.md)** route is mostly concerned with `Company Branch` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Company Service** is one of the main features of **GophersCom** and **Company Branch** is to provide the `connection & relationship between Company and Company Branches and how are related and where are they located with how many people in office`.
##

##
##