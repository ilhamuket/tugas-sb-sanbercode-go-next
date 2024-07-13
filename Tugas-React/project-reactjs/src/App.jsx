// import { useState } from 'react';
// import './App.css';
// import TugasIntroReact from './Tugas-Intro-ReactJS/TugasIntroReact';
// import TugasHooks from './Tugas-Hooks/TugasHooks';
import TugasCrudHooks from './Tugas-CRUD-Hooks/TugasCRUDHooks';

const App = () => {
  // const [showHooks, setShowHooks] = useState(true);

  // const handleCountdownEnd = () => {
  //   setShowHooks(false);
  // };

  return (
    <div className="App">
      {/* {showHooks && <TugasHooks onCountdownEnd={handleCountdownEnd} className="hooks-component" />}
      <div className="container">
        <TugasIntroReact />
      </div> */}
      <TugasCrudHooks />
    </div>
  );
};

export default App;