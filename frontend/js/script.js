let mediaRecorder;
let audioChunks = [];
let recognition;
let isRecognizing = false;

const recordBtn = document.getElementById('record-btn');
const waveContainerLeft = document.getElementById('wave-container-left');
const waveContainerRight = document.getElementById('wave-container-right');
const saveRecordingQuestion = document.getElementById('save-recording');
const saveBtn = document.getElementById('save-btn');
const discardBtn = document.getElementById('discard-btn');
const transcriptDiv = document.getElementById('transcript');
const micContainer = document.getElementById('mic-container');

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
    saveRecordingQuestion.classList.add('hidden');
    micContainer.classList.remove('hidden');
});

discardBtn.addEventListener('click', () => {
    discardRecording();
    saveRecordingQuestion.classList.add('hidden');
    micContainer.classList.remove('hidden');
});

function startRecording() {
    // Start media recording
    navigator.mediaDevices.getUserMedia({ audio: true })
        .then(stream => {
            mediaRecorder = new MediaRecorder(stream);
            mediaRecorder.start();
            audioChunks = [];

            mediaRecorder.addEventListener('dataavailable', event => {
                audioChunks.push(event.data);
            });

            // Start transcription (SpeechRecognition)
            startTranscription();

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
            // Stop transcription
            stopTranscription();

            // Hide the wavy lines and show the save/discard prompt
            waveContainerLeft.classList.add('hidden');
            waveContainerRight.classList.add('hidden');
            saveRecordingQuestion.classList.remove('hidden');
            micContainer.classList.add('hidden');
        });
    }
}

function saveRecording() {
    const transcript = transcriptDiv.innerHTML;
    console.log(transcript);

    // Define the URL of the PHP file
    const phpUrl = 'save_transcript.php'; // Your PHP file

    fetch(phpUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: `transcript=${encodeURIComponent(transcript)}`
    })
    .then(response => response.json())
    .then(data => {
        // Handle the PHP response
        console.log('Transcript saved successfully:', data);
    })
    .catch(error => {
        // Handle errors
        console.error('Error saving transcript:', error);
    });
}


let fullTranscript = '';  // Store the accumulated transcription

function discardRecording() {
    audioChunks = [];
    transcriptDiv.textContent = "";
    fullTranscript.textContent = "";
    fullTranscript = "";
}

function startTranscription() {
    if (!('webkitSpeechRecognition' in window)) {
        alert('Your browser does not support speech recognition.');
        return;
    }

    if (isRecognizing) return;

    // Initialize the SpeechRecognition object
    recognition = new webkitSpeechRecognition();
    recognition.continuous = true;
    recognition.interimResults = true;
    recognition.lang = 'en-US';

    // On result event, display real-time transcription
    recognition.onresult = (event) => {
        let interimTranscript = '';  // Interim result

        for (let i = event.resultIndex; i < event.results.length; i++) {
            let transcript = event.results[i][0].transcript;

            // Detect if a new sentence starts (based on capitalization of the first letter)
            if (transcript && transcript[0] === transcript[0].toUpperCase()) {
                // Add space between the previous sentence and the new one
                if (fullTranscript.length > 0 && !fullTranscript.endsWith(' ')) {
                    fullTranscript += ' ';
                }
            }

            if (event.results[i].isFinal) {
                fullTranscript += transcript.trim();  // Add final results to full transcript
            } else {
                interimTranscript += transcript;  // Handle interim results
            }
        }
        
        // Display full transcript + interim
        transcriptDiv.textContent = fullTranscript + interimTranscript;
    };

    recognition.onerror = (event) => {
        console.error('Speech recognition error:', event.error);
    };

    recognition.onend = () => {
        isRecognizing = false;
    };

    recognition.start();
    isRecognizing = true;
}

function stopTranscription() {
    if (recognition && isRecognizing) {
        recognition.stop();
        isRecognizing = false;
    }
}
