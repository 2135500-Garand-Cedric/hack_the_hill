from flask import Flask, request, jsonify
import speech_recognition as sr

app = Flask(__name__)

@app.route('/transcribe', methods=['POST'])
def transcribe_audio():
    # Check if the file is in the request
    if 'file' not in request.files:
        return jsonify({"error": "No file part in the request"}), 400

    file = request.files['file']

    # Initialize recognizer
    recognizer = sr.Recognizer()

    try:
        # Use SpeechRecognition to transcribe the audio file
        with sr.AudioFile(file) as source:
            audio_data = recognizer.record(source)
            transcript = recognizer.recognize_google(audio_data)
            return jsonify({"transcript": transcript})

    except sr.UnknownValueError:
        return jsonify({"error": "Audio could not be understood"}), 400
    except sr.RequestError:
        return jsonify({"error": "Error with transcription service"}), 500

if __name__ == '__main__':
    app.run(debug=True)
