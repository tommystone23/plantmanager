import React, { useState } from "react";

function App() {
  const [file, setFile] = useState(null);
  const [image, setImage] = useState("")
  const [message, setMessage] = useState("");

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) {
      setMessage("Please choose a file first.");
      return;
    }

    const formData = new FormData();
    formData.append("image", file);

    try {
      const res = await fetch("/api/upload/image", {
        method: "POST",
        body: formData,
      });

      if(res.ok)
        setMessage("File Uploaded Successfully")

      const data = await res.json()
      setImage(data.image)

    } catch (err) {
      console.error(err);
      setMessage("Error uploading.");
    }
  };

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Plant Manager</h1>
      <h2>Upload a Photo</h2>
      <input type="file" accept="image/*" onChange={handleFileChange} />
      <button onClick={handleUpload}>Upload</button>
      <p>{message}</p>

      {image && (
        <div style={{ marginTop: "1rem" }}>
          <h3>Preview:</h3>
          <img src={image} alt="Preview" style={{ maxWidth: "300px" }} />
        </div>
      )}
    </div>
  );
}

export default App;
