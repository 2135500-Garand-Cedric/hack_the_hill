<?php
require './include/configurations.php';
$api_url = "http://127.0.0.1:3000/api/getsidebardata"; // Replace with your API URL
$token = $_SESSION['token'];

// Initialize cURL
$ch = curl_init($api_url);

// Set cURL options
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

// Execute the cURL request
$response = curl_exec($ch);

// Check for errors
if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
    $response = ""; // Clear response on error
}

// Close cURL session
curl_close($ch);

// You can decode the response if it's JSON
$data = json_decode($response, true);
print_r($data);
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Microphone Recorder</title>
    <link rel="stylesheet" href="css/styles.css">
    <link rel="stylesheet" href="css/snackbar.css">
    <link rel="stylesheet" href="css/sidebar.css">
    <link rel="stylesheet" href="https://netdna.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.css">
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

    <!-- Sidebar -->
    <div class="sidebar collapsed" id="sidebar">
        <div class="toggle-btn" id="toggle-btn"><i class="fa fa-bars"></i></div>
        <div class="sidebar-subcontainer">
            <div class="icon-icon-container">
                <div class="icon-container collapsed" id="logo-icon"><img src="./images/logo.png" alt="logo" class="logo"></div>
                <div class="icon-container collapsed" id="history-icon"><i class="fa fa-history history-icon"></i></div>
                <div class="icon-container collapsed" id="cog-icon"><i class="fa fa-cog cog-icon"></i></div>
            </div>
            <div class="information">
                <div class="username hidden">Hi [username]</div>
                <div class="date hidden"></div>
                <div class="first-entry hidden">First Entry:</div>

            </div>
        </div>
    </div>

    <header>
        <a href="login.php">
            <button id="login-btn">Login</button>
        </a>
    </header>

    <main>
        <textarea id="transcript"></textarea>
        <div class="mic-container" id="mic-container">
            <div id="wave-container-left" class="wave-container-left hidden">
                <?php 
                    for ($i = 0; $i < 20; $i++) {
                        echo "<div class='wave' id='wave-left-${i}'></div>";
                    }
                ?>
            </div>
            <div class="button-container">
                <button id="record-btn" class="mic-btn">
                    <!-- <a href="https://www.flaticon.com/free-icons/microphone" title="microphone icons">Microphone icons created by Uniconlabs - Flaticon</a> -->
                    <img src="images/microphone.png" alt="Microphone Icon" class="mic-icon">
                </button>
            </div>
            <div id="wave-container-right" class="wave-container-right hidden">
                <?php 
                    for ($i = 0; $i < 20; $i++) {
                        echo "<div class='wave' id='wave-right-${i}'></div>";
                    }
                ?>
            </div>
        </div>
        <div id="save-recording" class="save-recording hidden">
            <div class="save-recording-content">
                <p>Do you want to save the recording?</p>
                <button id="save-btn">Save</button>
                <button id="discard-btn">Discard</button>
            </div>
        </div>
    </main>

    <script src="js/snackbar.js"></script>
    <script src="js/script.js"></script>
    <script src="js/sidebar.js"></script>
</body>
</html>
