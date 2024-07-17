// import { useState } from 'react';
// import './App.css';
// import TugasIntroReact from './Tugas-Intro-ReactJS/TugasIntroReact';
// import TugasHooks from './Tugas-Hooks/TugasHooks';
// import TugasCrudHooks from './Tugas-CRUD-Hooks/TugasCRUDHooks';
// import TugasAxios from './TugasAxios/TugasAxios';
import TugasContext from './TugasContext/TugasContext';

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
      {/* <TugasCrudHooks /> */}
      {/* <TugasAxios /> */}
      <TugasContext />
    </div>
  );
};

export default App;