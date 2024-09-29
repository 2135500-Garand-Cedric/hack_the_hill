<?php
if (!isset($_SESSION['token'])) {
    header('Location: login.php');
}