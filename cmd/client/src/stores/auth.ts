import { defineStore } from 'pinia'
import { Profile } from '@internal/user/user'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import type { SigninReq, SignupReq } from '@internal/user/dto/user'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { getCookie, removeCookie } from 'typescript-cookie'
import { FetchProfileReq, UpdateProfileReq } from '@internal/user/dto/profile'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const token = ref('' as string)
  const profile = ref(null as Profile | null)

  const signupURL = new URL('signup', config.web_client_url).href
  const signinURL = new URL('signin', config.web_client_url).href
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
      refreshProfile()
    }

    return resp
  }

  const signinGoogle = async (token: string) => {
    const resp = await fetch(signinGoogleURL, {
      method: 'POST',
      body: token
    })

    if (resp.status === 200) {
      refreshProfile()
    }

    return resp
  }

  const refreshProfile = async () => {
    token.value = getCookie('access') ?? ''

    if (!token.value) {
      return
    }

    try {
      const req = FetchProfileReq.create({
        me: true
      })

      const resp = await api.fetchProfile(req, { meta: { token: token.value } })
      profile.value = resp.response
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const updateProfile = async () => {
    try {
      const req = UpdateProfileReq.create({
        firstname: { value: profile?.value?.firstName },
        lastname: { value: profile?.value?.lastName }
      })

      const resp = await api.updateProfile(req, { meta: { token: token.value } })
      profile.value = resp.response
    } catch (err: any) {
      switch (err.code) {
        default:
          logger.error(err)
      }
    }
  }

  const refreshToken = async () => {
    if (!token.value) {
      return
    }

    await fetch(refreshTokenURL, {
      method: 'POST'
    })
  }

  const signout = async () => {
    token.value = ''
    profile.value = null
    removeCookie('g_state')
    removeCookie('refresh', { path: '', domain: '.legacyfactory.com' })
    removeCookie('access', { path: '', domain: '.legacyfactory.com' })
  }

  return {
    token,
    profile,
    signup,
    signin,
    signout,
    signinGoogle,
    refreshProfile,
    updateProfile,
    refreshToken
  }
})
