<?php
require './include/configurations.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $username = $_POST['username'];
    $email = $_POST['email'];
    $password = $_POST['password'];
    $confirmPassword = $_POST['confirm-password'];

    $errors = "";

    if (empty($username)) {
        $errors .= "Name is required.<br />";
    }
    if (empty($email)) {
        $errors .= "Email is required.<br />";
    }
    if (empty($password)) {
        $errors .= "Password is required.<br />";
    }
    if (empty($confirmPassword)) {
        $errors .= "Confirm password is required.<br />";
    } else if ($password !== $confirmPassword) {
        $errors .= "Passwords do not match.<br />";
    }

    if (empty($errors)) {
        $data = [
            'email' => $email,
            'password' => $password,
            'username' => $username,
        ];

        $apiUrl = 'http://127.0.0.1:3000/register';

        $ch = curl_init($apiUrl);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Type: multipart/form-data',
        ]);

        $response = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        curl_close($ch);

        if ($httpCode === 200) {
            $responseData = json_decode($response, true);
            $data = [
                'email' => $email,
                'password' => $password,
            ];
    
            $apiUrl = 'http://127.0.0.1:3000/login';

            $ch = curl_init($apiUrl);
            curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt($ch, CURLOPT_POST, true);
            curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
            curl_setopt($ch, CURLOPT_HTTPHEADER, [
                'Content-Type: multipart/form-data',
            ]);

            $response = curl_exec($ch);
            $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
            curl_close($ch);
    
            if ($httpCode === 200) {
                $responseData = json_decode($response, true);
                $_SESSION['token'] = $responseData['token'];
                header("Location: onboarding.php");
            } else {
                $_SESSION['snackbar_message'] = "An error occured during sign up.<br />";
                header("Location: login.php");
            }
        } else {
            $_SESSION['snackbar_message'] = "An error occured during sign up.<br />";
            header("Location: login.php");
        }
    } else {
        $_SESSION['snackbar_message'] = $errors;
        header("Location: login.php");
    }
} else {
    header("Location: index.php");
    exit();
}
?>
