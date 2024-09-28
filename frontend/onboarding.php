<?php
require './include/configurations.php';
?>
<!-- https://freefrontend.com/css-login-forms/#google_vignette -->
<!DOCTYPE html>
<html>
<head>
	<title>Onboarding</title>
	<link rel="stylesheet" type="text/css" href="css/onboarding.css">
	<link href="https://fonts.googleapis.com/css2?family=Jost:wght@500&display=swap" rel="stylesheet">
	<link rel="stylesheet" href="https://netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css">
</head>
<body>
	<div id="snackbar"></div>
	<?php
		if (isset($_SESSION['snackbar_message'])) {
			$message = $_SESSION['snackbar_message'];
			echo "<input type='hidden' id='snackbar-message' value='${message}'>";
			unset($_SESSION['snackbar_message']);
		}
	?>
	<div class="main">  	
		<input type="checkbox" id="chk" aria-hidden="true">
		<div class="signup">
			<form action="index.php">
				<h2 for="chk" aria-hidden="true">Tell us about yourself!</h2>
				<label for="q1">Question 1?</label>
                <textarea name="q1" required="" rows="4"></textarea>
				<label for="q2">What are your hobbies?</label>
                <textarea name="q2" required="" rows="4"></textarea>
                <label for="q3">What are your interests?</label>
                <textarea name="q3" required="" rows="4"></textarea>
                <label for="q4">What is your occupation?</label>
                <textarea name="q4" required="" rows="4"></textarea>
                <label for="q5">In which city do you live? (City/Country)</label>
                <textarea name="q5" required="" rows="4"></textarea>
                <label for="q6">Gender?</label>
                <textarea name="q6" required="" rows="4"></textarea>
				<label for="q7">Date of birth?</label>
                <textarea name="q7" required="" rows="4"></textarea>
                <button>Submit</button>
			</form>
		</div>
	</div>
	<script src="js/snackbar.js"></script>
</body>
</html>