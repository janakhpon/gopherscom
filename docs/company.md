### Protected - COMPANY
<code> GetCompanyList, AddCompany, ResetCompanyCache </code> : routes are available now at the moment and more
services/features will be added soon.

##

```go

const companyURL = "https://gopherscom.herokuapp.com/public/company/"
const privateCompanyURL = "https://gopherscom.herokuapp.com/protected/company/"

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
                    companyURL+"list"
                </code>
            </td>
            <td><code>{ String, Float, Time, Branch  }</code></td>
            <td> <code>`company/list`</code> route will allow you to fetch <code>slices of company object</code> data from
                <code>Postgresql/Redis</code>
                company.<code>{ Authenitication with valid accessToken is required }</code> </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                companyURL+"new"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point}</code></td>
            <td> <code>'/company/new'</code> route will allow you to publish a `company` on `GophersCom`.</td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                privateCompanyURL+"update"
                </code>
            </td>
            <td><code>{ String, Object }</code></td>
            <td> <code>'/company/update'</code> route will allow you to update the <code>company</code> with related <code>{id}</code>. <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                privateCompanyURL+"remove"
                </code>
            </td>
            <td><code>{ String, Object }</code></td>
            <td> <code>'/company/update'</code> route will allow you to delete the <code>company</code> with related <code>{id}</code>.<code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 privateCompanyURL+"resetcache"
                </code>
            </td>
            <td><code>{ String }</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>company</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Company** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .