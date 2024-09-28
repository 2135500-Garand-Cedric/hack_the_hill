let mediaRecorder;
let audioChunks = [];

const recordBtn = document.getElementById('record-btn');
const waveContainerLeft = document.getElementById('wave-container-left');
const waveContainerRight = document.getElementById('wave-container-right');
const modal = document.getElementById('modal');
const saveBtn = document.getElementById('save-btn');
const discardBtn = document.getElementById('discard-btn');

let min = 50;
let max = 100;

for (let i = 0; i < 20; i++) {
    const wave = document.getElementById('wave-right-' + i);
    min = 50 - (2 * i);
    max = 100 - (4 * i);
    const randomHeight = Math.floor(Math.random() * (max - min + 1)) + min;
    wave.style.height = randomHeight + 'px';
}

for (let i = 0; i < 20; i++) {
    k = 19 - i;
    const wave = document.getElementById('wave-left-' + k);
    min = 50 - (2 * i);
    max = 100 - (4 * i);
    const randomHeight = Math.floor(Math.random() * (max - min + 1)) + min;
    wave.style.height = randomHeight + 'px';
}
            

// Add event listeners for both mouse and touch events
recordBtn.addEventListener('mousedown', startRecording);
recordBtn.addEventListener('mouseup', stopRecording);

recordBtn.addEventListener('touchstart', startRecording);
recordBtn.addEventListener('touchend', stopRecording);

saveBtn.addEventListener('click', () => {
    saveRecording();
    modal.classList.add('hidden');
});

discardBtn.addEventListener('click', () => {
    discardRecording();
    modal.classList.add('hidden');
});

function startRecording() {
    navigator.mediaDevices.getUserMedia({ audio: true })
        .then(stream => {
            mediaRecorder = new MediaRecorder(stream);
            mediaRecorder.start();
            audioChunks = [];

            mediaRecorder.addEventListener('dataavailable', event => {
                audioChunks.push(event.data);
            });

            // Show the wavy lines when recording starts
            waveContainerLeft.classList.remove('hidden');
            waveContainerRight.classList.remove('hidden');
        })
        .catch(error => {
            console.error('Error accessing microphone:', error);
        });
}

function stopRecording() {
    if (mediaRecorder) {
        mediaRecorder.stop();

        mediaRecorder.addEventListener('stop', () => {
            // Hide the wavy lines and show the save/discard prompt
            waveContainerLeft.classList.add('hidden');
            waveContainerRight.classList.add('hidden');
            modal.classList.remove('hidden');
        });
    }
}

function saveRecording() {
    const audioBlob = new Blob(audioChunks, { type: 'audio/wav' });
    const audioUrl = URL.createObjectURL(audioBlob);
    const a = document.createElement('a');
    a.href = audioUrl;
    a.download = 'recording.wav';
    a.click();
}

function discardRecording() {
    audioChunks = [];
}
