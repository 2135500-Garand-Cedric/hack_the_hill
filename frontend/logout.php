<?php
require 'include/configurations.php';
unset($_SESSION['token']);
header('Location: index.php');