import Navbar from "react-bootstrap/Navbar"
import Container from "react-bootstrap/Container"
import Button from "react-bootstrap/Button"
import { LOGIN_ROUTE, MEMBERS_LIST_ROUTE } from "../utils/consts";
import { useNavigate } from "react-router-dom";
import { useContext } from "react";
import { Context } from "..";
import { observer } from "mobx-react-lite";

const NavBar = observer(() => {
    const navigator = useNavigate();

    const {user} = useContext(Context)
    
    function logout(){
        localStorage.removeItem("token")
        user.setIsAuth(false)
        navigator(LOGIN_ROUTE)
    };

    return (
        <Navbar bg="dark">
            <Container>
                <Navbar.Brand href={MEMBERS_LIST_ROUTE} style={{color:"white"}}>Metamemorys</Navbar.Brand>
                <Navbar.Collapse className="justify-content-end">
                    {user.isAuth ? <Button variant="outline-light" onClick={logout}>Logout</Button> : <div></div>}
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
});

export default NavBar