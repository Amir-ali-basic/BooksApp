import type { TokenData } from './TokenData'

export interface LoginResponse {
  success: boolean
  data: {
    token: TokenData
  }
}
