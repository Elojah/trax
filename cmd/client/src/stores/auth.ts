import router from '@/router';
import { defineStore } from 'pinia';
import {Profile} from '@internal/user/user';
import {config, logger} from "@/config"
import {SignInError} from "@/errors"
import {APIClient} from '@api/api.client';
import * as grpc from 'grpc-web';

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
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
        const client = new APIClient('https://api.legacyfactory.com:8082', null)
        const metadata: grpc.Metadata = {}//{ 'token': getCookie('access')! }

        // const refreshRooms = () => {
        //   const req = new ListRoomReq()
        //   req.setSize(100)
        //   setRooms({ rooms: [], loaded: false })
        //   client.listRoomPublic(req, metadata).then((result) => {
        //     console.log('found ', result.getRoomsList().length, ' rooms')

        //     setRooms({ rooms: result.getRoomsList(), loaded: true })
        //   }).catch((err) => {
        //     console.log(err)
        //   })
        // }


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
