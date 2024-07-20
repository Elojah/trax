<script setup lang="ts">
import { toRefs, ref } from "vue";
import { SigninReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import type { VForm } from "vuetify/components/VForm";
import { type CredentialResponse } from "vue3-google-signin";
import router from "@/router";

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


const errorsStore = useErrorsStore()
const {
  message,
  success,
  error,
} = toRefs(errorsStore)

const signin = async function () {
  const req = SigninReq.create({
    email: email.value,
    password: password.value,
  });

  const resp = await authStore.signin(req);

  switch (resp.status) {
    case 200: // success
      router.push({ name: 'dashboard' })
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

const signInGoogle = async function (credentials: CredentialResponse) {
  const resp = await authStore.signinGoogle(credentials.credential!);

  switch (resp.status) {
    case 200: // success
      router.push({ name: 'dashboard' })
      break;
    case 401: // unauthorized
      message.value = 'Failed to signin, wrong email/password.'
      error.value = true
      form?.value?.reset()
      break;
    case 500: // internal
      message.value = 'Internal error occurred, please retry later.'
      error.value = true
      form?.value?.reset()
      break;
  }
}

const signInGoogleError = async function (error: any) {
  console.log(error)
  snackbarInternal.value = true
  form?.value?.reset()
}
</script>

<template>
  <v-card class="px-6 py-6 rounded" variant="elevated">
    <v-card-item class="justify-center mb-2" prepend-icon="mdi-login">
      <v-card-title>
        SIGNIN
      </v-card-title>
      <template v-slot:prepend>
        <v-icon color="primary"></v-icon>
      </template>
    </v-card-item>
    <v-divider color="primary"></v-divider>
    <v-card-text class="mt-6">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="12">
            <v-text-field v-model="email" label="Email" :rules="emailRules" prepend-inner-icon="mdi-email" underlined
              required clearable></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="password" prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :rules="passwordRules"
              :type="showPassword ? 'text' : 'password'" label="Password" hint="At least 8 characters" counter
              @click:append-inner="showPassword = !showPassword" underlined required clearable>
            </v-text-field>
          </v-col>
          <v-col class="mb-6" cols="12" align="center">
            <v-btn size="large" :disabled="!valid" variant="tonal" append-icon="mdi-account-circle"
              @click="signin">Signin
              <template v-slot:append>
                <v-icon color="primary"></v-icon>
              </template>
            </v-btn>
          </v-col>
          <v-divider color="primary"></v-divider>
          <v-col class="mt-6 color-auto" cols="12" align="center">
            <GoogleSignInButton @success="signInGoogle" @error="signInGoogleError" theme="filled_black" size="large"
              shape="pill">
            </GoogleSignInButton>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>
<style scoped>
.color-auto {
  color-scheme: auto;
}
</style>
