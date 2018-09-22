package public

var SignupSucces = `{{define "content"}}
<div class="center-screen">
Sign up successfull. Redirecting to home pagein 5seconds...
</div>
<script>
	setTimeout(function() {
		window.location.href = "/home";
	}, 5000)
</script>
{{end}}`
