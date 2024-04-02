import React, { useEffect, useState } from 'react';
import { Utterance } from './types';
import './App.css';
import userProfileIcon from './no-profile-picture-15260.svg';

const App = () => {
  const [transcript, setTranscript] = useState<Utterance[]>([]);

  useEffect(() => {
    const apiUrl = `${process.env.REACT_APP_BACKEND_URL}/api/transcript`;
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => {
        if (data.utterances) {
          setTranscript(data.utterances);
        }
      })
      .catch((error) => console.error('Fetching error:', error));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        {transcript.map((utterance, index) => (
          <div key={index} className="Utterance">
            <img src={userProfileIcon} alt="Profile Icon" className="ProfileIcon" />
            <div className="TextWithTimestamp">
              <strong className="Speaker">{utterance.speaker}</strong>
              <p className="Text">{utterance.text}</p>
              {/* Update the Timestamp rendering here */}
              <div className="Timestamp">
                <span className="Date">{new Date(utterance.timestamp).toLocaleDateString()}</span>
                <span className="Time">{new Date(utterance.timestamp).toLocaleTimeString()}</span>
              </div>
            </div>
          </div>
        ))}
      </header>
    </div>
  );
};

export default App;
