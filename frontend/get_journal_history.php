<?php
require './include/configurations.php';
header('Content-Type: application/json');

// Ensure the value is passed in the URL
if (!isset($_GET['date'])) {
    echo json_encode(["error" => "No date provided in the URL."]);
    exit;
}

// Retrieve the value from the URL
$value = $_GET['date'];

// Set the API endpoint (replace with the actual API URL)
$api_url = "http://127.0.0.1:3000/api/journal?date=" . urlencode($value);

// Set the Bearer token (replace with your actual token)
$token = $_SESSION['token'];

// Initialize cURL
$ch = curl_init();

// Set the URL
curl_setopt($ch, CURLOPT_URL, $api_url);

// Set the request method to GET
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

// Set the Authorization header with Bearer token
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

// Execute the cURL request
$response = curl_exec($ch);

// Check if there were any errors
if (curl_errno($ch)) {
    echo json_encode(['error' => curl_error($ch)]);
    curl_close($ch);
    exit;
}

// Close the cURL session
curl_close($ch);

$data = json_decode($response, true);

// Return the JSON response
echo json_encode($data);
