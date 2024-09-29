// Toggle sidebar function
const toggleBtn = document.getElementById('toggle-btn');
const sidebar = document.getElementById('sidebar');
// const mainContent = document.getElementById('main-content');
const header = document.querySelector('header');
const historyIcon = document.getElementById('history-icon');
const cogIcon = document.getElementById('cog-icon');
const logoIcon = document.getElementById('logo-icon');




toggleBtn.addEventListener('click', function() {
    sidebar.classList.toggle('collapsed');
    // mainContent.classList.toggle('collapsed');
    header.classList.toggle('collapsed');

    historyIcon.classList.toggle('collapsed');
    logoIcon.classList.toggle('collapsed');
    cogIcon.classList.toggle('collapsed');
});