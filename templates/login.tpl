{{define "title"}}Login{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        <div class="row">
            <div class="col-md-6 mx-auto">
                {{if .Data.Error}}
                    <ul>
                        <li style="color:red">{{.Data.Error}}</li>
                    </ul>
                {{end}}
                <form method="POST" action="/login">
                    <div class="form-group">
                        <label for="exampleInputEmail1">Username</label>
                        <input type="username" name="name" class="form-control"
                         id="exampleInputEmail1" aria-describedby="emailHelp"
                          placeholder="Username"required>
                    </div>
                    <div class="form-group">
                        <label for="exampleInputPassword1">Password</label>
                        <input type="password" name="password" class="form-control"
                         id="exampleInputPassword1" placeholder="Password" required>
                    </div>
                    <div class="form-group form-check">
                        <input type="checkbox" class="form-check-input"
                         id="exampleCheck1" required>
                        <label class="form-check-label" for="exampleCheck1">Check me out</label>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>
            </div>
        </div>
    </div>
{{end}}