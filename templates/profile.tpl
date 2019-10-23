{{define "title"}}Profile{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        <div class="row">
            <div class="col-md-6 mx-auto">
            {{.User}}
            </div>
        </div>
    </div>
{{end}}