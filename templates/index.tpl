{{define "title"}}Home{{end}}
{{define "content"}}
    <div class="container" style="margin-top:50px">
        <div class="row">
            <div class="col-md-6 mx-auto">
                {{ if .User}}
                    <h2>Welome {{.User}}</h2>
                {{else}}
                    <p><a href="/login" >Click Here</a> to login or <a href="/register" >register</a></p>
                {{end}}
            </div>
        </div>
    </div>
{{end}}