

### Protected - TAG
<code> GetTagList, GetTag, CreateTag, UpdateTag, DeleteTag, ResetTagCache </code> : routes are available now at the moment and more
services/features will be added soon.

##

```go
const tagURL = "https://gopherscom.herokuapp.com/protected/tag/"

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
                    tag+"list"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'tag/list'</code> route will allow you to fetch <code>slices of tag object</code> data from
                <code>Postgresql/Redis</code>
                tag.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>GET</code></td>
            <td>
                <code>
                    tag+"byid?id={}"
                </code>
            </td>
            <td><code>{String, Objects}</code></td>
            <td> <code>'/'</code> route will allow you to fetch only specific <code>tag Object</code> based
                on <code>tagid</code>.<code>{ refreshToken need to be valid }</code> </td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>POST</code></td>
            <td>
                <code>
                tag+"new"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/tag/new'</code> route will allow you to add user tag data to tag.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                tag+"update"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/tag/update'</code> route will allow you to update <code>tag</code> with
                <code>{id}</code> to tag.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 tag+"resetcache"
                </code>
            </td>
            <td><code>{String}</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>tag</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Tag** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .