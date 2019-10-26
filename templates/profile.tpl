{{define "title"}}Profile{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        {{if .Data.Employee.Uname}}
            <div class="row">
                <div class="col-md-12 mx-auto">
                    <div class="card">
                        <div class="card-header">Basic Details</div>
                        <div class="card-body">
                            <div class="row">
                                <div class="col-md-4">
                                    <b>Username : </b>{{.User}}
                                </div>
                                <div class="col-md-4">
                                    <b>Name : </b>{{.Data.Employee.Fname}}
                                    {{.Data.Employee.Lname}}
                                </div>
                                <div class="col-md-4">
                                    <b>Email : </b>{{.Data.Employee.Email}}
                                </div>
                            </div>
                        </div>
                    </div><br><br>
                </div>
                {{if .Data.Faculty.Uname}}
                    <div class="col-md-12 mx-auto">
                        <div class="card">
                            <div class="card-header">Faculty Details</div>
                            <div class="card-body">
                                <div class="row">
                                    <div class="col-md-4">
                                        <b>Dept : </b>{{.Data.Faculty.Dept}}
                                    </div>
                                    <div class="col-md-4">
                                        <b>Start Date : </b>{{.Data.Faculty.StartDate}}
                                    </div>
                                    <div class="col-md-4">
                                        <b>End Date : </b>{{.Data.Faculty.EndDate}}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                {{end}}
            </div>
            <div class="row">
                <div class="container">
                    <br>
                    <!-- Nav tabs -->
                    <ul class="nav nav-tabs" role="tablist">
                        <li class="nav-item">
                            <a class="nav-link active" data-toggle="tab" href="#home">
                                <i class="fas fa-user"></i> Overview</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" data-toggle="tab" href="#menu1">
                                <i class="fas fa-cube"></i> Project</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" data-toggle="tab" href="#menu2">
                                <i class="fas fa-trophy"></i> Prizes</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" data-toggle="tab" href="#menu2">
                                <i class="far fa-newspaper"></i> Press/Media</a>
                        </li>
                    </ul>

                    <!-- Tab panes -->
                    <div class="tab-content">
                        <div id="home" class="container tab-pane active"><br>
                            <h3>HOME</h3>
                            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                        </div>
                        <div id="menu1" class="container tab-pane fade"><br>
                            <h3>Menu 1</h3>
                            <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
                        </div>
                        <div id="menu2" class="container tab-pane fade"><br>
                            <h3>Menu 2</h3>
                            <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam.</p>
                        </div>
                    </div>
                </div>
            </div>
        {{else}}
            <h3>Profile not found!</h3>
        {{end}}
    </div>
{{end}}