import { observer } from "mobx-react-lite";
import { useContext } from "react";
import {Navigate, Route, Routes} from "react-router-dom"
import { Context } from ".";
import Login from "./pages/auth/Login";
import Registration from "./pages/auth/Registration";
import MembersList from "./pages/members/MembersList";
import { LOGIN_ROUTE, MEMBERS_LIST_ROUTE, REG_ROUTE } from "./utils/consts";

const AppRouter = observer(() => {
    function RequireAuth({ children }) {
        const {user} = useContext(Context)
        let isAuthenticated = user.isAuth;
        return isAuthenticated ? children : <Navigate to={LOGIN_ROUTE} />;
    }

    return (
        <Routes>
            <Route path={REG_ROUTE} element={<Registration/>}/>
            <Route path={LOGIN_ROUTE} element={<Login/>}/>

            <Route path={MEMBERS_LIST_ROUTE} element={<RequireAuth><MembersList /></RequireAuth>} />

            <Route path={"*"} element={<Navigate to={LOGIN_ROUTE}/>}/>
        </Routes>
    );
});

export default AppRouter

