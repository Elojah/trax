import { defineStore } from 'pinia'
import { U } from '@internal/user/user'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import type { SigninReq, SignupReq } from '@internal/user/dto/user'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { removeCookie } from 'typescript-cookie'
import { FetchUserReq, UpdateUserReq } from '@internal/user/dto/user'
import { ref } from 'vue'
import { ClaimAuth } from '@internal/user/claims'
import { Empty } from '@pkg/pbtypes/empty'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string>('')
  const user = ref<U | null>(null)
  const claims = ref<ClaimAuth | null>(null)

  const signupURL = new URL('signup', config.web_client_url).href
  const signinURL = new URL('signin', config.web_client_url).href
  const signoutURL = new URL('signout', config.web_client_url).href
  const signinGoogleURL = new URL('signin_google', config.web_client_url).href
  const refreshTokenURL = new URL('refresh_token', config.web_client_url).href

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )

  const signup = async (req: SignupReq): Promise<Response> => {
    return await fetch(signupURL, {
      method: 'POST',
      body: JSON.stringify(req)
    })
  }

  const signin = async (req: SigninReq) => {
    const resp = await fetch(signinURL, {
      method: 'POST',
      body: JSON.stringify(req)
    })

    if (resp.status === 200) {
      token.value = await resp.text()
      refreshUser()
    }

    return resp
  }

  const signinGoogle = async (body: string) => {
    const resp = await fetch(signinGoogleURL, {
      method: 'POST',
      body: body
    })

    if (resp.status === 200) {
      token.value = await resp.text()
      refreshUser()
    }

    return resp
  }

  const refreshUser = async () => {
    if (!token.value) {
      refreshToken()
    }

    try {
      const req = FetchUserReq.create({
        me: true
      })

      const resp = await api.fetchUser(req, { meta: { token: token.value } })
      user.value = resp.response

      await refreshClaims()
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const refreshClaims = async () => {
    if (!token.value) {
      refreshToken()
    }

    try {
      const req = Empty.create({})

      const resp = await api.fetchClaims(req, { meta: { token: token.value } })

      claims.value = resp.response
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const updateUser = async (firstName: string | undefined, lastName: string | undefined, avatarURL: string | undefined = undefined) => {
    try {
      const req = UpdateUserReq.create({
        ...(firstName && { firstname: { value: firstName } }),
        ...(lastName && { lastname: { value: lastName } }),
        ...(avatarURL !== undefined && { avatarURL: { value: avatarURL } })
      })

      const resp = await api.updateUser(req, { meta: { token: token.value } })
      user.value = resp.response
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const refreshToken = async () => {
    const resp = await fetch(refreshTokenURL, {
      method: 'POST'
    })

    if (resp.status === 200) {
      token.value = await resp.text()
      await refreshUser()
    }

    return resp
  }

  const signout = async () => {
    const resp = await fetch(signoutURL, {
      method: 'GET'
    })

    if (resp.status !== 200) {
      logger.error('failed to sign out')
    }

    token.value = ''
    user.value = null

    removeCookie('g_state')
    removeCookie('g_csrf_token')
  }

  return {
    token,
    user,
    claims,
    signup,
    signin,
    signout,
    signinGoogle,
    refreshUser,
    updateUser,
    refreshToken
  }
})
