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
            </div>
            {{if .Data.Faculty.Uname}}
                <div class="row">
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
                </div>
                <div class="row">
                    {{template "cv_template" .}}
                </div>
            {{end}}

        {{else}}
            <h3>Profile not found!</h3>
        {{end}}
    </div>
{{end}}