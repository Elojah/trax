import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore({
  id: 'user',
  state: () => ({
    users: null,
  }),
  actions: {
    async fetchUsers() {
      // this.users = await fetch('/api/users').then((res) => res.json())
    },
  },
})
