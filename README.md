# gopherscom

**gopherscom** is a Gophers coding solution, trend in go and knowledge sharing community especially for Golang developers. It is implemented using [Go](https://golang.org/) and **Go's** framework [Gin](https://github.com/gin-gonic) and [Postgresql](https://www.postgresql.org/) as persistant database. **gopherscom** provides a [RESTful Service](https://restfulapi.net/)
 and features like [Geojson](http://geojson.io/), [Email
auth](https://mail.google.com/), [JWT with refreshToken](http://www.passportjs.org/) and [Redis Performance Caching](https://redis.io/) are included currently and more will be
added soon.Here is a deployed link [HEROKU](https://gopherscom.herokuapp.com/)


##

##


## API LIST

### User
[User](./docs/user.md)
Signup, Signin, Refreshtoken, Getusers</code> : routes are available now at the moment and more services/features will
be added soon.
<table class="table table-hover">
    <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
            <th scope="col">Types</th>
            <th scope="col">Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">1</th>
            <td><code>POST</code></td>
            <td>
                <code>
                  https://gopherscom.herokuapp.com/signup
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/signUp'</code> route will allow you to save new user data to
                Database.<code>{ unique email is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/signin
                </code>
            </td>
            <td><code>{String, String}</code></td>
            <td> <code>'/signIn'</code> route will allow you to login via database and generate
                two tokens.<code>accessToken</code> and <code>refreshToken</code>. </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>GET</code></td>
            <td>
                <code>
                 https://gopherscom.herokuapp.com/refreshToken
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/refreshToken'</code> route will allow you to refresh tokens without
                <code>sigining in</code> and
                <code>refreshToken</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
    </tbody>
</table>

- It's recommended to use **genuine email** in **signup** cuz there will be a feature to add like:: `every user account
need to be activated` using code via **email**

### Protected - USER
<code> GetUserList, GetUser, ResetUserCache </code> : routes are available now at the moment and more
services/features will be added soon.


<table class="table table-hover">
    <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
            <th scope="col">Types</th>
            <th scope="col">Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">1</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/user/list
                </code>
            </td>
            <td><code>{String, Objects, Slice}</code></td>
            <td> <code>'/'</code> route will allow you to fetach <code>users</code> data from
                <code>Postgresql/Redis</code>
                Database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/user?id={}
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>User</code> based
                on <code>id</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 https://gopherscom.herokuapp.com/protected/user/resetcache
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>User</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **User** if the request doesn't stuck in `errors` or `overcached`.



### Protected - PROFILE
<code> GetProfileList, GetProfileByUser, GetByID, CreateProfile, UpdateProfile, ResetProfileCache </code> : routes are available now at the moment and more
services/features will be added soon.


<table class="table table-hover">
    <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
            <th scope="col">Types</th>
            <th scope="col">Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">1</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/profile/list
                </code>
            </td>
            <td><code>{String, Objects, Slice}</code></td>
            <td> <code>'/profile/list'</code> route will allow you to fetch <code>slices of profile object</code> data from
                <code>Postgresql/Redis</code>
                Database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">1</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/profile/byuser
                </code>
            </td>
            <td><code>{String, Objects }</code></td>
            <td> <code>'/profile/byuser'</code> route will allow you to fetch <code>profile object</code> data from
                <code>Postgresql/Redis</code>
                Database using `userid`.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/profile/byid?id={}
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>Profile Object</code> based
                on <code>profileid</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/protected/profile/new
                </code>
            </td>
            <td><code>{String, Slices, Points, Number}</code></td>
            <td> <code>'/profile/new'</code> route will allow you to add user profile data to database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/protected/profile/update
                </code>
            </td>
            <td><code>{String, Slices, Points, Number}</code></td>
            <td> <code>'/profile/update'</code> route will allow you to update <code>Profile</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 https://gopherscom.herokuapp.com/protected/profile/resetcache
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>Profile</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Profile** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .


### Protected - BLOG
<code> GetBlogList, GetBlog, CreateBlog, UpdateBlog, SetBlogPublic, DeleteBlog, ResetBlogCache </code> : routes are available now at the moment and more
services/features will be added soon.


<table class="table table-hover">
    <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
            <th scope="col">Types</th>
            <th scope="col">Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">1</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/blog/list
                </code>
            </td>
            <td><code>{String, Objects, Slice}</code></td>
            <td> <code>'/blog/list'</code> route will allow you to fetach <code>slices of blog object</code> data from
                <code>Postgresql/Redis</code>
                Database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://gopherscom.herokuapp.com/protected/blog/byid?id={}
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>Blog Object</code> based
                on <code>id</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/protected/blog/new
                </code>
            </td>
            <td><code>{String, Slices, Points}</code></td>
            <td> <code>'/blog/new'</code> route will allow you to add Blog data to database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/protected/blog/update
                </code>
            </td>
            <td><code>{String, Slices, Points}</code></td>
            <td> <code>'/blog/update'</code> route will allow you to update <code>Blog</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                https://gopherscom.herokuapp.com/protected/blog/remove
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/blog/remove'</code> route will allow you to remove <code>Blog</code> with
                <code>{id}</code> from database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 https://gopherscom.herokuapp.com/protected/blog/resetcache
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>Blog</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Blog** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .




### Protected - APPTYPE
<code> GetApptypeList, GetApptype, CreateApptype, UpdateApptype, DeleteApptype, ResetApptypeCache </code> : routes are available now at the moment and more
services/features will be added soon.

var apptypeURL = "https://gopherscom.herokuapp.com/protected/apptype/"


<table class="table table-hover">
    <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Methods</th>
            <th scope="col">Path</th>
            <th scope="col">Types</th>
            <th scope="col">Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">1</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    apptype+"list"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'apptype/list'</code> route will allow you to fetch <code>slices of apptype object</code> data from
                <code>Postgresql/Redis</code>
                Database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    apptype+"byid?id={}"
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>Apptype Object</code> based
                on <code>apptypeid</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                apptype+"new"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/apptype/new'</code> route will allow you to add user apptype data to database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                apptype+"update"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/apptype/update'</code> route will allow you to update <code>Apptype</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 apptype+"resetcache"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>Apptype</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Apptype** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .