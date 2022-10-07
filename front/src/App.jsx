import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";

function App() {
  const [input, setInput] = useState("");

  return (
    <div className="App">
      <div>
        <img src={reactLogo} className="logo react" alt="React logo" />
      </div>
      <h1>Text lié à l'image</h1>
      <div className="card">
        <input
          value={input}
          onChange={(e) => {
            setInput(e.target.value);
          }}
        />
        <button
          onClick={() => {
            //call api
          }}
        >
          Déploiement
        </button>
      </div>
    </div>
  );
}

export default App;
