import Modal from "react-bootstrap/Modal"
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react"
import { createMember } from "../API/members"
import FileSelector from "./FileSelector"
import { sendFile } from "../API/files"

const NewMember = (props) => {
    const [name, setName] = useState("")
    const [selectedFiles, setSelectedFiles] = useState([]);

    function create() {
        createMember(name).then((res) => {
            sendFile({
                member_id: res.data.id,
                files: selectedFiles
            })
        });

        props.onHide();
    }

    return (
        <Modal
            {...props}
            size="lg"
            centered
        >
            <Modal.Header>
                <Modal.Title>New Member</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form>
                    <Form.Group>
                        <Form.Label>Member name</Form.Label>
                        <Form.Control type="text" placeholder="Enter member name" value={name} onChange={e => setName(e.target.value)}/>
                    </Form.Group>
                    <Form.Group className="mt-3">
                        <FileSelector handleFiles={(files) => {setSelectedFiles(files)}} />
                    </Form.Group>
                </Form>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={props.onHide}>Close</Button>
                <Button variant="primary" onClick={create}>Create</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default NewMember