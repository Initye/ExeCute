<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
   <!-- Left navigation bar -->
   <NavigationBar 
      activeSection="teamsInfo"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />
   
   <!-- Main content area -->
   <div class="flex-1 flex flex-col min-h-screen overflow-y-auto">
     <!-- Header bar -->
     <div class="bg-black pt-5 pb-3 px-8 mb-20 border-b border-white/50 ">
       <div class="flex items-center">
         <h1 class="text-5xl font-semibold text-white font-adlam">Execute</h1>
       </div>
     </div>

     <!-- Main content area -->
      <div class="flex justify-center w-full">
        <div 
          @click="openTeamDialog" 
          class="w-full max-w-md mx-auto mb-12 px-6 py-4 bg-borderColor rounded-xl border-borderColor border-2 border-solid flex justify-center cursor-pointer hover:opacity-90 transition-opacity"
        >
            <span class="text-6xl">{{ teamData.name }}</span> 
        </div>
      </div>
      <!-- Members, meeting, scoreboard -->
      
     <div class="grid grid-cols-1 lg:grid-cols-3 gap-12 items-start w-full px-4 mx-auto max-w-7xl">
      <!-- Members section -->
      <div class="w-full bg-infoBg rounded-xl border-borderColor border-2 border-solid p-6 flex flex-col">
        <!-- Headline -->
        <div class="flex justify-center pb-8">
          <h1 class="text-white text-3xl font-bold">Members</h1>
        </div>

        <!-- Members -->
        <div class="max-h-64 overflow-y-auto pr-1">
          <div v-if="usersData.length > 0" class="space-y-2">
            <div v-for="user in usersData" :key="user.id" class="w-full bg-fillingInfo rounded-xl px-4 py-3 flex items-center space-x-4" >
              <div class="w-6 h-5 rounded-full bg-purple-50"></div>
              <div class="flex justify-between w-full text-white text-sm">
                <span> {{ user.display_name || user.username }} </span>
                <span class="text-right"> {{ user.role }} </span>
              </div>
            </div>
          </div> 
            <div v-else class="w-full bg-fillingInfo rounded-xl px-4 py-3 flex items-center space-x-4">
              <div class="flex justify-center w-full text-white text-sm">
                <span>Loading members...</span>
              </div>
          </div>
        </div>
      </div>
      
      <!-- Next meeting section -->
      <div @click="openMeetingDialog" class="w-full bg-infoBg rounded-xl border-borderColor border-2 border-solid p-6 flex flex-col justify-center hover:opacity-90 transition-opacity cursor-pointer">
        <div class="flex justify-center items-center pb-4">
          <h1 class="text-white text-4xl font-bold">Next meeting:</h1>
        </div>
        <div class="flex justify-center">
          <h1 class="text-white text-3xl">{{ teamData.meeting }}</h1>
        </div>
     </div>
     <!-- Scoreboard section -->
     <div class="w-full bg-infoBg rounded-xl border-borderColor border-2 border-solid p-6 flex flex-col">
          <!-- Headline -->
         <div class="flex justify-center pb-8">
            <h1 class="text-white text-3xl font-bold">Scoreboard</h1>
          </div>
        <div class="max-h-64 overflow-y-auto pr-1">
          <div v-if="teamsData.length > 0">
            <div v-for="team in teamsData" :key="team.id" class="w-full bg-fillingInfo rounded-xl mb-1 px-4 py-3 flex items-center space-x-4">
              <div class="flex justify-between w-full text-white text-sm">
                <span>{{ team.name }}</span>
                <span class="text-right">{{ team.points_score }}</span>
              </div>
            </div>
          </div>
          <div v-else>
            <div class="w-full bg-fillingInfo rounded-xl px-4 py-3 flex items-center space-x-4">
              <div class="flex justify-center w-full text-white text-sm">
                <span>Loading groups...</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
   </div>
 </div>
  <div class="fixed bottom-4 right-4 pointer-events-none mt-6 hidden lg:block">
        <img
          v-if="activeGif === null"
          src="/Bunny/teamsGif.gif"
          class="mx-auto w-64"
          :key="`mirror-${gifTimestamp}`"
        />
        <img
          v-if="activeGif === 'info'"
          src="/Bunny/infoGif.gif"
          class="mx-auto w-64"
          :key="`mirror-${gifTimestamp}`"
        />
        <img
          v-if="activeGif === 'settingsUp'"
          src="/Bunny/settingsUp.gif"
          class="mx-auto w-64"
          :key="`mirror-${gifTimestamp}`"
        />
        <img
          v-if="activeGif === 'settingsDown'"
          src="/Bunny/settingsDown.gif"
          class="mx-auto w-64"
          :key="`mirror-${gifTimestamp}`"
        />
    </div>
 <!-- Dialog settings & design -->
 <SettingsDialog v-model:show="openSettings" @navigate="navigateTo" @close="() => { openSettings = false; activeGif = 'settingsDown'; gifTimestamp = Date.now()}"/>
 <!-- Info Dialog -->
 <InfoDialog v-model:show="openInfo" @close="() => { openInfo = false; activeGif = null; gifTimestamp = Date.now()}" />
 <!-- Team Dialog -->
 <TeamDialog v-model:show="showTeamDialog" />
 <!-- Meeting Dialog -->
 <MeetingDialog v-model:show="showMeetingDialog" />
</template>

<!-- Setting up router navigation  -->
<script lang="ts" setup>
  import { useRouter } from 'vue-router'
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
  import { ref, watch, onMounted } from 'vue'
  import NavigationBar from './NavigationBar.vue'
  import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
  import InfoDialog from './PresetsDialogs/InfoDialog.vue'
  import TeamDialog from './PresetsDialogs/TeamDialog.vue'
  import MeetingDialog from './PresetsDialogs/MeetingDialog.vue'
  import { fetchTeamInfo, teamData } from './PresetsScripts/GroupInfo'
  import { fetchTeamUsersInfo, usersData } from './PresetsScripts/FetchUsersGroup'
  import { fetchScoreboardInfo, teamsData } from './PresetsScripts/FetchScoreboard'

  const router = useRouter()
  //navigation using NavigationBar Preset
  const navigateTo = (section) => {
    if (section === 'dashboard') {
      router.push('/dashboard')
    } else if (section === 'teamsinfo') {
      router.push('/teamsinfo')
    } else if (section === 'userInfo') {
      router.push('/userInfo')
    } else if (section === 'calendar') {
      router.push('/calendar')
    } else if (section === 'logout') {
      router.push('/')
    }
  }
  
  const openSettings = ref(false)
  const openInfo = ref(false)
  const showTeamDialog = ref(false)
  const showMeetingDialog = ref(false)


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
  
  const openTeamDialog = () => {
    showTeamDialog.value = true
  }

  const openMeetingDialog = () => {
    showMeetingDialog.value = true
  }
  
  //Mascot
  const gifTimestamp = ref(Date.now())
    const activeGif = ref<'info' | 'settingsUp' | 'settingsDown' | null>(null);


  onMounted(() => {
    fetchTeamInfo()
    fetchTeamUsersInfo()
    fetchScoreboardInfo()
  })
</script>