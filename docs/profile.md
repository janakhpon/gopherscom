

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