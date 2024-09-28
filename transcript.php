<?php
if ($_SERVER['REQUEST_METHOD'] == 'POST' && isset($_FILES['audio'])) {
    $filePath = $_FILES['audio']['tmp_name'];

    // API URL
    $apiUrl = 'http://127.0.0.1:5000/transcribe';

    // Initialize cURL session
    $ch = curl_init($apiUrl);

    // Set cURL options
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, [
        'file' => new CURLFile($filePath, 'audio/wav', $_FILES['audio']['name'])
    ]);

    // Execute cURL session and get the response
    $response = curl_exec($ch);

    // Close cURL session
    curl_close($ch);

    // Show the transcription
    $transcription = json_decode($response, true);
    if (isset($transcription['transcript'])) {
        echo "Transcript: " . $transcription['transcript'];
    } else {
        echo "Error: " . $transcription['error'];
    }
}