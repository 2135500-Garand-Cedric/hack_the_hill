<?php
require './include/configurations.php';
header('Content-Type: application/json');

if (!isset($_GET['date'])) {
    echo json_encode(["error" => "No date provided in the URL."]);
    exit;
}

$value = $_GET['date'];

$api_url = "http://127.0.0.1:3000/api/journal?date=" . urlencode($value);
$token = $_SESSION['token'];

$ch = curl_init();

curl_setopt($ch, CURLOPT_URL, $api_url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $token,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo json_encode(['error' => curl_error($ch)]);
    curl_close($ch);
    exit;
}

curl_close($ch);

$data = json_decode($response, true);
echo json_encode($data);
