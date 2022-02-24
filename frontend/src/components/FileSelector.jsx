import { useRef } from "react";
import Button from "react-bootstrap/Button";

const FileSelector = (props) => {
    const fileInput = useRef(null)

    const toBase64 = (file) => new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result);
        reader.onerror = error => reject(error);
    });

    const handleSelectFile = async (event) => {
        const selectedFiles = event.target.files
        let filesArray = []

        await Array.from(selectedFiles).forEach( (file) => {
            toBase64(file).then((res) => {
                filesArray.push({
                    name: file.name,
                    value: res
                })
            })
        });


        props.handleFiles(filesArray)
    };


    return (
        <>
            <Button variant="primary" onClick={() => {fileInput.current.click()}}>Select a file</Button>
            <input type="file" ref={fileInput} onChange={handleSelectFile} style={{display:"none"}} multiple accept="image/*"/>
        </>
    );
};

export default FileSelector