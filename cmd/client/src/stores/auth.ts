import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {APIClient} from '@api/api.client';
import { Empty } from '@pkg/pbtypes/empty';
import type { SigninReq, SignupReq } from '@internal/user/dto/user';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import { getCookie, removeCookie } from 'typescript-cookie'
import { UpdateProfileReq } from '@internal/user/dto/profile';

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
      token: "" as string,
      profile: null as Profile | null,

      signupURL: new URL('signup', config.web_client_url).href,
      signinURL: new URL('signin', config.web_client_url).href,
      signinGoogleURL: new URL('signin_google', config.web_client_url).href,
      refreshTokenURL: new URL('refresh_token', config.web_client_url).href,

      api: new APIClient(new GrpcWebFetchTransport({
          baseUrl: config.api_url,
      })),
    }),
    actions: {
      async signup (req: SignupReq): Promise<Response> {
        return await fetch(this.signupURL, {
          method: 'POST',
          body: JSON.stringify(req),
        })
      },
      async signin(req:SigninReq) {
        const resp = await fetch(this.signinURL, {
          method: 'POST',
          body: JSON.stringify(req),
        })

        if (resp.status === 200) {
          this.refreshProfile();
        }

        return resp
      },
      async signinGoogle(token: string) {
        const resp = await fetch(this.signinGoogleURL, {
          method: 'POST',
          body: token,
        });

        if (resp.status === 200) {
          this.refreshProfile();
        }

        return resp
      },
      async refreshProfile() {
        this.token = getCookie('access') ?? "";

        if (!this.token) {
          return;
        }

        try {
          const resp = await this.api.fetchProfile(Empty, {meta: {token: this.token}})
          this.profile = resp.response;
        } catch (err: any) {
          switch (err.code) {
            default:
              logger.error(err);
          }
        }
      },
      async updateProfile(firstName: string|null, lastName: string|null) {
        try {
          const req = UpdateProfileReq.create({
            firstname: firstName,
            lastname: lastName,
          });

          const resp = await this.api.updateProfile(req, {meta: {token: this.token}})
          this.profile = resp.response;
        } catch (err: any) {
          switch (err.code) {
            default:
              logger.error(err);
          }
        }
      },
      async refreshToken() {
        if (!this.token) {
          return;
        }

        await fetch(this.refreshTokenURL, {
          method: 'POST',
        })
      },
      async signout() {
          this.token = "";
          this.profile = null;
          removeCookie('g_state');
          removeCookie('refresh', { path: '', domain: '.legacyfactory.com' })
          removeCookie('access', { path: '', domain: '.legacyfactory.com' })
      }
    }
});
