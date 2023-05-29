import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useFirewallStore = defineStore('firewall', () => {
  const port = ref(0)
  const firewallState = ref(true)
  const firewallVersion = ref('')
  function getPort() {
    
  }
  
  return { port, getPort }
})
