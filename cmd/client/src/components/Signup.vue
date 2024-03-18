<script setup lang="ts">
import { config } from '@/config';
import { useAuthStore } from '@/stores/auth';
import { SignupReq } from '@internal/user/dto/user';
import { ref } from 'vue';


// const authStore = useAuthStore();

const email = ref(null as string | null)
const password = ref(null as string | null)
const firstname = ref(null as string | null)
const lastname = ref(null as string | null)
const valid = ref(null as boolean | null)
const passwordCheck = ref(null as string | null)
const showPassword = ref(false)
const showPasswordCheck = ref(false)
const emailRules = [
  (v: string) => !!v || "Required",
  (v: string) => /.+@.+\..+/.test(v) || "Email must be valid"
]
const passwordRules = [
  (v: string) => !!v || "Required.",
  (v: string) => (v && v.length >= 8) || "Min 8 characters"
]
const firstnameRules = [
  (v: string) => !!v || "Required.",
]
const lastnameRules = [
  (v: string) => !!v || "Required.",
]
const passwordMatch = () => password.value === passwordCheck.value || "Password must match";

const signUp = function () {
  const url = new URL('/signup', config.web_client_url);
  const req = SignupReq.create({
    email: email.value,
    password: password.value,
    firstname: firstname.value,
    lastname: lastname.value,
  });

  fetch(url, {
    method: 'POST',
    body: JSON.stringify(req),
  })
}

</script>

<template>
  <v-card class="px-6 py-6" rounded>
    <v-card-item class="justify-center mb-2" prepend-icon="mdi-account-plus">
      <v-card-title>
        SIGNUP
      </v-card-title>
      <template v-slot:prepend>
        <v-icon color="success"></v-icon>
      </template>
    </v-card-item>
    <v-divider></v-divider>
    <v-card-text class="mt-6">
      <v-form ref="signup-form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="6">
            <v-text-field v-model="firstname" :rules="firstnameRules" label="First Name" maxlength="20" underlined
              append-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>
          <v-col cols="6">
            <v-text-field v-model="lastname" :rules="lastnameRules" label="Last Name" maxlength="20" underlined
              append-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>
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
          <v-col cols="12">
            <v-text-field v-model="passwordCheck" :append-icon="showPasswordCheck ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[...passwordRules, passwordMatch]" :type="showPasswordCheck ? 'text' : 'password'"
              label="Confirm Password" hint="At least 8 characters" counter
              @click:append="showPasswordCheck = !showPasswordCheck" underlined required clearable>
            </v-text-field>
          </v-col>
          <v-col cols="12" align="center">
            <v-btn size="large" :disabled="!valid" variant="tonal" append-icon="mdi-account-circle"
              @click="signUp">Signup
              <template v-slot:append>
                <v-icon color="success"></v-icon>
              </template>
            </v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>
<style scoped></style>
