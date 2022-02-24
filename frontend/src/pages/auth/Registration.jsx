import { useContext, useState } from "react"
import {NavLink, useNavigate} from "react-router-dom"
import { Context } from "../.."
import { registration } from "../../API/auth"
import { LOGIN_ROUTE, MEMBERS_LIST_ROUTE } from "../../utils/consts"

import Container from "react-bootstrap/Container";
import Card from "react-bootstrap/Card";
import Form from "react-bootstrap/Form";
import Row from "react-bootstrap/Row";
import Button from "react-bootstrap/Button";


const Registration = () => {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [pass, setPass] = useState("")
    const [repass, setRepass] = useState("")
    const {user} = useContext(Context)
    const navigate = useNavigate()

    function signup() {
        registration(name, email, pass, repass).then((res) => {
            user.setIsAuth(true)
            localStorage.setItem("token", "Bearer " + res.data.token)
            navigate(MEMBERS_LIST_ROUTE)
        })
    }

    return (
        <Container 
        className="d-flex justify-content-center align-items-center"
        style={{height: window.innerHeight - 54}}
        >
            <Card style={{width: 600}} className="p-5">
                <h2 className="m-auto">Registration</h2>
                <Form className="d-flex flex-column">
                    <Form.Control
                        type="text"
                        className="mt-3" 
                        placeholder="Enter name"
                        value={name}
                        onChange = {e => setName(e.target.value)}
                    />

                    <Form.Control
                        type="email"
                        className="mt-3" 
                        placeholder="Enter email"
                        value={email}
                        onChange = {e => setEmail(e.target.value)}
                    />

                    <Form.Control
                        type="password" 
                        className="mt-3" 
                        placeholder="Enter password"
                        value={pass}
                        onChange = {e => setPass(e.target.value)}
                    />

                    <Form.Control
                        type="password" 
                        className="mt-3" 
                        placeholder="Repeat password"
                        value={repass}
                        onChange = {e => setRepass(e.target.value)}
                    />

                    <Button className="mt-3" variant={"outline-success"} onClick={signup}>Registration</Button>

                    <Row className="mt-3 pl-3 pr-3">
                        <div>
                            Have account? <NavLink to={LOGIN_ROUTE}>Authorization!</NavLink>
                        </div>
                    </Row>
                </Form>
            </Card>
        </Container>
    )
}

export default Registration