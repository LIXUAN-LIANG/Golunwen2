<script setup lang="js">
import '../css/style.css'
import '../css/boxicons.min.css'
import axios from 'axios'
import { ref } from 'vue'

axios.defaults.baseURL = 'http://localhost:8080'

const nameInput = ref('')
const passwordInput = ref('')
const telephoneInput = ref('')
const errorMessage = ref('')
const container = ref(null)

const toggleRegister = () => {
  container.value.classList.add('active')
}

const toggleLogin = () => {
  container.value.classList.remove('active')
}

async function loginPost() {
  errorMessage.value = ''
  
  if (!nameInput.value || !passwordInput.value) {
    errorMessage.value = '用户名或密码不能为空'
    return
  }

  try {
    const response = await axios.post('/login', {
      name: nameInput.value,
      password: passwordInput.value
    }, {
          headers: {
        'Content-Type': 'application/json'
      }
    })

    if (response.data.code === 200) {
      // 登录成功
      console.log('登录成功：', response.data)
      // 保存token
      if (response.data.token) {
        localStorage.setItem('token', response.data.token.token)
      }
      // 这里可以添加登录成功后的跳转逻辑
      window.location.href = '/index'
    } else {
      errorMessage.value = response.data.msg || '登录失败'
    }
  } catch (error) {
    console.error('登录时发生错误：', error)
    errorMessage.value = error.response?.data?.msg || '登录失败，请稍后重试'
  }
}

async function registerPost() {
  errorMessage.value = ''

  if (!nameInput.value || !passwordInput.value|| !telephoneInput.value) {
    errorMessage.value = '用户名或密码或者电话号码不能为空'
    return
  }

  try {
    const response = await axios.post('/register', {
      name: nameInput.value,
      password: passwordInput.value,
      telephone: telephoneInput.value
    }, {
      headers: {
        'Content-Type': 'application/json'
      }
    })

    if (response.data.code === 200) {
      // 登录成功
      console.log('注册成功：', response.data)
      // 保存token
      if (response.data.token) {
        localStorage.setItem('token', response.data.token.token)
      }
      // 这里可以添加登录成功后的跳转逻辑
    } else {
      errorMessage.value = response.data.msg || '注册失败'
    }
  } catch (error) {
    console.error('注册时发生错误：', error)
    errorMessage.value = error.response?.data?.msg || '注册失败，请稍后重试'
  }
}
</script>

<template>
  <div class="container" ref="container">
    <div class="form-box login">
      <form @submit.prevent="loginPost">
        <h1>登录</h1>
        <div class="input-box">
          <input type="text" placeholder="Username" required v-model="nameInput">
          <i class="bx bxs-user"></i>
        </div>
        <div class="input-box">
          <input type="password" placeholder="Password" required v-model="passwordInput">
          <i class="bx bxs-lock-alt"></i>
        </div>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        <div class="forgot-link">
          <a href="#">忘记密码?</a>
        </div>
        <button type="submit" class="btn">Login</button>
        <p>或者选用以下平台登录</p>
        <div class="social-icons">
          <a href="#"><i class="bx bxl-google"></i></a>
          <a href="#"><i class="bx bxl-facebook"></i></a>
          <a href="#"><i class="bx bxl-github"></i></a>
          <a href="#"><i class="bx bxl-linkedin"></i></a>
        </div>
      </form>
    </div>

    <div class="form-box register">
      <form @submit.prevent="registerPost">
        <h1>注册</h1>
        <div class="input-box">
          <input type="text" placeholder="Username" required v-model="nameInput">
          <i class="bx bxs-user"></i>
        </div>
        <div class="input-box">
          <input type="text" placeholder="Telephone" required v-model="telephoneInput">
          <i class="bx bxs-phone"></i>
        </div>
        <div class="input-box">
          <input type="password" placeholder="Password" required v-model="passwordInput">
          <i class="bx bxs-lock-alt"></i>
        </div>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        <div class="forgot-link">
          <a href="#">忘记密码?</a>
        </div>
        <button type="submit" class="btn">Register</button>
        <p>或者选用以下平台登录</p>
        <div class="social-icons">
          <a href="#"><i class="bx bxl-google"></i></a>
          <a href="#"><i class="bx bxl-facebook"></i></a>
          <a href="#"><i class="bx bxl-github"></i></a>
          <a href="#"><i class="bx bxl-linkedin"></i></a>
        </div>
      </form>
    </div>

    <div class="toggle-box">
      <div class="toggle-panel toggle-left">
        <h1>Hello, Welcome!</h1>
        <p>还没有一个账户嘛?</p>
        <button class="btn register-btn" @click="toggleRegister">Register</button>
      </div>

      <div class="toggle-panel toggle-right">
        <h1>Welcome Back!</h1>
        <p>已经有一个账户啦?</p>
        <button class="btn login-btn" @click="toggleLogin">Login</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.error-message {
  color: #ff0000;
  margin: 10px 0;
  text-align: center;
}
</style>