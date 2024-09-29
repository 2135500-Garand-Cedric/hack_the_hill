const toggleBtn = document.getElementById('toggle-btn');
const sidebar = document.getElementById('sidebar');
const date = document.querySelector('.date');
const firstEntry = document.querySelector('.first-entry');
const secondEntry = document.querySelector('.second-entry');
const adviceEntry = document.querySelector('.advice-entry');
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
    adviceEntry.classList.toggle('hidden');
    datePicker.style.display = 'none';
});

document.getElementById('history-icon').addEventListener('click', function() {
    if ((datePicker.style.display === 'none' || datePicker.style.display === '') && !sidebar.classList.contains('collapsed')) {
        const iconPosition = historyIcon.getBoundingClientRect();
        datePicker.style.display = 'block';

        datePicker.style.position = 'absolute';
        datePicker.style.left = iconPosition.right + 'px';
        datePicker.style.top = (iconPosition.top + 15) + 'px';
        datePicker.focus();
    } else {
        datePicker.style.display = 'none';
    }
});

document.getElementById('date-picker').addEventListener('change', function() {
    const selectedDate = this.value;
    datePicker.style.display = 'none';
    let currentDate = new Date(selectedDate);
    currentDate.setHours(currentDate.getHours() + 4);
    dateDiv.innerHTML = formatDate(new Date(currentDate));

    const url = `get_journal_history.php?date=${encodeURIComponent(currentDate.toISOString().split('T')[0])}`;
    
    fetch(url)
        .then(response => response.json())
        .then(data => {
            const tasksArray = JSON.parse(data.data);
            if (data.error) {
                alert(data.error);
            } else {
                let html = "First Entry:<br />";
                tasksArray.forEach((data, index) => {
                    html += `<b>${data.task || 'N/A'}</b>: ${data.description || 'N/A'}<br />`;
                });
                firstEntry.innerHTML = html;
            }
        })
        .catch(error =>  {
            firstEntry.innerHTML = "First Entry:<br />";
            console.error('Error fetching the API:', error)
        });

    const url2 = `get_reflection_history.php?date=${encodeURIComponent(currentDate.toISOString().split('T')[0])}`;
    
    fetch(url2)
        .then(response => response.json())
        .then(data => {
            const tasksArray = JSON.parse(data.data);
            if (data.error) {
                alert(data.error);
            } else {
                let html = "Second Entry:<br />";
                tasksArray.forEach((data, index) => {
                    html += `<b>${data.task || 'N/A'}</b>: ${data.description || 'N/A'}<br />`;
                });
                secondEntry.innerHTML = html;
            }
        })
        .catch(error =>  {
            secondEntry.innerHTML = "Second Entry:<br />";
            console.error('Error fetching the API:', error)
        });

    const url3 = `get_advice_history.php?date=${encodeURIComponent(currentDate.toISOString().split('T')[0])}`;
    
    fetch(url3)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const tasksArray = JSON.parse(data.advice);
            if (data.error) {
                alert(data.error);
            } else {
                let html = "Advice:<br />";
                tasksArray.forEach((data, index) => {
                    html += `<b>${data.category || 'N/A'}</b>: ${data.advice || 'N/A'}<br />`;
                });
                adviceEntry.innerHTML = html;
            }
        })
        .catch(error =>  {
            adviceEntry.innerHTML = "Advice:<br />";
            console.error('Error fetching the API:', error)
        });
});

function formatDate(date) {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString(undefined, options);
}
const dateDiv = document.querySelector('.date');

const currentDate = new Date();

dateDiv.innerHTML = formatDate(currentDate);
