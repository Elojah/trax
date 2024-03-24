import router from '@/router';
import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {APIClient} from '@api/api.client';
import { Empty } from '@pkg/pbtypes/empty';
import type { SigninReq, SignupReq } from '@internal/user/dto/user';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import { getCookie } from 'typescript-cookie'

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
      token: "" as string,
      profile: null as Profile | null,

      signupURL: new URL('signup', config.web_client_url).href,
      signinURL: new URL('signin', config.web_client_url).href,
      signinGoogleURL: new URL('signin_google', config.web_client_url).href,

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
      async refresh() {
      },
        async signout() {
            this.profile = null;
            router.push('/signin');
      }
    }
});
