import Modal from "react-bootstrap/Modal"
import Button from "react-bootstrap/Button"
import ListGroup from "react-bootstrap/ListGroup"

const Member = (props) => {

    return (
        <Modal
            {...props}
            size="lg"
            centered
        >
            <Modal.Header>
                <Modal.Title>{props.member.name}</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                    {props.files.map((file) => {
                        console.log(file.value)
                        return (<img src={file.value} key={file.id} alt={file.name}/>)
                    }
                    )}
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={props.onHide}>Close</Button>
                <Button variant="primary">Create</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default Member