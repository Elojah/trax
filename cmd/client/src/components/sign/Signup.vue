<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';
import { SignupReq } from '@internal/user/dto/user';
import { toRefs, ref } from "vue";

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
      // form?.value?.reset()
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
</template>
<style scoped></style>
