function showSnackbar(message) {
    console.log("test");
    var snackbar = document.getElementById("snackbar");
    snackbar.innerHTML = message;
    snackbar.className = "show";
    setTimeout(function() {
        snackbar.className = snackbar.className.replace("show", "");
    }, 3000);
}

document.addEventListener('DOMContentLoaded', function() {
    var message = document.getElementById("snackbar-message");
    if (message) {
        showSnackbar(message.value);
    }
});