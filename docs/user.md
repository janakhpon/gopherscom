
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
            <th scope="row">3</th>
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