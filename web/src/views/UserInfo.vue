<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
    <!-- Left navigation bar -->
    <NavigationBar 
      activeSection="userInfo"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />
   
    <!-- Main content area -->
    <div class="flex-1 flex flex-col">
      <!-- Header bar -->
      <div class="bg-black pt-5 pb-3 px-8 border-b border-white/50 ">
        <div class="flex items-center">
          <h1 class="text-5xl font-semibold text-white font-adlam">Execute</h1>
        </div>
      </div>
     
      <div class="flex flex-col bg-black items-center justify-center pt-[2rem]">
        <div class="w-64 h-64 rounded-full flex items-center justify-center overflow-hidden mb-4">
          <img
            v-if="userAvatar"
            :src="userAvatar.avatar"
            alt="User Avatar"
            class="w-full h-full object-cover"
          />
        </div>

        <div class="flex px-8 py-4 bg-infoBg items-center justify-center text-center rounded-3xl border-borderColor border-2 border-solid">
          <h1 class="text-[48px]" v-if="userData">{{ userData.display_name || 'Set your data'}}</h1>
          <h1 class="text-[48px]" v-else>Loading...</h1>
        </div>
        <div class="px-4 mt-[-1px] bg-infoBg text-center rounded-3xl border-borderColor border-2 border-solid">
          <h1 class="text-[18px]" v-if="userData">{{ userData.role || 'None'}}</h1>
          <h1 class="text-[18px]" v-else>Loading...</h1>
        </div>

        <div class="space-y-2 mt-4 w-full max-w-xl px-6 py-12 bg-infoBg bg-white-200 rounded-xl border-borderColor border-2 border-solid relative">
          <!-- Dots/settings -->
           <button @click="toggleUserSettings">
            <div class="absolute top-3 right-6 cursor-pointer">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="text-gray-400">
              <circle cx="3" cy="12" r="1.5"></circle>
              <circle cx="11" cy="12" r="1.5"></circle>
              <circle cx="19" cy="12" r="1.5"></circle>
            </svg>
          </div>
           </button>

           <div v-if="userData" class="flex flex-col gap-3">
            <!-- Email -->
            <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
              <span class="text-left absolute text-white">Username: </span>
              <span class="text-center w-full text-white">{{ userData.username }}</span>
            </div>
            <!-- Phone Number -->
            <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
              <span class="text-left absolute text-white">Phone: </span>
              <span class="text-center w-full text-white">{{ userData.phone || 'Set your user data'}}</span>
            </div>
            <!-- Birth date -->
            <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
              <span class="text-left absolute text-white">Birth date: </span>
              <span class="text-center w-full text-white">{{ userData.birthdate || 'Set your user data'}}</span>
            </div>
           </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Mascot section -->
  <div class="fixed bottom-4 right-4 pointer-events-none mt-6 hidden lg:block">
      <img
        v-if="activeGif === null"
        src="/Bunny/mirrorUp.gif"
        class="mx-auto w-64"
        :key="`mirror-${gifTimestamp}`"
      />
      <img
        v-if="activeGif === 'userInfo'"
        src="/Bunny/task.gif"
        class="mx-auto w-64"
        :key="`task-${gifTimestamp}`"
      />
      <img
        v-if="activeGif === 'settingsUp'"
        src="/Bunny/settingsUp.gif"
        class="mx-auto w-64"
        :key="`settingsUp-${gifTimestamp}`"
      />
      <img
        v-if="activeGif === 'settingsDown'"
        src="/Bunny/settingsDown.gif"
        class="mx-auto w-64"
        :key="`settingsDown-${gifTimestamp}`"
      />
      <img
          v-if="activeGif === 'info'"
          src="/Bunny/infoGif.gif"
          class="mx-auto w-64"
          :key="`mirror-${gifTimestamp}`"
        />
    </div>

  <!-- Dialog settings & design -->
  <SettingsDialog v-model:show="openSettings" @navigate="navigateTo" @close="() => { openSettings = false; activeGif = 'settingsDown'; gifTimestamp = Date.now()}"/>
  <!-- Info Dialog -->
  <InfoDialog v-model:show="openInfo" @close="() => { openInfo = false; activeGif = null; gifTimestamp = Date.now()}"/>
  <!-- User Settings Dialog -->
  <SettingsUserDialog 
    v-model:show="openUserSettings" 
    :userData="userData" 
    @update:userData="updateUserData"
    @update:userAvatar="updateUserAvatar"
    @close="() => { openUserSettings = false; activeGif = null; gifTimestamp = Date.now()}"
  />
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ref, onMounted, computed } from 'vue'
import NavigationBar from './NavigationBar.vue'
import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
import InfoDialog from './PresetsDialogs/InfoDialog.vue'
import SettingsUserDialog from './PresetsDialogs/SettingsUserDialog.vue'

const router = useRouter()
const userData = ref(null)
const userAvatar = ref(null)
const openSettings = ref(false)
const openInfo = ref(false)
const openUserSettings = ref(false)

// Navigation using NavigationBar Preset
const navigateTo = (section) => {
    gifTimestamp.value = Date.now()
    activeGif.value = null
  if (section === 'dashboard') {
    router.push('/dashboard')
  } else if (section === 'teamsInfo') {
    router.push('/teamsInfo')
  } else if (section === 'userInfo') {
    router.push('/userInfo')
  } else if (section === 'calendar') {
    router.push('/calendar')
  } else if (section === 'logout') {
    router.push('/')
  }
}
    
const gifTimestamp = ref(Date.now())

const toggleSettings = () => {
  openSettings.value = !openSettings.value
  if (openSettings.value) {
        activeGif.value = 'settingsUp'
        gifTimestamp.value = Date.now()
      } else {
        activeGif.value = 'settingsDown'
        gifTimestamp.value = Date.now()
      }
}

const toggleInfo = () => {
  openInfo.value = !openInfo.value
  if (openInfo.value) {
        activeGif.value = 'info'
        gifTimestamp.value = Date.now()
      } else {
        activeGif.value = null
        gifTimestamp.value = Date.now()
      }
}

const toggleUserSettings = () => {
  openUserSettings.value = !openUserSettings.value
  if(openUserSettings.value) {
    activeGif.value = 'userInfo'
  } else {
    activeGif.value = null
  }
   gifTimestamp.value = Date.now()
}

// Update userData when dialog emits changes
const updateUserData = (newUserData) => {
  userData.value = newUserData
}
const updateUserAvatar = (newUserAvatar) => {
  userAvatar.value = newUserAvatar
}

onMounted(async () => {
  try {
    const userResponse = await fetch('api/v1/user/current', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include'
    })

    if (userResponse.ok) {
      userData.value = await userResponse.json()

      const avatarResponse = await fetch(`api/v1/avatar?id=${userData.value.id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: 'include'
      })

      if (avatarResponse.ok) {
        userAvatar.value = await avatarResponse.json()
      } else {
        console.error('Failed to fetch avatar', avatarResponse.status)
      }
    } else {
      console.error('Failed to fetch user data', userResponse.status)
    }
  } catch(error) {
    console.error('Error fetching user data:', error)
  }
})

  //Mascot
  const activeGif = ref<'userInfo' | 'settingsUp' | 'settingsDown' | 'info' | null>(null);

</script>
