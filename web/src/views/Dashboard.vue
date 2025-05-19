<template>
   <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
    <!-- Left navigation bar -->
    <NavigationBar 
      activeSection="dashboard"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />

    <!-- Main content area -->
    <div class="flex-1 flex flex-col min-h-screen overflow-y-auto">
      <!-- Header bar -->
      <div class="bg-black pt-5 pb-3 px-8 border-b border-white/50 ">
        <div class="flex items-center">
          <h1 class="text-5xl font-semibold text-white font-adlam">Execste</h1>
        </div>
      </div>
      
      <!-- Team info section -->
      <div class="px-16 py-6">
        <div class="flex items-center space-x-4">
          <h2 class="text-3xl text-white font-adlam">{{ teamData.name }}</h2>
           <!-- Points -->
          <div class="flex space-x-3">
            <div class="ml-1 text-[16px] text-white-300 font-adlam">Points earned: {{ teamData.pointsScore || 'No points earned yet' }} </div>
            
            <div class="ml-2 text-[16px] text-white-300 font-adlam">Points left: {{ teamData.points || 'No points yet' }} </div>
          </div>
        </div>        
      </div>
       
      <!-- Main content area -->
      <div class="flex-grow bg-black m-12">
        <!-- Dashboard elements  -->
         <div class="flex flex-wrap gap-12 items-start">
            <!-- To-Do -->
              <div class="w-[28rem] flex-none bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl  font-bold">To-Do</h1>
                </div>

                <!-- Adding tasks design and logic  -->
                <draggable 
                  v-model="toDoTasks"
                  group="tasks"
                  item-key="id"
                  class="min-h-24"
                  data-step="1"
                  @change="(event) => onDragChange(event, 1)">
                    <template #item="{element: item, index}">
                      <div class="bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between cursor-move">
                        <div class="flex items-center">
                            <div :class="{'bg-green-500': item.completed, 'bg-gray-500': !item.completed}" class="w-4 h-4 rounded-full mr-3 cursor-pointer" @click="toggleCompletion(item)"></div>
                            <button @click="openTaskSettings(item, 'todo', index)">
                          <!-- Elements inside  -->
                        <div class="text-left">
                            <div class="text-xl">{{ item.name.trim() }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ formatDate(item.dueDate) }}</div> <!-- Show only when it's provided -->
                            </div>
                        </button>
                      </div>
                      <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                  </template>
                  <!-- No tasks yet  -->
                   <div v-if="!toDoTasks || toDoTasks.length === 0" class="">No tasks yet. Add one</div>
                </draggable>
            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('todo')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-sans font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>


          <!-- In progress -->
          <div class="w-[28rem] bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl  font-bold">In progress</h1>
                </div>

                <!-- Adding tasks design and logic  -->
                <draggable 
                  v-model="inProgressTasks"
                  group="tasks"
                  item-key="id"
                  class="min-h-24"
                  data-step="2"
                  @change="(event) => onDragChange(event, 2)">
                  
                    <template #item="{element: item, index}">
                        <div class="bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between cursor-move">
                        <div class="flex items-center">
                            <div :class="{'bg-green-500': item.completed, 'bg-gray-500': !item.completed}" class="w-4 h-4 rounded-full mr-3 cursor-pointer" @click="toggleCompletion(item)"></div>
                            <button @click="openTaskSettings(item, 'inProgress', index)">
                          <!-- Elements inside  -->
                        <div class="text-left">
                            <div class="text-xl">{{ item.name.trim() }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ formatDate(item.dueDate) }}</div> <!-- Show only when it's provided -->
                            </div>
                        </button>
                      </div>
                      <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                  </template>

                  <!-- No tasks yet  -->
                <div v-if="inProgressTasks.length == 0" class="">No tasks yet. Add one</div>
                </draggable>

            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('inProgress')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-sans font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>


          <!-- Completed -->
          <div class="w-[28rem] bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl font-bold">Completed</h1>
                </div>

                <!-- Adding tasks design and logic  -->
                <draggable 
                  v-model="completedTasks"
                  group="tasks"
                  item-key="id"
                  class="min-h-24"
                  data-step="3"
                  @change="(event) => onDragChange(event, 3)">
                  
                    <template #item="{element: item, index}">
                        <div class="bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between cursor-move">
                        <div class="flex items-center">
                            <div :class="{'bg-green-500': item.completed, 'bg-gray-500': !item.completed}" class="w-4 h-4 rounded-full mr-3 cursor-pointer" @click="toggleCompletion(item)"></div>
                            <button @click="openTaskSettings(item, 'completed', index)">
                          <!-- Elements inside  -->
                        <div class="text-left">
                            <div class="text-xl">{{ item.name.trim() }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ formatDate(item.dueDate) }}</div>
                            </div>
                        </button>
                      </div>
                      <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                  </template>

                  <!-- No tasks yet  -->
                <div v-if="completedTasks.length == 0 " class="">No tasks yet. Add one</div>
                </draggable>

            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('completed')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-sans font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>
        </div>
    </div>
  </div>
