<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
   <!-- Left navigation bar -->
   <NavigationBar 
      activeSection="calendar"
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
     
     <!-- Main content area -->
     <div class="flex-grow bg-black">
        <div class="flex justify-between items-center">
          <div class="flex items-center m-4">

            <!-- PreviousWeek button nav -->
            <button @click="previousWeek" class="bg-black text-white ">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>

            <!-- Today button nav -->
            <button @click="today" class="bg-borderColor text-white px-4 py-2 rounded-md">
              Today
            </button>

            <!-- NextWeek button nav -->
            <button @click="nextWeek" class="bg-black text-white ">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
              <div class="text-sm ">{{  formatMonthYear(currentWeekStart) }}</div>
          </div>
          <div class="flex items-center">
        </div>
      </div>

      <div class="flex flex-col w-full h-screen">
        <div class="flex bg-black text-white border-t border-gray-700">
          <!-- Days of the week section -->
          <div 
            v-for="(day, index) in weekDays" 
            :key="index" 
            class="flex-1 border-r border-gray-700 text-center py-2"
          >
            <div class="font-bold text-xl">{{ formatWeekday(day) }}</div>
            <div class="font-bold text-xs">{{ formatDayNumber(day) }}</div>
            <div class="text-xs text-gray-300">{{ isToday(day) ? 'Today' : '' }}</div>
          </div>
        </div>
        
        <div class="flex-grow overflow-y-auto">
            <!-- grid for calendar -->
            <div class="grid grid-cols-7 h-full "> 
              <!-- column for a day -->
              <div 
                v-for="(day, dayIndex) in weekDays" 
                :key="dayIndex" 
                class="border-r border-gray-700 h-full "
              >
                <div class="p-2 h-full">

                  <div class="p-2 h-full">
                    <div 
                      v-for="task in tasksDay(day)" 
                      :key="task.id" 
                      class="bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between"
                    >
                      <div class="flex items-center">
                        <div class="w-4 h-4 rounded-full mr-3 bg-green-200"></div>
                        <div>
                          <div class="text-xl">{{ task.name }}</div>
                          <div class="text-sm text-gray-300">{{ task.description }}</div>
                        </div>
                      </div>
                      <span class="text-white text-sm">{{ task.pointsValue }} pts</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
</div>
 
    <!-- Dialog settings & design -->
    <SettingsDialog v-model:show="openSettings"  @navigate="navigateTo"/>
    <!-- Info Dialog -->
    <InfoDialog v-model:show="openInfo" />
</template>

<!-- Setting up router navigation  -->
<script lang="ts" setup>
  import { useRouter } from 'vue-router'
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
  import { computed, onMounted, ref } from 'vue';
  import { 
  format, 
  startOfWeek, 
  endOfWeek, 
  eachDayOfInterval, 
  addDays,
  subDays,
  addWeeks,
  subWeeks,
  isToday as dateFnsIsToday,
} from 'date-fns';
//Language for dates, zone
import { isSameDay, parseISO } from 'date-fns'
import { enAU } from 'date-fns/locale';
import NavigationBar from './NavigationBar.vue'
import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
import InfoDialog from './PresetsDialogs/InfoDialog.vue'

   const router = useRouter()
   //Navigation
   const navigateTo = (section) => {
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

  //Dialogs
   const openSettings = ref(false)
   const openInfo = ref(false)
   const toggleSettings = () => {
      openSettings.value = !openSettings.value
    }

    const toggleInfo = () => {
      openInfo.value = !openInfo.value
    } 

//Calendar
   //Current week
   const currentWeekStart = ref(startOfWeek(new Date(), { weekStartsOn: 1}))

   const weekDays = computed(() => {
    return eachDayOfInterval({
      start: currentWeekStart.value,
      end: endOfWeek(currentWeekStart.value, { weekStartsOn: 1})
    })
   })

   const tasks = ref([])

   onMounted(async () => {
      try {
        const response = await fetch('api/v1/task', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
        });
        const data = await response.json()
        if (Array.isArray(data)) {
          tasks.value = data;
        } else {
          tasks.value = [];
        }
      } catch (error) {
        console.error('Fetching tasks failed:', error);
        tasks.value = [];
      }
    })
    //show per dueDate
    const tasksDay = (day: Date) => {
      return tasks.value.filter((task) => 
      isSameDay(parseISO(task.dueDate), day))
    }
    
//Navigations
   //Next week nav
   const nextWeek = () => {
    currentWeekStart.value = addWeeks(currentWeekStart.value, 1)
   }
   //Prev week nav
   const previousWeek = () => {
    currentWeekStart.value = subWeeks(currentWeekStart.value, 1)
   }
   //Today nav
   const today = () => {
    currentWeekStart.value = startOfWeek(new Date(), { weekStartsOn: 1})
   }

// Date functions
   //Month current
    const formatMonthYear = (date: Date) => {
      return format(date, 'MMMM yyyy', { locale: enAU });
    };
    //Days
    const formatDayNumber = (date: Date) => {
      return format(date, 'dd');
    };
    //Current week
    const formatWeekday = (date: Date) => {
      return format(date, 'EEEE', { locale: enAU });
    };
    //Current day
    const isToday = (date: Date) => {
      return dateFnsIsToday(date);
    };

    
    //Delete later when delete placeholder for elements inside
    const active = ref(false)
</script>