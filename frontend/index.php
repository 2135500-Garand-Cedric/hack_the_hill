<?php
require './include/configurations.php';
require './include/header.php';


$api_url = "http://127.0.0.1:3000/api/getsidebardata";
$token = $_SESSION['token'];

$ch = curl_init($api_url);

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
    $response = "";
}

curl_close($ch);

$outerData = json_decode($response, true);
$htmlFirstEntry = "";

if (!empty($outerData['data'])) {
    $innerJsonString = $outerData['data'];

    $tasks = json_decode($innerJsonString, true);

    foreach ($tasks as $task) {
        $taskName = $task['task'];
        $description = $task['description'];

        $htmlFirstEntry .= "<b>${taskName}</b>: ${description}<br />";
    }
}

########################################
# Second entry
########################################

$api_url = "http://127.0.0.1:3000/api/getreflectiondata";

$ch = curl_init($api_url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
    $response = "";
}

curl_close($ch);

$outerData = json_decode($response, true);
$htmlSecondEntry = "";

if (!empty($outerData['data'])) {
    $innerJsonString = $outerData['data'];

    $tasks = json_decode($innerJsonString, true);

    foreach ($tasks as $task) {
        $taskName = $task['task'];
        $description = $task['description'];
        $htmlSecondEntry .= "<b>${taskName}</b>: ${description}<br />";
    }
}

########################################
# Profile Info
########################################

$api_url = "http://127.0.0.1:3000/api/profile";

$ch = curl_init($api_url);

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
    $response = "";
}

curl_close($ch);

$data = json_decode($response, true);
$username = isset($data['username']) ? $data['username'] : "Guest";

########################################
# Get Advice
########################################

$api_url = "http://127.0.0.1:3000/api/getadvice";

$ch = curl_init($api_url);

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
    $response = "";
}

curl_close($ch);

$advice = json_decode($response, true);
$jsonString = $advice['advice'];
$adviceArray = json_decode($jsonString, true);
$adviceHtml = "";

foreach ($adviceArray as $index => $item) {
    $adviceText = $item['advice'];
    $category = $item['category'];
    $adviceHtml .= "<b>${category}: </b>${adviceText}.<br />";
}

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
                <input type="date" id="date-picker" class="hidden" style="display: none;">
                <div class="icon-container collapsed" id="cog-icon"><a href="logout.php"><i class="fa fa-cog cog-icon"></i></a></div>
            </div>
            <div class="information">
                <div class="date hidden" id="date"></div>
                <div class="first-entry hidden">First Entry: <br /><?php echo $htmlFirstEntry; ?></div>
                <div class="second-entry hidden">Second Entry: <br /><?php echo $htmlSecondEntry; ?></div>
            </div>
        </div>
    </div>

    <main>
        <div class="advice-title"><b>Hi <?php echo $username ?>, today's suggestion is:</b></div>
        <div class="advice"><?php echo $adviceHtml; ?></div>
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
