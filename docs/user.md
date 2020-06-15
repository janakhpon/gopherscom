
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