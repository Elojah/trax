<script setup lang="ts">
import { ref } from "vue";
import { SigninReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { GoogleLogin } from 'vue3-google-login';
import { config } from '@/config';
import type { VForm } from "vuetify/components/VForm";

const form = ref<VForm | null>(null);

const authStore = useAuthStore()

const email = ref(null as string | null)
const password = ref(null as string | null)
const valid = ref(null as boolean | null)
const showPassword = ref(false as boolean)
const emailRules = [
  (v: string) => !!v || "Required",
  (v: string) => /.+@.+\..+/.test(v) || "Email must be valid"
]

const passwordRules = [
  (v: string) => !!v || "Required.",
  (v: string) => (v && v.length >= 8) || "Min 8 characters"
]

const googleClientID = config.google_client_id

const snackbarUnauthorized = ref(false)
const snackbarInternal = ref(false)

const signin = async function () {
  const req = SigninReq.create({
    email: email.value,
    password: password.value,
  });

  const resp = await authStore.signin(req);

  switch (resp.status) {
    case 200: // success
      // TODO: redirect
      break;
    case 401: // unauthorized
      snackbarUnauthorized.value = true
      form?.value?.reset()
      break;
    case 500: // internal
      snackbarInternal.value = true
      form?.value?.reset()
      break;
  }
}
const signInGoogle = async function (token: string) {
  const resp = await authStore.signinGoogle(token);

  switch (resp.status) {
    case 200: // success
      authStore.refreshProfile()
      // TODO: redirect
      break;
    case 401: // unauthorized
      snackbarUnauthorized.value = true
      form?.value?.reset()
      break;
    case 500: // internal
      snackbarInternal.value = true
      form?.value?.reset()
      break;
  }
}

authStore.refreshProfile()
</script>

<template>
  <v-card class="px-6 py-6 rounded-xl" variant="elevated">
    <v-card-item class="justify-center mb-2" prepend-icon="mdi-login">
      <v-card-title>
        SIGNIN
      </v-card-title>
      <template v-slot:prepend>
        <v-icon color="success"></v-icon>
      </template>
    </v-card-item>
    <v-divider color="success"></v-divider>
    <v-card-text class="mt-6">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="12">
            <v-text-field v-model="email" label="Email" :rules="emailRules" append-icon="mdi-email" underlined required
              clearable></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="password" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="passwordRules" :type="showPassword ? 'text' : 'password'" label="Password"
              hint="At least 8 characters" counter @click:append="showPassword = !showPassword" underlined required
              clearable>
            </v-text-field>
          </v-col>
          <v-col class="mb-6" cols="12" align="center">
            <v-btn size="large" :disabled="!valid" variant="tonal" append-icon="mdi-account-circle"
              @click="signin">Signin
              <template v-slot:append>
                <v-icon color="success"></v-icon>
              </template>
            </v-btn>
          </v-col>
          <v-divider color="success"></v-divider>
          <v-col class="mt-6" cols="12" align="center">
            <GoogleLogin :client-id="googleClientID" :callback="signInGoogle" prompt auto-login>
              <v-btn size="large" variant="tonal" append-icon="mdi-google">
                Signin with Google
                <template v-slot:append>
                  <v-icon color="success"></v-icon>
                </template>
              </v-btn>
            </GoogleLogin>
          </v-col>
        </v-row>
      </v-form>
      <v-snackbar :timeout="2000" v-model="snackbarUnauthorized" class="mt-6" color="error">
        Failed to signin, wrong email/password.
        <template v-slot:actions>
          <v-btn variant="text" @click="snackbarUnauthorized = false">Close</v-btn>
        </template>
      </v-snackbar>
      <v-snackbar :timeout="2000" v-model="snackbarInternal" class="mt-6" color="error">
        Internal error occurred, please retry later.
        <template v-slot:actions>
          <v-btn variant="text" @click="snackbarInternal = false">Close</v-btn>
        </template>
      </v-snackbar>
    </v-card-text>
  </v-card>
</template>
<style scoped></style>
