<script setup lang="ts">
import { toRefs, ref } from "vue";
import { SigninReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import router from "@/router";
import { type CredentialResponse } from "vue3-google-signin";

import { Form } from '@primevue/forms';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';

import { logger } from "@/config";

const authStore = useAuthStore()

const email = ref(null as string | null)
const password = ref(null as string | null)
const valid = ref(null as boolean | null)
const showPassword = ref(false as boolean)

const resolver =  zodResolver(
    z.object({
        username: z.string().min(1, { message: 'Username is required.' }).email('Username must be a valid email.'),
        password: z.string().min(1, { message: 'Password is required.' }).min(8, { message: 'Password must be at least 8 characters long.' }),
    })
);

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
<template>
    <div class="card flex justify-center">
        <Toast />
        <Form v-slot="$form" :resolver @submit="signin" class="flex flex-col gap-4 w-full sm:w-60">
            <div class="flex flex-col gap-1">
                <InputText name="username" type="text" placeholder="Username" fluid />
                <Message v-if="$form.username?.invalid" severity="error" size="small" variant="simple">{{ $form.username.error.message }}</Message>
            </div>
            <div class="flex flex-col gap-1">
                <Password name="password" placeholder="Password" :feedback="false" toggleMask fluid />
                <Message v-if="$form.password?.invalid" severity="error" size="small" variant="simple">
                    <ul class="my-0 px-4 flex flex-col gap-1">
                        <li v-for="(error, index) of $form.password.errors" :key="index">{{ error.message }}</li>
                    </ul>
                </Message>
            </div>
            <Button type="submit" severity="secondary" label="Submit" />
        </Form>
    </div>
</template>
</template>
<style scoped>
</style>
