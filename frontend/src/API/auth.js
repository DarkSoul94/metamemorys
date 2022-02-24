import { $host } from "."

export const registration = (name, email, pass, repass) => {
    return $host.post("/auth/signup", {
        name: name,
        email:email,
        pass: pass,
        pass_repeat: repass
    })
}

export const login = (email, password) => {
    return $host.post("/auth/signin", {
        email:email, 
        pass: password
    })    
}