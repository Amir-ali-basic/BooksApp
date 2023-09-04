import type { AuthLoginDTO } from '@/Models/AuthLoginDTO'
export class LoginPresentable {
  email: string
  password: string
  constructor(data?: AuthLoginDTO) {
    //willl be model
    this.email = data?.email ? data.email : ''
    this.password = data?.password ? data.password : ''
  }
}
