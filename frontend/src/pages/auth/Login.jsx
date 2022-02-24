import React, { useContext, useState } from "react"
import { Context } from "../.."
import { login } from "../../API/auth"
import {NavLink, useNavigate} from "react-router-dom"
import { MEMBERS_LIST_ROUTE, REG_ROUTE } from "../../utils/consts"
import {observer} from "mobx-react-lite"

import Container from "react-bootstrap/Container";
import Card from "react-bootstrap/Card";
import Form from "react-bootstrap/Form";
import Row from "react-bootstrap/Row";
import Button from "react-bootstrap/Button";

 const Login = observer( () => {
    const [email, setEmail] = useState("")
    const [pass, setPass] = useState("")
    const {user} = useContext(Context)
    const navigate = useNavigate()

     function signin() {
         login(email, pass).then((res)=>{
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
                  <h2 className="m-auto">Authorization</h2>
                  <Form className="d-flex flex-column">
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

                        <Button className="mt-3" variant={"outline-success"} onClick={signin}>Login</Button>

                        <Row className="mt-3 pl-3 pr-3">
                           <div>
                                No account? <NavLink to={REG_ROUTE}>Registration!</NavLink>
                           </div>
                        </Row>
                  </Form>
               </Card>
            </Container>
        )
});

export default Login