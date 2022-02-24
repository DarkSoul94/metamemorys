import { $authHost } from "."

export const getMemberList = () => {
    return $authHost.get("/member/list")
}

export const createMember = (name) => {
    return $authHost.post("/member/create", {name:name});
}