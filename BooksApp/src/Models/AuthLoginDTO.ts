export class AuthLoginDTO {
  email: string | null
  password: string | null
  constructor(data?: AuthLoginDTO) {
    this.email = data?.email ? data.email : null
    this.password = data?.password ? data.password : null
  }
}
