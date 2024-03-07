import router from '@/router';
import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {SignInError} from "@/errors"
import {APIClient} from '@api/api.client';
import { Empty } from '@pkg/pbtypes/empty';

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
      token: null as string | null,
      profile: null as Profile | null,
    }),
    actions: {
      async signin(email:string, password:string) {
        // const signin = await fetch(
        //   new URL('signin', config.web_client_url).href, {
        //   method: 'POST',
        //   headers: {
        //     'Content-Type': 'application/json'
        //   },
        //   body: JSON.stringify({email, password})
        // });
        // this.user = new U();
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