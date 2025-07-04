<script setup lang="ts">
import { toRefs, ref } from "vue";
import { SigninReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { type CredentialResponse } from "vue3-google-signin";
import router from "@/router";
import { logger } from "@/config";

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
      message.value = 'Failed to signin, wrong email/password.'
      error.value = true;
      // form?.value?.reset()
      break;
    case 500: // internal
      message.value = 'Internal error occurred, please retry later.'
      error.value = true;
      // form?.value?.reset()
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
      // form?.value?.reset()
      break;
    case 500: // internal
      message.value = 'Internal error occurred, please retry later.'
      error.value = true
      // form?.value?.reset()
      break;
  }
}

const signInGoogleError = async function (error: any) {
  message.value = `Google returned an error: ${error}`
  error.value = true
  // form?.value?.reset()
}

</script>

<template>
</template>
<style scoped>
</style>
