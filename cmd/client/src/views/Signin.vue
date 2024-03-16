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

      valid: null as boolean | null,
      showPassword: false,

      emailRules: [
        (v: string) => !!v || "Required",
        (v: string) => /.+@.+\..+/.test(v) || "Email must be valid"
      ],
      passwordRules: [
        (v: string) => !!v || "Required.",
        (v: string) => (v && v.length >= 8) || "Min 8 characters"
      ],
    };
  },
  methods: {
    signIn() {
      // Add your sign in logic here
      this.authStore.signinGoogle(this.email!);
      console.log('Signing in...');
    }
  }
});

</script>

<template>
  <v-card class="px-6 py-8 rounded-xl">
    <v-card-item class="justify-center mb-12" prepend-icon="mdi-login">
      <v-card-title>
        SIGNIN
      </v-card-title>
    </v-card-item>
    <v-card-text>
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
          <v-col cols="12" align="center">
            <v-btn size="large" :disabled="!valid" variant="tonal" append-icon="mdi-account-circle"
              @click="signIn">Signin
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
