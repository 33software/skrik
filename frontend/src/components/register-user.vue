<template>
  <div class="register">
    <form @submit.prevent="register">
      <div>
        <label for="email">Username:</label>
        <input type="username" v-model="username" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Register</button>
    </form>

    <div v-if="errorMessage" class="error">
      {{ errorMessage }}
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      email: '',
      password: '',
      errorMessage: null,
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post('http://localhost:8080/api/account/register', {
          username: this.username,
          email: this.email,
          password: this.password,
        });

        const token = response.data.token;
        localStorage.setItem('token', token);

        this.$router.push('/');
      } catch (error) {
        this.errorMessage = error.response.data.message || 'Registration failed';
      }
    },
  },
};
</script>

<style>
.error {
  color: red;
}
</style>