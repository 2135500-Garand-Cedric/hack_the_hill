<?php
// Check if the form was submitted
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Retrieve form data
    $username = $_POST['username'];
    $email = $_POST['email'];
    $password = $_POST['password'];
    $confirmPassword = $_POST['confirm-password'];

    // Initialize an array to hold errors
    $errors = [];

    // Validate form fields
    if (empty($username)) {
        $errors['username'] = "Name is required.";
    }
    if (empty($email)) {
        $errors['email'] = "Email is required.";
    }
    if (empty($password)) {
        $errors['password'] = "Password is required.";
    }
    if (empty($confirmPassword)) {
        $errors['confirm_password'] = "Confirm password is required.";
    } else if ($password !== $confirmPassword) {
        $errors['confirm_password'] = "Passwords do not match.";
    }

    // Check if there are no errors
    if (empty($errors)) {
        // Prepare data for API request
        $data = [
            'email' => $email,
            'password' => $password,
            'username' => $username,
        ];

        // API URL
        $apiUrl = 'http://127.0.0.1:3000/register'; // Replace with your API endpoint

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
            header("Location: onboarding.html");
        } else {
            // Handle error response
            echo "Error occurred: " . $response;
        }
    } else {
        // Handle validation errors
        foreach ($errors as $field => $error) {
            echo "<p>$field: $error</p>";
        }
    }
} else {
    // Redirect or handle invalid request method
    header("Location: index.php");
    exit();
}
?>
