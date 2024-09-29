<?php
require './include/configurations.php';

if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    $goal = $_POST['q1'];
    $hobbies = $_POST['q2'];
    $interests = $_POST['q3'];
    $occupation = $_POST['q4'];
    $city_country = $_POST['q5'];
    $gender = $_POST['q6'];
    $dob = $_POST['q7'];

    $errors = "";

    if (empty($goal)) $errors .= "Goal is required.<br />";
    if (empty($hobbies)) $errors .= "Hobbies are required.<br />";
    if (empty($interests)) $errors .= "Interests are required.<br />";
    if (empty($occupation)) $errors .= "Occupation is required.<br />";
    if (empty($city_country)) $errors .= "City and Country are required.<br />";
    if (empty($gender)) $errors .= "Gender is required.<br />";
    if (empty($dob)) $errors .= "Date of Birth is required.<br />";

    if (empty($errors)) {
        $api_url = "http://127.0.0.1:3000/api/generateprofile";
        $data = [
            'goals' => $goal,
            'hobbies' => $hobbies,
            'interests' => $interests,
            'occupation' => $occupation,
            'city' => $city_country,
            'gender' => $gender,
            'dob' => $dob
        ];

        $ch = curl_init($api_url);
        $token = $_SESSION['token'];
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data));
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $token,
            'Content-Type: application/x-www-form-urlencoded'
        ]);

        $response = curl_exec($ch);
        $http_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        curl_close($ch);

        if ($http_code === 200) {
            $responseData = json_decode($response, true);
            $_SESSION['snackbar_message'] = $responseData['message'];
            header("Location: index.php");
        } else {
            $_SESSION['snackbar_message'] = "An error occured with the onboarding.<br />";
            header("Location: onboarding.php");
        }
    } else {
        $_SESSION['snackbar_message'] = $errors;
        header("Location: onboarding.php");
    }
} else {
    header("Location: index.php");
    exit();
}
?>
