import { observer } from "mobx-react-lite";
import { useEffect, useState } from "react";
import { getMemberList } from "../../API/members";
import ListGroup from "react-bootstrap/ListGroup"
import Container from "react-bootstrap/Container"
import NewMember from "../../components/NewMember";
import Button from "react-bootstrap/Button";
import Member from "../../components/Member";
import { getMemberFiles } from "../../API/files";

const MembersList = observer(() => {
    const [showNewMemberModal, setShowMewMemberModal] = useState(false)
    const [showMemberModal, setShowMemberModal] = useState(false)
    const [selectedMember, setSelectedMember] = useState([])
    const [memberFiles, setMemberFiles] = useState([])

    const[members, setMembers] = useState([])

    useEffect(() => {
        getMemberList().then((res) => {
            setMembers(res.data.members)
        })
    }, []);

    function showMember(member) {
        setSelectedMember(member)

        getMemberFiles(member.id).then((res) => {
            console.log(res.data)
            setMemberFiles(res.data.files !== null ? res.data.files : [])
        })

        setShowMemberModal(true)
    }

    return (
        <Container style={{width: 800}} className="mt-3">
            <NewMember show={showNewMemberModal} onHide={() => setShowMewMemberModal(false)} />
            <Button variant={"outline-success"} onClick={() => setShowMewMemberModal(true)}>New Member</Button>

            <Member show={showMemberModal} onHide={() => setShowMemberModal(false)} member={selectedMember} files={memberFiles}/>

            <ListGroup className="mt-3">
                {members.map((member) => 
                    <ListGroup.Item key={member.id} action onClick={() => showMember({id: member.id, name: member.name})} >{member.name}</ListGroup.Item>
                )}
            </ListGroup>
        </Container>
    );
});

export default MembersList