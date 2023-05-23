export class LoginPresentable {
  email: string
  password: string
  constructor(data?: any) {
    //willl be model
    this.email = data?.email ? data.email : ''
    this.password = data?.password ? data.password : ''
  }
}
