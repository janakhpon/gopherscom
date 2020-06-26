### Protected - BOOTCAMP
<code> GetBootcampList, GetBootcamp, CreateBootcamp, UpdateBootcamp, SetBootcampAvailability, EnrollBootcamp, LikeBootcamp, CommentBootcamp, DeleteBootcamp, ResetBootcampCache </code> : routes are available now at the moment and more
services/features will be added soon.

##

```go
const bootcampURL = "https://gopherscom.herokuapp.com/protected/bootcamp/"

```
##


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
                    bootcampURL+"list"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point, Objects}</code></td>
            <td> <code>`bootcamp/list`</code> route will allow you to fetch <code>slices of bootcamp object</code> data from
                <code>Postgresql/Redis</code>
                bootcamp.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    bootcampURL+"byid?id={}"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point, Object}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>bootcamp Object</code> based
                on <code>bootcampid</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>POST</code></td>
            <td>
                <code>
                bootcampURL+"new"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point}</code></td>
            <td> <code>'/bootcamp/new'</code> route will allow you to add user bootcamp data to bootcamp.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                bootcampURL+"update"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point }</code></td>
            <td> <code>'/bootcamp/update'</code> route will allow you to update <code>bootcamp</code> with
                <code>{id}</code> to bootcamp.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                bootcampURL+"public"
                </code>
            </td>
            <td><code>{ String, Boolean }</code></td>
            <td> <code>'/bootcamp/public'</code> route will allow you to set the  <code>bootcamp availability</code> with
                <code>{id}</code> to bootcamp.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">6</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                bootcampURL+"remove"
                </code>
            </td>
            <td><code>{ String}</code></td>
            <td> <code>'/bootcamp/remove'</code> route will allow you to remove <code>Bootcamp</code> with
                <code>{id}</code> from database.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">7</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                bootcampURL+"enroll"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point, Object }</code></td>
            <td> <code>'/bootcamp/enroll'</code> route will allow you to enroll a <code>bootcamp</code> with
                <code>{id}s </code> to bootcamp.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">8</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                bootcampURL+"like"
                </code>
            </td>
            <td><code>{ String, Boolean }</code></td>
            <td> <code>'/bootcamp/like'</code> route will allow you to like the <code>bootcamp</code> with
                <code>{id}s</code> to bootcamp.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">9</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                bootcampURL+"comment"
                </code>
            </td>
            <td><code>{ String, Object }</code></td>
            <td> <code>'/bootcamp/comment'</code> route will allow you to comment the <code>bootcamp</code> with
                <code>{id}s</code> to bootcamp.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">10</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 bootcampURL+"resetcache"
                </code>
            </td>
            <td><code>{ String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>bootcamp</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Bootcamp** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .