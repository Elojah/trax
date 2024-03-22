import router from '@/router';
import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {SignInError} from "@/errors"
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
          this.token = getCookie('token') ?? "";
        }

        return resp
      },
      async signinGoogle(token: string) {
        logger.info('signin google attempt');

        const signin = await fetch(
          new URL('signin_google', config.web_client_url).href, {
          method: 'POST',
          headers: {},
          body: token
        });

        if (signin.status !== 200) {
          const err = `signin google failed with status ${signin.status}`
          logger.error(err);

          throw new SignInError(err);
        }
      },
      async refreshProfile() {
        await this.api.fetchProfile(Empty, {meta: {token: this.token}})
      },
      async signinTwitch(token:string) {},
      async refresh() {
      },
        async signout() {
            this.profile = null;
            router.push('/signin');
      }
    }
});
