### Protected - BRANCH
<code> GetCompanyBranches, GetBranch, AddCompanyBranch, UpdateCompanyBranch, DeleteCompanyBranch, ResetBranchCache </code> : routes are available now at the moment and more services/features will be added soon.

##

```go

const branchURL = "https://gopherscom.herokuapp.com/public/branch/"
const privatebranchURL = "https://gopherscom.herokuapp.com/protected/branch/"

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
                    branchURL+"branches"
                </code>
            </td>
            <td><code>{ String, Float, Time, Branch  }</code></td>
            <td> <code>`branch/branches`</code> route will allow you to fetch <code>slices of branch object</code> data from <code>Postgresql/Redis</code> using <code> { company_id }</code>. </td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td><code>POST</code></td>
            <td>
                <code>
                branchURL+"new"
                </code>
            </td>
            <td><code>{ String, Float, Time, Boolean, Point}</code></td>
            <td> <code>'/branch/new'</code> route will allow you to publish a company `branch` on `GophersCom`.</td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td><code>PUT</code></td>
            <td>
                <code>
                privatebranchURL+"update"
                </code>
            </td>
            <td><code>{ String, Object }</code></td>
            <td> <code>'/branch/update'</code> route will allow you to update the <code>branch</code> with related <code>{id}</code>. <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                privatebranchURL+"delete"
                </code>
            </td>
            <td><code>{ String, Object }</code></td>
            <td> <code>'/branch/update'</code> route will allow you to delete the <code>branch</code> with related <code>{id}</code>.<code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td><code>DELETE</code></td>
            <td>
                <code>
                 privatebranchURL+"resetcache"
                </code>
            </td>
            <td><code>{ String }</code></td>
            <td> <code>'/'</code> route will allow you to reset all <code>branch</code> data in 
                <code>Redis cache</code>.
                <code>{ Authenitication with valid accessToken is required }</code> . </td>
        </tr>
    </tbody>
</table>

- It's recommended not to reset **Cache Data** for **Company Branch** if the request doesn't stuck in `errors` or `overcached`. As a result of removing cache, data fetching from **Postgresql** will be a bit slower than from retrieving from **Cache** .