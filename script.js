let isRecording = false;
let mediaRecorder;
let audioChunks = [];

const recordBtn = document.getElementById('record-btn');
const modal = document.getElementById('modal');
const saveBtn = document.getElementById('save-btn');
const discardBtn = document.getElementById('discard-btn');

recordBtn.addEventListener('click', () => {
    if (!isRecording) {
        startRecording();
    } else {
        stopRecording();
    }
});

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

            mediaRecorder.addEventListener('stop', () => {
                modal.classList.remove('hidden');
            });
        })
        .catch(error => {
            console.error('Error accessing microphone:', error);
        });

    isRecording = true;
}

function stopRecording() {
    if (mediaRecorder) {
        mediaRecorder.stop();
    }
    isRecording = false;
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
