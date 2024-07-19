/* eslint-disable react/prop-types */
import logo from '../assets/logo.png'; 
import './TugasIntroReact.css';

const Checkbox = ({ label }) => {
  return (
    <div className="checkbox-container">
      <input type="checkbox" id={label} name={label} />
      <label htmlFor={label}>{label}</label>
    </div>
  );
};

const TugasIntroReact = () => {
  const thingsToDo = [
    "Belajar GIT & CLI", 
    "Belajar HTML & CSS", 
    "Belajar Javascript", 
    "Belajar ReactJS Dasar", 
    "Belajar ReactJS Advance"
  ];

  return (
    <div className="intro-react">
      <img src={logo} alt="sanbercode logo" className="logo" />
      <h2>THINGS TO DO</h2>
      <p>During bootcamp in sanbercode</p>
      {thingsToDo.map((item, index) => (
        <Checkbox key={index} label={item} />
      ))}
      <button type="button">SEND</button>
    </div>
  );
};

export default TugasIntroReact;