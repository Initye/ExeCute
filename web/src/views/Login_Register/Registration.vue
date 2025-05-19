<template>
  <div class="min-h-screen flex items-center justify-center bg-black font-adlam">
    <div class="text-center space-y-6">
      <h1 class="text-white select-none text-9xl font-bold pb-8">Execute</h1>
      <div>
        <label class="block text-white select-none text-3xl mb-2">Login</label>
        <input
          type="text"
          placeholder=" Login..."
          class="w-64 px-4 py-2 rounded text-black focus:outline-none"
          v-model="login"
          @focus="activeGif = 'login'"
        />
        <p v-if="loginError" class="text-error text-sm">{{ loginError }}</p>
      </div>
      <div>
        <label class="block text-white select-none text-3xl mb-2">Password</label>
        <input
          type="password"
          placeholder="Enter Password..."
          class="w-64 px-4 py-2 rounded text-black focus:outline-none"
          v-model="password"
          @focus="activeGif = 'password'"
        />
        <p v-if="passwordError" class="text-error text-sm">{{ passwordError }}</p>
      </div>
      <div>
        <label class="block text-white select-none text-3xl mb-2">Repeat Password</label>
        <input
          type="password"
          placeholder="Please repeat password..."
          class="w-64 px-4 py-2 rounded text-black focus:outline-none"
          v-model="rePassword"
          @focus="activeGif = 'password'"
        />
        <p v-if="rePasswordError" class="text-error text-sm">{{ rePasswordError }}</p>
      </div>

      <div>
        <button
          class="bg-white hover:bg-gray-300 text-gray-800 py-1 px-9 border border-gray-400 rounded"
          @click="registerUser"
        >
          Register
        </button>
        <p v-if="registrationError" class="text-error mt-2">{{ registrationError }}</p>
        <p v-if="registrationSuccess" class="text-accepted mt-2">{{ registrationSuccess }}</p>
      </div>
    </div>
  </div>
  <!-- Mascot section -->
  <div class="fixed bottom-4 right-4 pointer-events-none mt-6 hidden md:block">
      <img
        v-if="activeGif === null"
        src="/Bunny/standing.png"
        class="mx-auto w-64"
      />
      <img
        v-if="activeGif === 'login'"
        src="/Bunny/loginGif.gif"
        class="mx-auto w-64"
      />
      <img
        v-if="activeGif === 'password'"
        src="/Bunny/passwordGif.gif"
        class="mx-auto w-64"
      />
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

//Register
const login = ref('');
const password = ref('');
const rePassword = ref('');

const loginError = ref('');
const passwordError = ref('');
const rePasswordError = ref('');
const registrationError = ref('');
const registrationSuccess = ref('');

const registerUser = async () => {
  loginError.value = '';
  passwordError.value = '';
  rePasswordError.value = '';
  registrationError.value = '';
  registrationSuccess.value = '';

  try {
    const response = await fetch('api/v1/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: login.value,
        password: password.value,
        repassword: rePassword.value,
      }),
    });
   
    if (!response.ok) {
      if (response.status === 401) {
        registrationError.value = 'Invalid username or password';
      } else if (response.status === 400) {
        registrationError.value = 'Missing required fields';
      } else {
        registrationError.value = `Error: ${response.status} ${response.statusText}`;
      }
      return;
    }
      const data = await response.json();
      registrationSuccess.value = 'Registration successful! Redirecting to login...';
      setTimeout(() => {
        router.push('/'); //Pushing back to login
      }, 1500); //delay
    } catch (error: any) { 
      registrationError.value = `Connection error: ${error.message || 'Unknown error'}`;
      console.error('Registration error:', error);
  }
}

const activeGif = ref<'login' | 'password' | 'register' | null>(null);
</script>