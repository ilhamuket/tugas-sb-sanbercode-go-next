/* eslint-disable react/prop-types */
import { useState, useEffect } from 'react';
import './TugasHooks.css';
import TugasIntroReact from '../Tugas-Intro-ReactJS/TugasIntroReact';

const TugasHooks = ({ onCountdownEnd }) => {
  const [currentTime, setCurrentTime] = useState(new Date().toLocaleTimeString());
  const [countdown, setCountdown] = useState(100);

  useEffect(() => {
    const interval = setInterval(() => {
      const now = new Date();
      const currentTime = now.toLocaleTimeString();
      const newCountdown = countdown - 1;

      setCurrentTime(currentTime);
      setCountdown(newCountdown);

      if (newCountdown <= 0) {
        onCountdownEnd(); 
        clearInterval(interval);
      }
    }, 1000);

    return () => clearInterval(interval);
  }, [countdown, onCountdownEnd]);

  return (
    <div>
      <div className="hooks-component">
      <h2>Now At - {currentTime}</h2>
      <p>Countdown: {countdown}</p>
      </div>
      <TugasIntroReact />
    </div>
  );
};

export default TugasHooks;