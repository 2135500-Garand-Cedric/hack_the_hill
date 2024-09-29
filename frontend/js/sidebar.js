// Toggle sidebar function
const toggleBtn = document.getElementById('toggle-btn');
const sidebar = document.getElementById('sidebar');
const date = document.querySelector('.date');
const firstEntry = document.querySelector('.first-entry');
const secondEntry = document.querySelector('.second-entry');
const historyIcon = document.getElementById('history-icon');
const cogIcon = document.getElementById('cog-icon');
const logoIcon = document.getElementById('logo-icon');
const datePicker = document.getElementById('date-picker');

toggleBtn.addEventListener('click', function() {
    sidebar.classList.toggle('collapsed');
    historyIcon.classList.toggle('collapsed');
    logoIcon.classList.toggle('collapsed');
    cogIcon.classList.toggle('collapsed');
    date.classList.toggle('hidden');
    firstEntry.classList.toggle('hidden');
    secondEntry.classList.toggle('hidden');
    datePicker.style.display = 'none';
});

document.getElementById('history-icon').addEventListener('click', function() {
    if ((datePicker.style.display === 'none' || datePicker.style.display === '') && !sidebar.classList.contains('collapsed')) {
        // Get the position of the history icon
        const iconPosition = historyIcon.getBoundingClientRect();

        // Display the date picker
        datePicker.style.display = 'block';

        // Position the date picker to the right of the icon
        datePicker.style.position = 'absolute';
        datePicker.style.left = iconPosition.right + 'px';  // Position to the right
        datePicker.style.top = (iconPosition.top + 15) + 'px';     // Align top with icon

        // Focus on the date picker to automatically show the calendar
        datePicker.focus();
    } else {
        datePicker.style.display = 'none';
    }
});

document.getElementById('date-picker').addEventListener('change', function() {
    // Get the selected date
    const selectedDate = this.value;
    datePicker.style.display = 'none';

    // Do something with the selected date
    console.log('Selected date:', selectedDate);

    // You can add any JS logic here, like calling a function or updating the DOM
    alert('Date selected: ' + selectedDate);
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
