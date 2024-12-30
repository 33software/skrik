<template>
    <div>
      <h2>Login</h2>
      <form @submit.prevent="loginUser">
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Password" required />
        <button type="submit">Login</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: '',
        password: '',
      };
    },
    methods: {
      async loginUser() {
        const response = await fetch('http://localhost:8080/api/account/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
          }),
        });
        const data = await response.json();
        if (data.token) {
          localStorage.setItem('token', data.token);
          this.$router.push('/');
        } else {
          alert('Login failed');
        }
      },
    },
  };
  </script>
  