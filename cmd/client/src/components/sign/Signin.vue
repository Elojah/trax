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
import { Form, type FormSubmitEvent } from "@primevue/forms";
import FormField from '@primevue/forms/formfield';
import Message from 'primevue/message';

import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';

import { logger } from "@/config";

const checked2 = ref(true);

const authStore = useAuthStore()

const resolver = zodResolver(
  z.object({
    email: z.string().min(1, { message: 'Email is required.' }).email('Email must be valid.'),
    password: z.string().min(1, { message: 'Password is required.' }).min(8, { message: 'Password must be at least 8 characters long.' }),
  })
);

const initialValues = {
  email: '',
  password: '',
};

const errorsStore = useErrorsStore()
const {
  message,
  error,
} = toRefs(errorsStore)

const signin = async function (e: FormSubmitEvent) {
  if (!e.valid) {
    logger.error('Signin form is invalid', e);
    return;
  }

  const req = SigninReq.create({
    email: e.values.email,
    password: e.values.password,
  });

  const resp = await authStore.signin(req);

  switch (resp.status) {
    case 200: // success
      router.push({ name: 'home' })
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
      router.push({ name: 'home' })
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
  <Form v-slot="$form" :resolver :initialValues @submit="signin"
    class="w-full md:w-6/12 px-8 md:px-12 lg:px-24 py-16 bg-surface-0 dark:bg-surface-900 rounded-2xl md:rounded-l-none flex flex-col gap-12">
    <div class="text-surface-900 dark:text-surface-0 text-2xl font-medium leading-tight">Login</div>
    <div class="flex flex-col gap-6">
      <FormField v-slot="$field" name="email" class="flex flex-col gap-2">
        <label for="email3" class="text-surface-900 dark:text-surface-0 font-medium">Email Address</label>
        <InputText id="email3" type="text" name="email" class=" w-full rounded-md shadow-sm" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
        </Message>
      </FormField>
      <FormField v-slot="$field" name="password" class="flex flex-col gap-2">
        <label for="password3" class="text-surface-900 dark:text-surface-0 font-medium">Password</label>
        <InputText id="password3" type="password" name="password" class="w-full rounded-md shadow-sm" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
        </Message>
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
    </div>
    <div class="flex flex-col gap-6">
      <Button label="Login" class="w-full" type="submit" />
      <div class="text-center">
        <span class="text-surface-600 dark:text-surface-300">Don't have an account? </span>
        <span class="text-primary font-medium cursor-pointer hover:text-primary-emphasis">Sign up</span>
      </div>
      <Divider align="center">
        <span class="text-surface-500 dark:text-surface-400">or</span>
      </Divider>
      <div class="flex flex-col items-center w-full" style="color-scheme: auto;">
        <GoogleSignInButton @success="signInGoogle" @error="signInGoogleError" theme="filled_black" auto-select="true"
          size="large">
        </GoogleSignInButton>
      </div>
    </div>
  </Form>
</template>
<style scoped></style>
