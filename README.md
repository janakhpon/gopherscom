# gopherscom

**gopherscom** is a Gophers coding solution, trend in go and knowledge sharing community especially for Golang developers. It is implemented using [Go](https://golang.org/) and **Go's** framework [Gin](https://github.com/gin-gonic) and [Postgresql](https://www.postgresql.org/) as persistant database. **gopherscom** provides a [RESTful Service](https://restfulapi.net/)
built using [TypeScript](https://www.typescriptlang.org/) and features like [Geojson](http://geojson.io/), [Email
auth](https://mail.google.com/), [JWT with refreshToken](http://www.passportjs.org/) and [Redis Performance Caching](https://redis.io/) are included currently and more will be
added soon.Here is a deployed link [HEROKU](https://gopherscom.herokuapp.com/)


##

##


## API LIST

### User
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

### Profile
addProfile, getProfile, getProfiles, updateProfile </code> : routes are available now at the moment and more
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
            <td><code>POST</code></td>
            <td>
                <code>
                    https://restletsplant.herokuapp.com/profile/
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to save new <code>Profile</code> data related to
                <code>User</code> to
                Database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                    https://restletsplant.herokuapp.com/profile/?userid={}id={}
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to update <code>Profile</code> with
                <code>{userid}</code> and <code>{id}</code> in database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://restletsplant.herokuapp.com/profile/?userid={}
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>Profile</code> based
                on <code>userid</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    https://restletsplant.herokuapp.com/profile/profiles
                </code>
            </td>
            <td><code>{Array, Objects}</code></td>
            <td> <code>'/profiles'</code> route will allow you to fetch <code>Profile</code> list within
                Database.
                <code>{ Authenitication with valid accessToken is required }</code> .</td>
        </tr>
    </tbody>
</table>

- It's recommended to provide <code>location</code> data in this formate :
<code>No.xxx || home no , xxx Street || Street Name, Myinetharyar || Ward Name, Mawlamyine || City Name, MON || State Name, 12011 || Zip Code, Myanmar || Country</code>,
.<code> Zip code</code> && <code>Country</code> is a must provide data here. <br /> After all data for
<code>location</code> is completed in <code>format</code> your address will be generatated in <code>geojson</code>
format included <code> [lat, lon] </code>.

## Plant
<code>addPlant, getPlant, updatePlant, removePlant, commentPlant, commentPlant, removeCommentPlant, updateCommentPlant </code>:
routes are available now at the moment and more services/features will be added soon. </p>
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
                                https://restletsplant.herokuapp.com/plant/
                                </code>
            </td>
            <td><code>{Array, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch plant list from
                database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/plant/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to add plant data to
                database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/plant/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to update <code>Plant</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/plant/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to delete <code>Plant</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td><code>POST</code></td>
            <td>
                <code>
                            https://restletsplant.herokuapp.com/plant/vote/
                            </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to vote and save in <code>Plant</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">6</th>
            <td><code>POST</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/plant/comment/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to comment and save in <code>Plant</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">7</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/plant/comment/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to update <code> Comment </code> with
                <code>{plantid}</code> and <code>{commentid}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">8</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/comment/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to delete <code> Comment </code> with
                <code>{plantid}</code> and <code>{commentid}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>
- It's recommended to provide specific data from <code>category</code> and <code>resource</code> model not to cause
<code>unstructured error</code>.

## Category
<code>addCategory, updateCategory, removeCategory, getCategories </code> : routes are available now at the moment and
more services/features will be added soon. </p>
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
                                https://restletsplant.herokuapp.com/category/
                                </code>
            </td>
            <td><code>{Array, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch category list from
                database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST, PUT</code></td>
            <td>
                <code>
                               https://restletsplant.herokuapp.com/category/
                               </code>
            </td>
            <td><code>{String}</code></td>
            <td><code>'/'</code> route will allow you to add category to
                database.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/category/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to update <code>Category</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                                https://restletsplant.herokuapp.com/category/
                                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to delete <code>Category</code> with
                <code>{id}</code> to database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>
- It's recommended to provide specific data from <code>category</code> and <code>resource</code> model not to cause
<code>unstructured error</code>

## Resource
- Resources service is not ready yet

