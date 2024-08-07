<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { SignupReq } from '@internal/user/dto/user';
import { toRefs, ref } from "vue";
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore()

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

const errorsStore = useErrorsStore()
const {
  message,
  success,
  error,
} = toRefs(errorsStore)

const signup = async function () {
  const req = SignupReq.create({
    email: email.value,
    password: password.value,
    firstname: firstname.value,
    lastname: lastname.value,
  });

  const resp = await authStore.signup(req);

  switch (resp.status) {
    case 200: // success
      message.value = 'Successfully Signup ! You can now signin.'
      success.value = true
      form?.value?.reset()
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
  <v-card class="px-6 py-6 rounded" variant="elevated">
    <v-card-item class="justify-center mb-2" prepend-icon="mdi-account-plus">
      <v-card-title>
        SIGNUP
      </v-card-title>
      <template v-slot:prepend>
        <v-icon color="primary"></v-icon>
      </template>
    </v-card-item>
    <v-divider color="primary"></v-divider>
    <v-card-text class="mt-6">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="6">
            <v-text-field v-model="firstname" :rules="firstnameRules" label="First Name" maxlength="20" underlined
              prepend-inner-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>
          <v-col cols="6">
            <v-text-field v-model="lastname" :rules="lastnameRules" label="Last Name" maxlength="20" underlined
              prepend-inner-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="email" label="Email" :rules="emailRules" preprnd-inner-icon="mdi-email" underlined
              required clearable></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="password" prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :rules="passwordRules"
              :type="showPassword ? 'text' : 'password'" label="Password" counter
              @click:append-inner="showPassword = !showPassword" underlined required clearable>
            </v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="passwordCheck" prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPasswordCheck ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[...passwordRules, passwordMatch]" :type="showPasswordCheck ? 'text' : 'password'"
              label="Confirm Password" counter @click:append-inner="showPasswordCheck = !showPasswordCheck" underlined
              required clearable>
            </v-text-field>
          </v-col>
          <v-col cols="12" align="center">
            <v-btn size="large" :disabled="!valid" variant="tonal" append-icon="mdi-account-circle"
              @click="signup">Signup
              <template v-slot:append>
                <v-icon color="primary"></v-icon>
              </template>
            </v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>
<style scoped></style>
