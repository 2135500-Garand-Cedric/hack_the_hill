<?php
require './include/configurations.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $email = $_POST['email'];
    $password = $_POST['password'];
    $errors = "";

    if (empty($email)) {
        $errors .= "Email is required.<br />";
    }
    if (empty($password)) {
        $errors .= "Password is required.<br />";
    }

    if (empty($errors)) {
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
            $_SESSION['token'] = $response['token'];
            header("Location: index.php");
        } else {
            $_SESSION['snackbar_message'] = 'Incorrect username or password.<br />';
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
