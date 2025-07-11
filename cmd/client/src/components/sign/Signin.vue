<script setup lang="ts">
import { toRefs, ref } from "vue";
import { SigninReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import router from "@/router";
import { type CredentialResponse } from "vue3-google-signin";

import Button from 'primevue/button';
import Checkbox from 'primevue/checkbox';
import Divider from 'primevue/divider';
import InputText from 'primevue/inputtext';
import Form from '@primevue/forms/form';
import FormField from '@primevue/forms/formfield';

import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';

import { logger } from "@/config";

const checked2 = ref(true);

const authStore = useAuthStore()

const email = ref(null as string | null)
const password = ref(null as string | null)
const valid = ref(null as boolean | null)
const showPassword = ref(false as boolean)

const resolver = zodResolver(
  z.object({
    email: z.string().min(1, { message: 'Email is required.' }).email('Email must be valid.'),
    password: z.string().min(1, { message: 'Password is required.' }).min(8, { message: 'Password must be at least 8 characters long.' }),
  })
);

const errorsStore = useErrorsStore()
const {
  message,
  error,
} = toRefs(errorsStore)

const signin = async function () {
  console.log('signin', email.value, password.value)

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
  <div
    class="w-full md:w-6/12 px-8 md:px-12 lg:px-24 py-16 bg-surface-0 dark:bg-surface-900 rounded-2xl md:rounded-l-none flex flex-col gap-12">
    <Form v-slot="$form" :resolver @submit="signin" class="flex flex-col gap-6">
      <div class="text-surface-900 dark:text-surface-0 text-2xl font-medium leading-tight">Login</div>
      <FormField class="flex flex-col gap-2">
        <label for="email3" class="text-surface-900 dark:text-surface-0 font-medium">Email
          Address</label>
        <InputText v-model="email" id="email3" type="text" class="w-full rounded-md shadow-sm" />
      </FormField>
      <FormField class="flex flex-col gap-2">
        <label for="password3" class="text-surface-900 dark:text-surface-0 font-medium">Password</label>
        <InputText v-model="password" id="password3" type="password" class="w-full rounded-md shadow-sm" />
      </FormField>
      <div
        class="flex flex-col sm:flex-row md:flex-col lg:flex-row items-start sm:items-center md:items-start lg:items-center justify-between w-full gap-2">
        <div class="flex items-center gap-2">
          <Checkbox id="rememberme2" v-model="checked2" :binary="true" />
          <label for="rememberme2" class="text-surface-900 dark:text-surface-0">Remember me</label>
        </div>
        <a
          class="text-surface-500 dark:text-surface-400 font-medium cursor-pointer hover:text-surface-600 dark:hover:text-surface-300">Forgot
          your password?</a>
      </div>
    </Form>
    <div class="flex flex-col gap-6">
      <Button label="Login" class="w-full" />
      <Divider align="center">
        <span class="text-surface-500 dark:text-surface-400">or</span>
      </Divider>
      <Button label="Sign In with Google" icon="pi pi-google !text-base !leading-none" severity="secondary"
        class="w-full" />
    </div>
    <div class="text-center">
      <span class="text-surface-600 dark:text-surface-300">Don't have an account? </span>
      <span class="text-primary font-medium cursor-pointer hover:text-primary-emphasis">Sign up</span>
    </div>
  </div>
</template>
<style scoped></style>
