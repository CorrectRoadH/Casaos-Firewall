import { ref } from "vue";

export const useCasaos = () => {
    const baseHost = 'http://localhost';
    const background = ref("");
    const casaosVersion = ref("");
    return {baseHost,background,casaosVersion};
}