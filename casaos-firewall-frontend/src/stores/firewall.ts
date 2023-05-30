import { ref, computed, reactive } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import type { Port } from '@/types'

const baseHost = 'http://localhost'
const PortMetaData = {
  fromJson(object: any):Port{
    return {
      port: object.port ? object.port : "error port",
      protocol: object.protocol ? object.protocol : "error protocol",
      action: object.action ? object.port : "error action",
    }
  }
}

interface FirewallState {
  Ports: Port[],
  firewallState: boolean,
  firewallVersion: string,
}

export const useFirewallStore = defineStore('firewall', {
  state: (): FirewallState => ({
    Ports : [],
    firewallState : true,
    firewallVersion : "",
  }),
  actions: {
    async getPort() {
      const ports = (await axios.get(baseHost+'/v2/firewall/port')).data
      this.Ports = ports.map(PortMetaData.fromJson)
    },
    async getFirewallState() {},
    async getFirewallVersion() {
      const version = (await axios.get(baseHost+'/v2/firewall/version')).data
      this.firewallVersion = version
    }
  }
})
