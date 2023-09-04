export class AuthLoginModel {
  email: string | undefined
  password: string | undefined
  constructor(data?: AuthLoginModel) {
    this.email = data?.email ? data.email : undefined
    this.password = data?.password ? data.password : undefined
  }
}
