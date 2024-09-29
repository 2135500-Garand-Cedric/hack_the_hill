// Toggle sidebar function
const toggleBtn = document.getElementById('toggle-btn');
const sidebar = document.getElementById('sidebar');
const header = document.querySelector('header');
const username = document.querySelector('.username');
const date = document.querySelector('.date');
const firstEntry = document.querySelector('.first-entry');
const historyIcon = document.getElementById('history-icon');
const cogIcon = document.getElementById('cog-icon');
const logoIcon = document.getElementById('logo-icon');




toggleBtn.addEventListener('click', function() {
    sidebar.classList.toggle('collapsed');
    header.classList.toggle('collapsed');
    historyIcon.classList.toggle('collapsed');
    logoIcon.classList.toggle('collapsed');
    cogIcon.classList.toggle('collapsed');
    username.classList.toggle('hidden');
    date.classList.toggle('hidden');
    firstEntry.classList.toggle('hidden');
});

// Function to format the date
function formatDate(date) {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString(undefined, options);
}
// Get the current date
const currentDate = new Date();
// Get the date div
const dateDiv = document.querySelector('.date');
// Set the inner HTML of the date div to the current date
dateDiv.innerHTML = formatDate(currentDate);