</div>

<!-- Task Modal -->
  <div v-if="isModalOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="border-2 border-solid border-borderColor bg-infoBg text-white p-6 rounded-xl w-full max-w-md space-y-4">
        <h2 class="text-2xl text-center font-semibold">Add Task</h2>
        <div>
          <label class="block mb-1">Task Name</label>
          <input v-model="task.name" type="text" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Description</label>
          <input v-model="task.description" type="text" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div></div>
        <div>
          <label class="block mb-1">Points</label>
          <input v-model="task.points" type="number" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Due Date</label>
          <input v-model="task.dueDate" type="datetime-local" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" /> <!-- input type adds calendar on the right btw -->
        </div>
        <p v-if="taskError" class = "text-error text-sm mt-1 text-center"> {{ taskError }}</p>
        <div class="flex justify-end space-x-2">
          <button @click="closeModal" class="px-4 py-2 rounded bg-gray-600 hover:bg-gray-500">Cancel</button>
          <button @click="saveTask" class="px-4 py-2 rounded bg-blue-600 hover:bg-blue-500">Save</button>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <div v-if="isTaskSettingOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
  <div class="border-2 border-solid border-borderColor bg-infoBg text-white p-6 rounded-xl w-full max-w-md space-y-4">
    <h2 class="text-2xl text-center font-semibold">{{ selectedTask?.name }}</h2>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Description:</h3>
      <!-- If no description 'No description provided'-->
      <p class="text-gray-300">{{ selectedTask?.description || 'No description provided' }}</p>
    </div>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Points:</h3>
      <!-- If no points 'No points set'-->
      <p class="text-gray-300">{{ selectedTask?.points || 'No points set' }} </p>
    </div>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Due Date:</h3>
      <!-- If no date 'No due date set'-->
      <p class="text-gray-300">{{ formatDate(selectedTask?.dueDate) || 'No due date set' }}</p>
    </div>
    <div class="flex justify-between space-x-2 mt-4">
      <button @click="handleTaskDeletion" class="px-4 py-2 rounded bg-red-600 hover:bg-red-500">Delete Task</button>
      <button @click="closeTaskSettings" class="px-4 py-2 rounded bg-gray-600 hover:bg-gray-500">Close</button>
    </div>
  </div>
</div>
  <!-- Mascot section -->
  <div class="fixed bottom-4 right-4 pointer-events-none mt-6 hidden lg:block">
      <img
        v-if="activeGif === null"
        src="/Bunny/standing.png"
        class="mx-auto w-64"
        :key="`standing-${gifTimestamp}`"
      />
      <img
        v-if="activeGif === 'task'"
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
    <SettingsDialog v-model:show="openSettings"  @navigate="navigateTo" @close="() => { openSettings = false; activeGif = 'settingsDown'; gifTimestamp = Date.now()}"/>
    <!-- Info Dialog -->
    <InfoDialog v-model:show="openInfo" @close="() => { openInfo = false; activeGif = null; gifTimestamp = Date.now()}" />
</template>


