{{define "title"}}CV | Edit{{end}}
{{define "content"}}
    <style>
        p:empty:not(:focus)::before {
            content: attr(data-placeholder);
        }
    </style>
    {{template "cv_template" .}}
    <div class="container">
    <div class="row">
        <div class="col-md-2 mx-auto">
            <button type="button" id="submitButton" class="btn btn-primary">Save</button>
        </div>
    </div>
</div>
{{end}}
{{define "js"}}
<script>
    $(document).ready(function () {
        $('#submitButton').click(function () {
            $.ajax({
                type: 'POST',
                url: '/faculty/editcv',
                data: {
                    'bio': $('#bio').text(),
                    'about': $('#about').text()
                },
                success: function (msg) {}
            });
        });
    });
</script>
{{end}}