<?php
require './include/configurations.php';
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $transcript = $_POST['transcript'] ?? '';

    if (!empty($transcript)) {
        $api_url = 'http://127.0.0.1:3000/api/createjournalentry';

        $ch = curl_init($api_url);

        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        
        $data = http_build_query([
            'data' => $transcript
        ]);

        $token = $_SESSION['token'];

        curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Type: application/x-www-form-urlencoded',
            'Authorization: Bearer ' . $token
        ]);

        // Execute the cURL request
        $response = curl_exec($ch);
        $http_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        // Close the cURL session
        curl_close($ch);

        // Send the response back to the JavaScript
        if ($http_code === 200) {
            echo json_encode(['status' => 'success', 'message' => 'Transcript saved successfully']);
        } else {
            echo json_encode(['status' => 'error', 'message' => 'Failed to save transcript']);
        }
    } else {
        echo json_encode(['status' => 'error', 'message' => 'Transcript is empty']);
    }
} else {
    // Invalid request method
    echo json_encode(['status' => 'error', 'message' => 'Invalid request method']);
}
?>