<!-- Setting up router navigation  -->
<script lang="ts" setup>
    import { useRouter } from 'vue-router'
    import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
    import { onMounted, ref } from 'vue';
    import draggable from 'vuedraggable';
    import NavigationBar from './NavigationBar.vue'
    import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
    import InfoDialog from './PresetsDialogs/InfoDialog.vue'
    import { fetchTeamInfo, teamData } from './PresetsScripts/GroupInfo'
    import { toggleCompletion } from './PresetsScripts/taskCompletion';
    import { updateTaskStep, onDragChange } from './PresetsScripts/onDragChange';
    import { deleteTask } from './PresetsScripts/DeleteTask';
    

    const router = useRouter()
    //Navigation
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
    const taskError = ref('')
    const openSettings = ref(false)
    const openInfo = ref(false)
    
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

    //Adding tasks logic
    const isModalOpen = ref(false) //For showing add tasks dialog
    

    function openModal(listType: string) {
      activeTaskList.value = listType
      isModalOpen.value = true
      task.value = {
        name: '',
        description: '',
        points: 0,
        dueDate: '',
        completed: false,
        id: undefined
      }
      activeGif.value = 'task'
    }
    
    function closeModal() {
      isModalOpen.value = false
      activeGif.value = null
    }

    interface TaskItem {
      id?: number,
      name: string,
      description: string,
      points: number,
      dueDate: string,
      completed: boolean,
      step?: number,
    }
    
    const task = ref<TaskItem>({
      name: '',
      description: '',
      points: 0, //Number
      dueDate: '',
      completed: false,
      id: undefined
    })
    

    const toDoTasks = ref<TaskItem[]>([])
    const inProgressTasks = ref<TaskItem[]>([])
    const completedTasks = ref<TaskItem[]>([])
    const activeTaskList = ref('')

    function formatDate(dateStr: string | undefined): string {
      return dateStr ? dateStr.slice(0, 16).replace('T', ' ') : '';
    }


    async function fetchTasks() {
      try { 
      //clearing arrays
      toDoTasks.value = []
      inProgressTasks.value = []
      completedTasks.value = []
      
        const response = await fetch('api/v1/task', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
        })
        if (response.ok) {
          const tasks = await response.json()

          if (Array.isArray(tasks)) {
             tasks.forEach((task: any) => {
            const taskItem: TaskItem = {
              id: task.id,
              name: task.name,
              description: task.description,
              points: task.pointsValue,
              dueDate: task.dueDate,
              completed: task.completed || false,
              step: task.step,
            }
              if (task.step === 1) {
              toDoTasks.value.push(taskItem)
              } else if (task.step === 2) {
              inProgressTasks.value.push(taskItem)
              } else if (task.step === 3) {
              completedTasks.value.push(taskItem)
            }
          })
          }
        }
      } catch (error) {
        console.error('Fetching tasks failed:', error)
      }
    }
 

    async function saveTask() {
      if (!task.value.name.trim()) {
        taskError.value = 'Task name is required'
        return
      }

      if (!task.value.dueDate.trim()) {
        taskError.value = 'Due date is required'
        return
      }
      
      const stepId = {
        todo: 1,
        inProgress: 2,
        completed: 3
      }

      const load = {
        name: task.value.name.trim(),
        description: task.value.description,
        dueDate: new Date(task.value.dueDate).toISOString(),
        pointsValue: Number(task.value.points),
        step: stepId[activeTaskList.value] || 1
      }

      
    taskError.value = '';
    try {
        const response = await fetch('api/v1/task', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(load),
          credentials: 'include',
        })        
      
        if (response.ok) {
          const result = await response.json();
          
          const newTask = {
            ...task.value,
            id: result.id 
          };
          
      await fetchTasks();
      await fetchTeamInfo();
    } if (!response.ok) {
      //Setting errors for codes
        if (response.status === 403) {
          taskError.value = 'To create a task, join/create group';
        } 
        return;
      }
      console.log("Load to server:", load)
      closeModal()

    } catch (error) {
        taskError.value = `Connection error: ${error.message || 'Unknown error'}`;
        console.error('Request error:', error);
    }
  } 

    
    //Tasks settings
    const isTaskSettingOpen = ref(false)
    const selectedTask = ref<TaskItem | null>(null)
    const selectedTaskList = ref('')
    const selectedTaskIndex = ref(-1)

  
    function openTaskSettings(task: TaskItem, listType: string, index: number) {
      selectedTask.value = task
      selectedTaskList.value = listType
      selectedTaskIndex.value = index
      isTaskSettingOpen.value = true
    }
    
    function closeTaskSettings() {
      isTaskSettingOpen.value = false
      activeGif.value = null
    }

    async function handleTaskDeletion() {
      try {
        if (!selectedTask.value || selectedTask.value.id === undefined) {
          console.error('No valid task selected for deletion');
          return;
        }
        
        const success = await deleteTask(selectedTask.value.id); //call to the script
        
        if (success) {
          console.log('Successfully deleted task');
          //Data refresh
          await fetchTasks();
          await fetchTeamInfo();
        } else {
          console.error('Delete task failed');
        }
        
        closeTaskSettings();
      } catch (error) {
        console.error('Error in handler:', error);
        closeTaskSettings();
      }
    }

    //Mascot
    const activeGif = ref<'task' | 'settingsUp' | 'settingsDown' | 'info' | null>(null);

    onMounted(() => {
      fetchTasks()
      fetchTeamInfo()
    })
</script>