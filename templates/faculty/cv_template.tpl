{{define "cv_template"}}
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <br>
                <!-- Nav tabs -->
                <ul class="nav nav-tabs" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link active" data-toggle="tab" href="#overview">
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
                    {{if .Data.Self}}
                        <li class="ml-auto">
                            <a class="nav-link" href="/faculty/editcv">
                                <i class="fas fa-edit"></i>Edit</a>
                        </li>
                    {{end}}
                </ul>

                <!-- Tab panes -->
                <div class="tab-content">
                    <div id="overview" class="container tab-pane active"><br>
                        <div >
                            <h3>Biography</h3>
                            <p data-placeholder="Describe yourself..." id="bio" contenteditable="{{ .Data.Style.Edit }}" class="{{ .Data.Style.Class }}">
                                {{ .Data.CVDetail.Overview.Biography}}</p>
                        </div>
                        <h4>About Me</h4>
                        <p data-placeholder="Write about you..." id="about" contenteditable="{{ .Data.Style.Edit }}" class="{{ .Data.Style.Class }}">
                            {{ .Data.CVDetail.Overview.AboutMe}}</p>
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
    </div>
{{end}}