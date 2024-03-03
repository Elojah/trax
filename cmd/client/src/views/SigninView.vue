<script lang="ts">
import { useAuthStore } from '@/stores/auth';
import { defineComponent } from 'vue';

const authStore = useAuthStore();

export default defineComponent({
  data() {
    return {
      email: null as string | null,
      password: null as string | null
    };
  },
  methods: {
    signIn() {
      // Add your sign in logic here
      authStore.signinGoogle(this.email!);
      console.log('Signing in...');
    }
  }
});

</script>

<template>
  <div>
    <h1>Sign In</h1>
    <form>
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" />
      </div>
      <button @click="signIn">Sign In</button>
    </form>
  </div>
</template>

<style scoped>
/* Add your component styles here */
</style>
