# GophersCom

**gopherscom** is a GOPHERS(Go Developers) coding solution, trend in **Go** Programming Language related area and also as a knowledge sharing community especially for Golang developers. It is implemented using [Go](https://golang.org/) and **Go's** framework [Gin](https://github.com/gin-gonic) and [Postgresql](https://www.postgresql.org/) as persistant database. **gopherscom** provides a [RESTful Service](https://restfulapi.net/)
 and features like [Geojson](http://geojson.io/), [Email
auth](https://mail.google.com/), [JWT with refreshToken](http://www.passportjs.org/) and [Redis Performance Caching](https://redis.io/) are included currently and more will be
added soon. Here is a deployed link **[live on heroku](https://gopherscom.herokuapp.com/)** .


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

#### User Model

```go

type User struct {
	ID        string    `json:"_id" json:"id"`
	NAME      string    `json:"name"`
	EMAIL     string    `json:"email"`
	PASSWORD  string    `json:"password"`
	CREATEDAT time.Time `json:"created_at" json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at" json:"updated_at"`
}


```
##



### Profile
**[Profile](./docs/profile.md)** route is mostly concerned with user detail information and it provides data from both `Postgresql database` and `Redis cache`.

#### Profile Model

```go

type Profile struct {
	ID         string    `json:"id"`
	USERID     string    `json:"userid"`
	CAREER     string    `json:"career"`
	FRAMEWORKS []string  `json:"frameworks"`
	LANGUAGES  []string  `json:"languages"`
	PLATFORMS  []string  `json:"platforms"`
	DATABASES  []string  `json:"databases"`
	OTHERS     []string  `json:"others"`
	SEX        string    `json:"sex"`
	BIRTHDATE  string    `json:"birthdate"`
	ADDRESS    string    `json:"address"`
	ZIPCODE    string    `json:"zipcode"`
	CITY       string    `json:"city"`
	STATE      string    `json:"state"`
	COUNTRY    string    `json:"country"`
	LAT        float64   `json:"lat"`
	LON        float64   `json:"lon"`
	CREATEDAT  time.Time `json:"created_at"`
	UPDATEDAT  time.Time `json:"updated_at"`
}


```
##




### Blog
**[Blog](./docs/blog.md)** route is mostly concerned with `Blog` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Blog Service** is one of the main features of **GophersCom**.

#### Blog Model

```go

type Blog struct {
	ID        string    `json:"id"`
	TITLE     string    `json:"title"`
	BODY      string    `json:"body"`
	PUBLIC    bool      `json:"public"`
	APPTYPE   string    `json:"apptype"`
	LANGUAGES []string  `json:"languages"`
	TAGS      []string  `json:"tags"`
	LIBRARIES []string  `json:"libraries"`
	AUTHOR    string    `json:"author"`
	LIKES     []Like    `json:"likes"`
	COMMENTS  []Comment `json:"comments"`
	CREATEDAT time.Time `json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at"`
}


```
##


### Apptype
**[Apptype](./docs/apptype.md)** route is mostly concerned with `Apptype tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **apptype tag items**.

#### Apptype Model

```go

type Apptype struct {
	ID          string    `json:"id"`
	NAME        string    `json:"name"`
	DESCRIPTION string    `json:"description"`
	AUTHOR      string    `json:"author"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}


```
##


### Library
**[Library](./docs/library.md)** route is mostly concerned with `Library tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **library tag items**.

#### Library Model

```go

Library := Apptype


```
##


### Other
**[Other](./docs/other.md)** route is mostly concerned with `Other tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **other tag items**.

#### Other Model

```go

Other := Apptype


```
##


### Platform
**[Platform](./docs/platform.md)** route is mostly concerned with `Platform tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **platform tag items**.

#### Platform Model

```go

Platform := Apptype


```
##


### Tag
**[Tag](./docs/tag.md)** route is mostly concerned with `tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **tag items**.

#### Tag Model

```go

Tag := Apptype


```
##


### Language
**[Language](./docs/language.md)** route is mostly concerned with `language tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **language tag items**.

#### Library Model

```go

Language := Apptype


```
##

### Framework
**[Framework](./docs/framework.md)** route is mostly concerned with `framework tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **framework tag items**.

#### Framework Model

```go

Framework := Apptype


```
##

### Database
**[Database](./docs/database.md)** route is mostly concerned with `database tag` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. Its main use case is in `Frontend` dropdown or list of **database tag items**.

#### Database Model

```go

Other := Apptype


```
##



### Bootcamp
**[Bootcamp](./docs/bootcamp.md)** route is mostly concerned with `Bootcamp` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Bootcamp Service** is one of the main features of **GophersCom**.

#### Bootcamp Model

```go

type Bootcamp struct {
	ID          string     `json:"id"`
	TOPIC       string     `json:"topic"`
	INSTRUCTORS []string   `json:"instructors"`
	ADDRESS     string     `json:"address"`
	LAT         float64    `json:"lat"`
	LON         float64    `json:"lon"`
	STUDENTS    string     `json:"students"`
	ENROLLMENTS []Enroller `json:"enrollments"`
	DESCRIPTION string     `json:"description"`
	AVAILABLE   bool       `json:"available"`
	STARTEDAT   string     `json:"startedat"`
	FINISHEDAT  string     `json:"finishedat"`
	AUTHOR      string     `json:"author"`
	LIKES       []Like     `json:"likes"`
	COMMENTS    []Comment  `json:"comments"`
	CREATEDAT   time.Time  `json:"created_at"`
	UPDATEDAT   time.Time  `json:"updated_at"`
}


```

##




### Company
**[Company](./docs/company.md)** route is mostly concerned with `Company` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Company Service** is one of the main features of **GophersCom**.

#### Company Model

```go

type Company struct {
	ID          string    `json:"id"`
	NAME        string    `json:"name"`
	PRODUCTS    []string  `json:"products"`
	EMPLOYEE    string    `json:"employee"`
	BRANCHES    []*Branch `pg:",one2many:company_branches"`
	FRAMEWORKS  []string  `json:"frameworks"`
	LANGUAGES   []string  `json:"languages"`
	PLATFORMS   []string  `json:"platforms"`
	DATABASES   []string  `json:"databases"`
	OTHERS      []string  `json:"others"`
	ADDRESS     string    `json:"address"`
	ZIPCODE     string    `json:"zipcode"`
	CITY        string    `json:"city"`
	STATE       string    `json:"state"`
	COUNTRY     string    `json:"country"`
	LAT         float64   `json:"lat"`
	LON         float64   `json:"lon"`
	FOUNDEDYEAR string    `json:"foundedyear"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}


```
##



### Branch
**[Branch](./docs/branch.md)** route is mostly concerned with `Company Branch` related information and it provides `CRUD` operations for both `Postgresql database` and `Redis cache`. **Company Service** is one of the main features of **GophersCom** and **Company Branch** is to provide the `connection & relationship between Company and Company Branches and how are related and where are they located with how many people in office`.

#### Company Branch Model

```go

type Branch struct {
	ID          string    `json:"id"`
	CID         string    `json:"cid"`
	NAME        string    `json:"name"`
	ADDRESS     string    `json:"address"`
	ZIPCODE     string    `json:"zipcode"`
	CITY        string    `json:"city"`
	STATE       string    `json:"state"`
	COUNTRY     string    `json:"country"`
	LAT         float64   `json:"lat"`
	LON         float64   `json:"lon"`
	FOUNDEDYEAR string    `json:"foundedyear"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}



```
##

##
##