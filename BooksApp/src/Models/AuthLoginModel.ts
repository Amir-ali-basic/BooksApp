export class AuthLoginModel {
  email: string | null
  password: string | null
  constructor(data?: AuthLoginModel) {
    this.email = data?.email ? data.email : null
    this.password = data?.password ? data.password : null
  }
}
