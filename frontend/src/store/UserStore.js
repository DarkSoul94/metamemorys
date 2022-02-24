import { makeAutoObservable } from "mobx"

export default class UserStore {
    constructor(){
        this._isAuth = localStorage.getItem("token") ? "true" : false

        makeAutoObservable(this)
    }

    setIsAuth(bool){
        this._isAuth = bool
    }

    get isAuth(){
        return this._isAuth
    }
}