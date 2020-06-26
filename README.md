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
**[Blog](./docs/blog.md)** route is mostly concerned with `Blog` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`.
##


### Apptype
**[Apptype](./docs/apptype.md)** route is mostly concerned with `Apptype tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of *apptype tag items*.
##


### Library
**[Library](./docs/library.md)** route is mostly concerned with `Library tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of *library tag items*.
##


### Other
**[Other](./docs/other.md)** route is mostly concerned with `Other tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of *other tag items*.
##


### Platform
**[Platform](./docs/platform.md)** route is mostly concerned with `Platform tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of *platform tag items*.
##


### Tag
**[Tag](./docs/tag.md)** route is mostly concerned with `tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of *tag items*.
##


### Language
- [Language](./docs/language.md)
##

### Framework
- [Framework](./docs/framework.md)
##

### Database
- [Database](./docs/database.md)
##

### Bootcamp
- [Bootcamp](./docs/bootcamp.md)
##

### Company
- [Company](./docs/company.md)
##

### Branch
- [Branch](./docs/branch.md)
##

##
##