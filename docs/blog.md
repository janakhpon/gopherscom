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

