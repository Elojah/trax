<script setup lang="ts">
import { toRefs, ref } from "vue";
import { SignupReq } from '@internal/user/dto/user';
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import router from "@/router";

import Button from 'primevue/button';
import Divider from 'primevue/divider';
import InputText from 'primevue/inputtext';
import { Form, type FormSubmitEvent } from "@primevue/forms";
import FormField from '@primevue/forms/formfield';
import Message from 'primevue/message';

import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';

import { logger } from "@/config";

const authStore = useAuthStore()

const resolver = zodResolver(
  z.object({
    email: z.string().min(1, { message: 'Email is required.' }).email('Email must be valid.'),
    password: z.string().min(1, { message: 'Password is required.' }).min(8, { message: 'Password must be at least 8 characters long.' }),
    passwordCheck: z.string().min(1, { message: 'Password confirmation is required.' }),
    firstname: z.string().min(1, { message: 'First name is required.' }),
    lastname: z.string().min(1, { message: 'Last name is required.' }),
  }).refine((data) => data.password === data.passwordCheck, {
    message: "Passwords don't match",
    path: ["passwordCheck"],
  })
);

const initialValues = {
  email: '',
  password: '',
  passwordCheck: '',
  firstname: '',
  lastname: '',
};

const errorsStore = useErrorsStore()
const {
  message,
  success,
  error,
} = toRefs(errorsStore)

const signup = async function (e: FormSubmitEvent) {
  if (!e.valid) {
    logger.error('Signup form is invalid', e);
    return;
  }

  const req = SignupReq.create({
    email: e.values.email,
    password: e.values.password,
    firstname: e.values.firstname,
    lastname: e.values.lastname,
  });

  const resp = await authStore.signup(req);

  switch (resp.status) {
    case 200: // success
      message.value = 'Successfully signed up! You can now sign in.'
      success.value = true
      // Redirect to signin after successful signup
      setTimeout(() => {
        router.push({ name: 'signin' })
      }, 2000);
      break;
    case 409: // conflict
      message.value = 'Failed to signup, email already exists.'
      error.value = true
      break;
    case 500: // internal
      message.value = 'Internal error occurred, please retry later.'
      error.value = true
      break;
  }
}

</script>

<template>
  <Form v-slot="$form" :resolver :initialValues @submit="signup"
    class="w-full md:w-6/12 px-8 md:px-12 lg:px-24 py-16 bg-surface-0 dark:bg-surface-900 rounded-2xl md:rounded-l-none flex flex-col gap-12">
    <div class="text-surface-900 dark:text-surface-0 text-2xl font-medium leading-tight">Create Account</div>
    <div class="flex flex-col gap-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <FormField v-slot="$field" name="firstname" class="flex flex-col gap-2 flex-1">
          <label for="firstname" class="text-surface-900 dark:text-surface-0 font-medium">First Name</label>
          <InputText id="firstname" type="text" name="firstname" class="w-full rounded-md shadow-sm" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
          </Message>
        </FormField>
        <FormField v-slot="$field" name="lastname" class="flex flex-col gap-2 flex-1">
          <label for="lastname" class="text-surface-900 dark:text-surface-0 font-medium">Last Name</label>
          <InputText id="lastname" type="text" name="lastname" class="w-full rounded-md shadow-sm" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
          </Message>
        </FormField>
      </div>
      <FormField v-slot="$field" name="email" class="flex flex-col gap-2">
        <label for="email" class="text-surface-900 dark:text-surface-0 font-medium">Email Address</label>
        <InputText id="email" type="text" name="email" class="w-full rounded-md shadow-sm" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
        </Message>
      </FormField>
      <FormField v-slot="$field" name="password" class="flex flex-col gap-2">
        <label for="password" class="text-surface-900 dark:text-surface-0 font-medium">Password</label>
        <InputText id="password" type="password" name="password" class="w-full rounded-md shadow-sm" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
        </Message>
      </FormField>
      <FormField v-slot="$field" name="passwordCheck" class="flex flex-col gap-2">
        <label for="passwordCheck" class="text-surface-900 dark:text-surface-0 font-medium">Confirm Password</label>
        <InputText id="passwordCheck" type="password" name="passwordCheck" class="w-full rounded-md shadow-sm" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}
        </Message>
      </FormField>
    </div>

    <!-- Display success/error messages -->
    <div v-if="success || error" class="flex flex-col gap-2">
      <Message v-if="success" severity="success" size="small">{{ message }}</Message>
      <Message v-if="error" severity="error" size="small">{{ message }}</Message>
    </div>

    <div class="flex flex-col gap-6">
      <Button label="Create Account" class="w-full" type="submit" />
      <div class="text-center">
        <span class="text-surface-600 dark:text-surface-300">Already have an account? </span>
        <router-link to="/signin" class="text-primary font-medium cursor-pointer hover:text-primary-emphasis">Sign
          in</router-link>
      </div>
      <Divider align="center">
        <span class="text-surface-500 dark:text-surface-400">or</span>
      </Divider>
      <div class="text-center text-surface-600 dark:text-surface-300">
        <span class="text-sm">By creating an account, you agree to our Terms of Service and Privacy Policy.</span>
      </div>
    </div>
  </Form>
</template>
<style scoped></style>
