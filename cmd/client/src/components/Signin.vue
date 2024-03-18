<script setup lang="ts">
import { ref } from "vue";
import { useAuthStore } from '@/stores/auth';
import { GoogleLogin } from 'vue3-google-login';
import { config } from '@/config';

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
const signIn = function () {
  // Add your sign in logic here
  authStore.signinGoogle(email.value!);
  console.log('Signing in with ', email.value, password.value);
}
const signInGoogle = function () {
  // Add your sign in logic here
  authStore.signinGoogle(email.value!);
  console.log('Signing in with ', email.value, password.value);
}

</script>

<template>
  <v-card class="px-6 py-6" rounded>
    <v-card-item class="justify-center mb-2" prepend-icon="mdi-login">
      <v-card-title>
        SIGNIN
      </v-card-title>
      <template v-slot:prepend>
        <v-icon color="success"></v-icon>
      </template>
    </v-card-item>
    <v-divider></v-divider>
    <v-card-text class="mt-6">
      <v-form ref="signin-form" v-model="valid" lazy-validation>
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
              @click="signIn">Signin
              <template v-slot:append>
                <v-icon color="success"></v-icon>
              </template>
            </v-btn>
          </v-col>
          <v-divider></v-divider>
          <v-col class="mt-6" cols="12" align="center">
            <GoogleLogin :client-id="googleClientID" :callback="signInGoogle">
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
    </v-card-text>
  </v-card>
</template>
<style scoped></style>
