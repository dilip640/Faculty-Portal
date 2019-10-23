{{define "title"}}Register{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        <div class="row">
            <div class="col-md-6 mx-auto">
                <form method="POST" action="/register">
                    <div class="form-row">
                        <div class="form-group col-md-6">
                            <label for="exampleInputfirstname" id="exampleInputfirstname">First Name</label>
                            <input type="text" name="first_name" class="form-control"
                             placeholder="First name" required>
                        </div>
                        <div class="form-group col-md-6">
                            <label for="exampleInputlastname" id="exampleInputlastname">Last Name</label>
                            <input type="text" name="last_name" class="form-control"
                             placeholder="Last name" required>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="exampleInputusername">Username</label>
                        <input type="username" name="name" class="form-control"
                         id="exampleInputusername" aria-describedby="emailHelp"
                          placeholder="Username"required>
                    </div>
                    <div class="form-group">
                        <label for="exampleInputPassword1">Password</label>
                        <input type="password" name="password" class="form-control"
                         id="exampleInputPassword1" placeholder="Password" required>
                    </div>
                    <div class="form-group">
                        <label for="exampleInputEmail">Email</label>
                        <input type="email" name="email" class="form-control"
                         id="exampleInputEmail" required>
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