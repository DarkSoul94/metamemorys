import { $authHost } from ".";

export const sendFile = data => {
   return $authHost.post("/file/create", data)
};

export const getMemberFiles = (id) => {
   return $authHost.get("/file/list", {params: {memberID: id}})
}