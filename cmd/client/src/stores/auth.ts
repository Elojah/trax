import router from '@/router';
import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {SignInError} from "@/errors"
import {APIClient} from '@api/api.client';
import { Empty } from '@pkg/pbtypes/empty';
import type { SigninReq, SignupReq } from '@internal/user/dto/user';

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
      token: null as string | null,
      profile: null as Profile | null,
      signupURL: new URL('signup', config.web_client_url).href,
      signinURL: new URL('signin', config.web_client_url).href,
    }),
    actions: {
      async signup (req: SignupReq): Promise<Response> {
        return await fetch(this.signupURL, {
          method: 'POST',
          body: JSON.stringify(req),
        })
      },
      async signin(req:SigninReq) {
        return await fetch(this.signinURL, {
          method: 'POST',
          body: JSON.stringify(req),
        })
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

        this.profile = await this.fetchProfile();
      },
      async fetchProfile(): Promise<Profile|null> {
        const client = new APIClient(config.api_url)
        const profile = await client.fetchProfile(Empty, {meta: {token: this.token}})


        return null
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
