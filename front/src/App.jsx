import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import "./App.css";

function App() {
  const { name } = useParams();
  const [input, setInput] = useState("");
  const [text, setText] = useState("");
  const apiText = import.meta.env.VITE_API_TEXT_URL;
  const apiImage = import.meta.env.VITE_API_IMAGE_URL;
  const apiClient = import.meta.env.VITE_API_CLIENT_URL;

  const getText = () => {
    setText();
    if (name !== undefined) {
      axios.get(`${apiText}/?name=${name}`).then((res) => {
        setText(res.data);
      });
    } else {
      axios.get(apiText).then((res) => {
        setText(res.data);
      });
    }
  };

  useEffect(() => {
    getText();
  }, []);

  return (
    <div className="App">
      <div>
        <img
          src={
            name !== undefined
              ? `${apiImage}/?name=${name}`
              : apiImage
          }
          className="logo react"
          alt="image"
        />
      </div>
      <h1>{name}</h1>
      <p>{text}</p>
      <div className="card">
        <input
          value={input}
          onChange={(e) => {
            setInput(e.target.value);
          }}
        />
        <button
          onClick={() => {
            axios
              .post(`${apiClient}/${input}`)
          }}
        >
          Déploiement
        </button>
      </div>
    </div>
  );
}

export default App;
