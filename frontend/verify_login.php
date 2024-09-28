<?php
require './include/configurations.php';
// Check if the form was submitted
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Retrieve form data
    $email = $_POST['email'];
    $password = $_POST['password'];

    // Initialize an array to hold errors
    $errors = "";

    // Validate form fields
    if (empty($email)) {
        $errors .= "Email is required.<br />";
    }
    if (empty($password)) {
        $errors .= "Password is required.<br />";
    }

    // Check if there are no errors
    if (empty($errors)) {
        // Prepare data for API request
        $data = [
            'email' => $email,
            'password' => $password,
        ];

        // API URL
        $apiUrl = 'http://127.0.0.1:3000/login'; // Replace with your API endpoint

        // Initialize cURL
        $ch = curl_init($apiUrl);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        // Set the data to send as form-data
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Type: multipart/form-data',
        ]);

        // Execute cURL request
        $response = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        // Close cURL
        curl_close($ch);

        // Check the response
        if ($httpCode === 200) {
            // Success, handle the response (you can decode JSON here if needed)
            $responseData = json_decode($response, true);
            // Do something with the response (e.g., redirect, display message)
            $_SESSION['token'] = $response['token'];
            header("Location: index.php");
        } else {
            // Handle error response
            $_SESSION['snackbar_message'] = 'An error occured with the login. <br />';
            header("Location: login.php");
        }
    } else {
        // Handle validation errors
        $_SESSION['snackbar_message'] = $errors;
        header("Location: login.php");
    }
} else {
    // Redirect or handle invalid request method
    header("Location: index.php");
    exit();
}
?>
