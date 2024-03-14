<script lang="ts">
import { useAuthStore } from '@/stores/auth';
import { defineComponent } from 'vue';


export default defineComponent({
  setup() {
    const authStore = useAuthStore();

    return {
      authStore
    };
  },
  data() {
    return {
      email: null as string | null,
      password: null as string | null,
      firstname: null as string | null,
      lastname: null as string | null,

      valid: null as boolean | null,
      passwordCheck: null as string | null,
      showPassword: false,
      showPasswordCheck: false,

      emailRules: [
        (v: string) => !!v || "Required",
        (v: string) => /.+@.+\..+/.test(v) || "Email must be valid"
      ],
      passwordRules: [
        (v: string) => !!v || "Required.",
        (v: string) => (v && v.length >= 8) || "Min 8 characters"
      ],
      firstnameRules: [
        (v: string) => !!v || "Required.",
      ],
      lastnameRules: [
        (v: string) => !!v || "Required.",
      ],
    };
  },
  computed: {
    passwordMatch() {
      return () => this.password === this.passwordCheck || "Password must match";
    }
  },
  methods: {
    signUp() {
      //   // Add your sign in logic here
      //   this.authStore.signinGoogle(this.email!);
      console.log('Signing up...');
    }
  }
});

</script>

<template>
  <v-card class="mx-auto px-6 py-8">
    <v-container>
      <v-form ref="signup-form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="12" sm="6" md="6">
            <v-text-field v-model="firstname" :rules="firstnameRules" label="First Name" maxlength="20"
              variant="underlined" append-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" md="6">
            <v-text-field v-model="lastname" :rules="lastnameRules" label="Last Name" maxlength="20"
              variant="underlined" append-icon="mdi-card-account-details" required clearable></v-text-field>
          </v-col>

          <v-col cols="12">
            <v-text-field v-model="email" label="Email" :rules="emailRules" append-icon="mdi-email" variant="underlined"
              required clearable></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="password" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="passwordRules" :type="showPassword ? 'text' : 'password'" label="Password"
              hint="At least 8 characters" counter @click:append="showPassword = !showPassword" variant="underlined"
              required clearable>
            </v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field v-model="passwordCheck" :append-icon="showPasswordCheck ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[...passwordRules, passwordMatch]" :type="showPasswordCheck ? 'text' : 'password'"
              label="Confirm Password" hint="At least 8 characters" counter
              @click:append="showPasswordCheck = !showPasswordCheck" variant="underlined" required clearable>
            </v-text-field>
          </v-col>
          <v-col cols="12" align="center">
            <v-btn size="large" :disabled="!valid" append-icon="mdi-account-circle" @click="signUp">Signup</v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-container>
  </v-card>
</template>
<style scoped></style>
