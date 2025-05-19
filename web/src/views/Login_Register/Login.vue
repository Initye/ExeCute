<template>
    <div class="min-h-screen flex items-center justify-center bg-black font-adlam">
        <!-- Text spacer -->
      <div class="text-center space-y-6">
        <!-- Header -->
        <h1 class="text-white select-none text-9xl font-bold pb-8">Execute</h1>
        <!-- Login input -->
        <div>
          <label class="block text-white select-none text-3xl mb-2">Login</label>
          <input
            type="text"
            placeholder="Enter text..."
            class="w-64 px-4 py-2 rounded text-black focus:outline-none"
            v-model= "login"
            @focus="activeGif = 'login'"
          />
          <p v-if="loginError" class = "text-error text-sm">{{ loginError }}</p>
          
        </div>
        <!-- Password input  -->
        <div>
          <label class="block text-white select-none text-3xl mb-2">Password</label>
          <input
            type="password"  
            placeholder="Enter text..."
            class="w-64 px-4 py-2 rounded text-black focus:outline-none"
            v-model = "password"
            @focus="activeGif = 'password'"
          />
          <p v-if="passwordError" class = "text-error text-sm">{{  passwordError }}</p>
        </div>
          <div>
            <div>
              <button class="bg-white hover:bg-gray-300 text-gray-800  py-1 px-9 border border-gray-400 rounded "
              @click = "loginUser"
              >
              Submit
              </button>
          </div>
          <span class="text-gray-400">Or</span>
          <div>
              <button class="bg-white hover:bg-gray-300 text-gray-800 py-1 px-9 border border-gray-400 rounded "
              @click = "goToRegister"
              >
              Register
              </button>
          </div>
          <p v-if="loggingError" class = "text-error text-sm mt-1"> {{ loggingError }}</p>
          <p v-if="loggingSuccess" class = "text-accepted text-sm">{{ loggingSuccess }}</p>
        </div>
      </div>
    </div>
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

<!-- Setting up router navigation  -->
<script lang="ts" setup>
    import { useRouter } from 'vue-router'
    import { ref } from 'vue';
    
    const router = useRouter()

    
    const goToRegister = () => {
      router.push('/register')
    }

  const login = ref('');
  const password = ref('');

  const loginError = ref('');
  const passwordError = ref('');
  const loggingError = ref('');
  const loggingSuccess = ref('');

  const loginUser = async () => {
    loginError.value = '';
    passwordError.value = '';
    loggingError.value = '';
    loggingSuccess.value = '';

    //Setting login error
    if (!login.value.trim()) {
    loginError.value = 'Login is required';
    return;
  }
  //Setting password error 
  if (!password.value.trim()) {
    passwordError.value = 'Password is required';
    return;
  }

  

  try { 
    const response = await fetch('api/v1/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
      username: login.value,
      password: password.value,
    }),
    credentials: 'include',
    });

    if (!response.ok) {
      //Setting errors for codes
      if (response.status === 400) {
        loggingError.value = 'Invalid username or password';
      } else if (response.status === 401) {
        loggingError.value = 'Invalid username or password';
      } else {
        loggingError.value = `Error: ${response.status} ${response.statusText}`;
      }
      return;
    }

    const data = await response.json();
    loggingSuccess.value = 'Logging in...';
    setTimeout(() => {
    router.push('/dashboard'); 
    }, 1500);
  } catch (error: any) { 
    loggingError.value = `Connection error: ${error.message || 'Unknown error'}`;
    console.error('Logging error:', error);
    }
  }

  //Mascot section
  const activeGif = ref<'login' | 'password' | 'register' | null>(null);

</script>