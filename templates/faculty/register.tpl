{{define "title"}}Faculty | Update{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        <div class="row">
            <div class="col-md-6 mx-auto">
                {{if .Data.Error}}
                    <ul>
                        <li style="color:red">{{.Data.Error}}</li>
                    </ul>
                {{end}}
                <form method="POST" action="/faculty/update">
                    <div class="form-row">
                        <div class="form-group col-md-6">
                            <label for="exampleInputstartdate" id="exampleInputstartdate">Start Date</label>
                            <input type="date" name="start_date" id="exampleInputstartdate"
                             class="form-control"required>
                        </div>
                        <div class="form-group col-md-6">
                            <label for="exampleInputenddate" id="exampleInputenddate">End Date</label>
                            <input type="date" name="end_date" id="exampleInputenddate"
                             class="form-control"required>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="exampleInputdepartment" id="exampleInputdepartment">Department</label>
                        <input type="text" name="dept" class="form-control"
                         id="exampleInputdepartment" aria-describedby="emailHelp"
                          placeholder="e.g CSE"required>
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