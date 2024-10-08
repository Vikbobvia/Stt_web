import React, { useState } from "react";
import { Upload, Input, Form, Button } from "antd";
import axios from "axios";

const FileInput = () => {
  const [creatorName, setCreatorName] = useState("");
  const [fileList, setFileList] = useState([]);
  const [wordCount, setWordCount] = useState(null); // Store word count result
  const [fileName, setFileName] = useState(""); // Store file name
  const [status, setStatus] = useState(""); // Store status message
  const [loading, setLoading] = useState(false); // Store loading state

  const onUploadChange = ({ fileList: newFileList }) => {
    setFileList(newFileList);
    if (newFileList.length > 0) {
      setFileName(newFileList[0].name); // Set the file name when a file is selected
    } else {
      setFileName(""); // Clear file name if no file is selected
    }
  };

  const handleSubmit = () => {
    // Check if a file is selected
    console.log("File List:", fileList);

    if (fileList.length === 0) {
      alert("Please upload a file.");
      return;
    }

    const fileReader = new FileReader();
    setLoading(true); // Start loading state

    // Reading the file content
    fileReader.onload = (event) => {
      console.log("File has been loaded."); // Check if this logs
      const fileContent = event.target.result; // File content as string

      const payload = {
        creator_name: creatorName,
        file_name: fileName, // Include file name in payload
        file_content: fileContent, // Sending file content to Go server
      };

      console.log("Payload being sent to server:", payload); // Debugging payload

      // Axios POST request to Go server
      axios
        .post("http://localhost:8080/api/countWords", payload)
        .then((response) => {
          console.log("Response from server:", response.data); // Debugging response
          // Update state with the response from the Go server
          setWordCount(response.data.word_count); // Handle the word count from Go server
          setStatus(response.data.status); // Handle the status message from Go server
        })
        .catch((error) => {
          console.error("Error occurred during request:", error); // Debugging error
          setStatus("Error occurred while processing the file."); // Update status on error
        })
        .finally(() => {
          setLoading(false); // Stop loading state
        });
    };

    fileReader.readAsText(fileList[0].originFileObj); // Read the content of the selected file
  };

  return (
    <Form layout="vertical">
      <Form.Item label="Creator Name">
        <Input
          value={creatorName}
          onChange={(e) => setCreatorName(e.target.value)}
        />
      </Form.Item>
      <Form.Item label="Sound File">
        <Upload
          name="soundFile"
          listType="picture"
          fileList={fileList}
          onChange={onUploadChange}
          beforeUpload={() => false} // Prevent automatic upload
        >
          <Button type="button">Click to upload</Button>
        </Upload>
        {fileList.length > 0 && (
          <p>
            Selected File: {fileName} (and {fileList.length - 1} more)
          </p>
        )}
      </Form.Item>
      <Form.Item>
        <Button type="primary" onClick={handleSubmit} loading={loading}>
          Submit
        </Button>
      </Form.Item>
      {status && (
        <div>
          <h3>Status: {status}</h3>
        </div>
      )}
      {wordCount !== null && (
        <div>
          <h3>Word Count Result:</h3>
          <p>{wordCount} words found</p>
        </div>
      )}
    </Form>
  );
};

export default FileInput;
